package releaseindex

import (
	"context"
	"fmt"
	"sync"

	"github.com/giantswarm/microerror"
	"github.com/giantswarm/micrologger"
	"github.com/giantswarm/versionbundle"
	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
	yaml "gopkg.in/yaml.v2"
)

type Config struct {
	// GithubToken is a Github token which is authorized to read from the
	// installations repository.
	GithubToken string
	Logger      micrologger.Logger
}

type ReleaseIndex struct {
	githubToken string
	logger      micrologger.Logger

	cache    map[Provider][]versionbundle.IndexRelease
	cacheMux *sync.RWMutex
}

func New(config Config) (*ReleaseIndex, error) {
	if config.GithubToken == "" {
		return nil, microerror.Maskf(invalidConfigError, "%T.GithubToken must not be empty", config)
	}
	if config.Logger == nil {
		return nil, microerror.Maskf(invalidConfigError, "%T.Logger must not be empty", config)
	}

	r := &ReleaseIndex{
		githubToken: config.GithubToken,
		logger:      config.Logger,

		cache:    map[Provider][]versionbundle.IndexRelease{},
		cacheMux: &sync.RWMutex{},
	}

	return r, nil
}

// GetVersion returns a component version for a given type and provider.
//
// NOTE: Release index is cached for a provider after first call. This makes it
// cheap to call this method multiple times with the same provider. But changes
// to the installations repository will not be visible.
func (r *ReleaseIndex) GetVersion(ctx context.Context, provider Provider, versionType VersionType, component Component) (string, error) {
	index, err := r.getIndex(ctx, provider)
	if err != nil {
		return "", microerror.Mask(err)
	}

	version, err := getVersion(index, versionType, component)
	if err != nil {
		return "", microerror.Mask(err)
	}

	return version, nil
}

func (r *ReleaseIndex) getFileContent(ctx context.Context, owner, repo, path string) (string, error) {
	var client *github.Client
	{
		t := &oauth2.Token{
			AccessToken: r.githubToken,
		}

		sts := oauth2.StaticTokenSource(t)
		tc := oauth2.NewClient(ctx, sts)

		client = github.NewClient(tc)
	}

	opt := &github.RepositoryContentGetOptions{}
	repoContent, _, _, err := client.Repositories.GetContents(ctx, owner, repo, path, opt)

	if err != nil {
		return "", microerror.Mask(err)
	}

	content, err := repoContent.GetContent()
	if err != nil {
		return "", microerror.Mask(err)
	}

	return content, nil
}

func (r *ReleaseIndex) getIndex(ctx context.Context, provider Provider) ([]versionbundle.IndexRelease, error) {
	var err error

	{
		r.cacheMux.RLock()
		index, ok := r.cache[provider]
		r.cacheMux.RUnlock()

		if ok {
			return index, nil
		}
	}

	var indexYAML string
	{
		owner := "giantswarm"
		repo := "installations"
		path := fmt.Sprintf("release/provider/%s.yaml", provider.String())

		indexYAML, err = r.getFileContent(ctx, owner, repo, path)
		if err != nil {
			return nil, microerror.Mask(err)
		}
	}

	var index []versionbundle.IndexRelease
	{
		err := yaml.Unmarshal([]byte(indexYAML), &index)
		if err != nil {
			return nil, microerror.Mask(err)
		}
	}

	{
		r.cacheMux.Lock()
		r.cache[provider] = index
		r.cacheMux.Unlock()
	}

	return index, nil
}

func getVersion(index []versionbundle.IndexRelease, versionType VersionType, component Component) (string, error) {
	if len(index) < 2 {
		return "", microerror.Maskf(executionFailedError, "expected at least 2 releases in the index but got %d", len(index))
	}

	var release versionbundle.IndexRelease
	{
		switch versionType {
		case VersionTypeLatest:
			release = index[0]
		case VersionTypeCurrent:
			release = index[1]
		default:
			return "", microerror.Maskf(executionFailedError, "unknown version type %#q", versionType.String())
		}
	}

	for _, a := range release.Authorities {
		if a.Name == component.String() {
			return a.Version, nil
		}
	}

	return "", microerror.Maskf(executionFailedError, "version of type %#q for component %#q not found", versionType.String(), component.String())
}
