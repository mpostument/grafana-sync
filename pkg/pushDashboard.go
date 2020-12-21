package pkg

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

func PushDashboard(grafanaURL string, apiKey string, directory string) {
	var (
		filesInDir []os.FileInfo
		rawBoard   []byte
		err        error
	)

	ctx := context.Background()
	c := sdk.NewClient(grafanaURL, apiKey, sdk.DefaultHTTPClient)
	filesInDir, err = ioutil.ReadDir(directory)
	if err != nil {
		log.Fatal(err)
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
			_, err := c.SetDashboard(ctx, board, params)
			if err != nil {
				log.Printf("error on importing dashboard %s", board.Title)
				continue
			}
		}
	}
}
