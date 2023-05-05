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

func PullNotifications(grafanaURL string, apiKey string, directory string) error {
	var (
		notifications []sdk.AlertNotification
		err           error
	)
	ctx := context.Background()

	c, err := sdk.NewClient(grafanaURL, apiKey, httpClient)
	if err != nil {
		return err
	}

	if notifications, err = c.GetAllAlertNotifications(ctx); err != nil {
		return err
	}

	for _, notification := range notifications {
		b, err := json.MarshalIndent(notification, "", "  ")
		if err != nil {
			return err
		}
		if err = writeToFile(directory, b, notification.Name, ""); err != nil {
			return err
		}
	}
	return nil
}

func PushNotification(grafanaURL string, apiKey string, directory string) error {
	var (
		filesInDir []os.FileInfo
		rawFolder  []byte
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
			if rawFolder, err = ioutil.ReadFile(filepath.Join(directory, file.Name())); err != nil {
				log.Println(err)
				ExecutionErrorHappened = true
				continue
			}
			var notification sdk.AlertNotification
			if err = json.Unmarshal(rawFolder, &notification); err != nil {
				log.Println(err)
				ExecutionErrorHappened = true
				continue
			}
			if _, err := c.CreateAlertNotification(ctx, notification); err != nil {
				log.Printf("error on importing notification %s", notification.Name)
				ExecutionErrorHappened = true
				continue
			}
		}
	}
	return nil
}
