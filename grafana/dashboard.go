package grafana

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/grafana-tools/sdk"
)

func PullDashboard(grafanaURL string, apiKey string, directory string, tag string) error {
	var (
		boardLinks []sdk.FoundBoard
		rawBoard   sdk.Board
		meta       sdk.BoardProperties
		err        error
	)
	ctx := context.Background()

	c := sdk.NewClient(grafanaURL, apiKey, sdk.DefaultHTTPClient)

	if boardLinks, err = c.Search(ctx, sdk.SearchTag(tag), sdk.SearchType(sdk.SearchTypeDashboard)); err != nil {
		return err
	}

	for _, link := range boardLinks {
		if rawBoard, meta, err = c.GetDashboardByUID(ctx, link.UID); err != nil {
			log.Printf("%s for %s\n", err, link.URI)
			continue
		}
		rawBoard.ID = 0
		b, err := json.Marshal(rawBoard)
		if err != nil {
			return err
		}
		if err = writeToFile(directory, b, meta.Slug, tag); err != nil {
			return err
		}
	}
	return nil
}

func PushDashboard(grafanaURL string, apiKey string, directory string) error {
	var (
		filesInDir []os.FileInfo
		rawBoard   []byte
		err        error
	)

	ctx := context.Background()
	c := sdk.NewClient(grafanaURL, apiKey, sdk.DefaultHTTPClient)
	if filesInDir, err = ioutil.ReadDir(directory); err != nil {
		return err
	}
	for _, file := range filesInDir {
		if strings.HasSuffix(file.Name(), ".json") {
			if rawBoard, err = ioutil.ReadFile(fmt.Sprintf("%s/%s", directory, file.Name())); err != nil {
				log.Println(err)
				continue
			}
			var board sdk.Board
			if err = json.Unmarshal(rawBoard, &board); err != nil {
				log.Println(err)
				continue
			}
			params := sdk.SetDashboardParams{
				FolderID:  sdk.DefaultFolderId,
				Overwrite: true,
			}
			if _, err := c.SetDashboard(ctx, board, params); err != nil {
				log.Printf("error on importing dashboard %s", board.Title)
				continue
			}
		}
	}
	return nil
}
