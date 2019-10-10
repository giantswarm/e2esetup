package catalog

import (
	"context"
	"fmt"

	"github.com/giantswarm/microerror"
	"github.com/google/go-github/github"
)

// GetLatestChart returns the latest chart tarball file in the specified catalog.
func GetLatestChart(ctx context.Context, catalog, app string) (string, error) {
	client := github.NewClient(nil)

	var path string
	{
		query := fmt.Sprintf("repo:giantswarm/%s filename:%s", catalog, app)
		searchOption := github.SearchOptions{
			Sort: "indexed",
		}
		result, _, err := client.Search.Code(ctx, query, &searchOption)
		if err != nil {
			return "", microerror.Mask(err)
		}

		path = result.CodeResults[0].GetPath()
	}

	var downloadURL string
	{
		r, _, _, err := client.Repositories.GetContents(ctx, "giantswarm", catalog, path, nil)
		if err != nil {
			return "", microerror.Mask(err)
		}

		downloadURL = r.GetDownloadURL()
	}

	return downloadURL, nil
}
