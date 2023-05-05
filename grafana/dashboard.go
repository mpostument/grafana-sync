package grafana

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/grafana-tools/sdk"
)

func PullDashboard(grafanaURL string, apiKey string, directory string, tag string, folderID int) error {
	var (
		boardLinks []sdk.FoundBoard
		rawBoard   sdk.Board
		meta       sdk.BoardProperties
		err        error
	)

	ctx := context.Background()
	c, err := sdk.NewClient(grafanaURL, apiKey, httpClient)

	if err != nil {
		return err
	}

	searchParams := []sdk.SearchParam{sdk.SearchType(sdk.SearchTypeDashboard)}
	if folderID != -1 {
		searchParams = append(searchParams, sdk.SearchFolderID(folderID))
	}

	if tag != "" {
		searchParams = append(searchParams, sdk.SearchTag(tag))
	}

	if boardLinks, err = c.Search(ctx, searchParams...); err != nil {
		return err
	}

	for _, link := range boardLinks {
		if rawBoard, meta, err = c.GetDashboardByUID(ctx, link.UID); err != nil {
			log.Printf("%s for %s\n", err, link.URI)
			ExecutionErrorHappened = true
			continue
		}
		rawBoard.ID = 0
		b, err := json.MarshalIndent(rawBoard, "", "  ")
		if err != nil {
			return err
		}
		if err = writeToFile(directory, b, meta.Slug, tag); err != nil {
			return err
		}
	}
	return nil
}

func PushDashboard(grafanaURL string, apiKey string, directory string, folderId int) error {
	var (
		filesInDir []os.FileInfo
		rawBoard   []byte
		err        error
	)

	ctx := context.Background()
	c, err := sdk.NewClient(grafanaURL, apiKey, httpClient)
	if err != nil {
		return err
	}

	if filesInDir, err = ioutil.ReadDir(directory); err != nil {
		return err
	}
	for _, file := range filesInDir {
		if filepath.Ext(file.Name()) == ".json" {
			if rawBoard, err = ioutil.ReadFile(filepath.Join(directory, file.Name())); err != nil {
				log.Println(err)
				ExecutionErrorHappened = true
				continue
			}
			var board sdk.Board
			if err = json.Unmarshal(rawBoard, &board); err != nil {
				log.Println(err)
				ExecutionErrorHappened = true
				continue
			}
			params := sdk.SetDashboardParams{
				FolderID:  folderId,
				Overwrite: true,
			}
			if _, err := c.SetDashboard(ctx, board, params); err != nil {
				log.Printf("error on importing dashboard %s", board.Title)
				ExecutionErrorHappened = true
				continue
			}
		}
	}
	return nil
}
