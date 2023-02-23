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

func PullDatasources(grafanaURL string, apiKey string, directory string) error {
	var (
		datasources []sdk.Datasource
		err         error
	)
	ctx := context.Background()

	c, err := sdk.NewClient(grafanaURL, apiKey, sdk.DefaultHTTPClient)
	if err != nil {
		return err
	}

	if datasources, err = c.GetAllDatasources(ctx); err != nil {
		return err
	}

	for _, datasource := range datasources {
		b, err := json.MarshalIndent(datasource, "", "  ")
		if err != nil {
			return err
		}
		if err = writeToFile(directory, b, datasource.Name, ""); err != nil {
			return err
		}
	}
	return nil
}

func PushDatasources(grafanaURL string, apiKey string, directory string) error {
	var (
		filesInDir []os.FileInfo
		rawFolder  []byte
		err        error
	)

	ctx := context.Background()
	c, err := sdk.NewClient(grafanaURL, apiKey, sdk.DefaultHTTPClient)
	if err != nil {
		return err
	}

	if filesInDir, err = ioutil.ReadDir(directory); err != nil {
		return err
	}
	for _, file := range filesInDir {
		if filepath.Ext(file.Name()) == ".json" {
			if rawFolder, err = ioutil.ReadFile(filepath.Join(directory, file.Name())); err != nil {
				log.Println(err)
				continue
			}
			var datasource sdk.Datasource
			if err = json.Unmarshal(rawFolder, &datasource); err != nil {
				log.Println(err)
				ExecutionErrorHappened = true
				continue
			}
			if _, err := c.CreateDatasource(ctx, datasource); err != nil {
				log.Printf("error on importing folder %s", datasource.Name)
				ExecutionErrorHappened = true
				continue
			}
		}
	}
	return nil
}
