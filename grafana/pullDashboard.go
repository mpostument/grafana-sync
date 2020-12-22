package grafana

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/grafana-tools/sdk"
)

func PullDashboard(grafanaURL string, apiKey string, directory string, tag string) {
	var (
		boardLinks []sdk.FoundBoard
		rawBoard   sdk.Board
		meta       sdk.BoardProperties
		err        error
	)
	ctx := context.Background()

	c := sdk.NewClient(grafanaURL, apiKey, sdk.DefaultHTTPClient)

	if boardLinks, err = c.Search(ctx, sdk.SearchTag(tag)); err != nil {
		log.Fatalln(err)
	}

	for _, link := range boardLinks {
		if rawBoard, meta, err = c.GetDashboardByUID(ctx, link.UID); err != nil {
			log.Printf("%s for %s\n", err, link.URI)
			continue
		}
		rawBoard.ID = 0
		writeDashboardToFile(directory, rawBoard, meta.Slug, tag)
	}
}

func writeDashboardToFile(directory string, dashboard sdk.Board, name string, tag string) {
	var (
		err           error
		path          string
		dashboardFile *os.File
		fileName      string
	)

	path = directory
	if tag != "" {
		path = filepath.Join(path, tag)
	}

	if _, err = os.Stat(path); os.IsNotExist(err) {
		if err = os.MkdirAll(path, 0755); err != nil {
			log.Fatalln("Directory is not created", err)
		}
	}
	fileName = fmt.Sprintf("%s/%s.json", path, name)
	dashboardFile, err = os.Create(fileName)
	if err != nil {
		log.Printf("failed to create file for dashboard %s", fileName)
	}

	defer dashboardFile.Close()

	err = json.NewEncoder(dashboardFile).Encode(dashboard)
	if err != nil {
		log.Printf("failed to encode dashboard json to file %s", fileName)
	}
}
