package grafana

import (
	"io/ioutil"
	"os"
	"path/filepath"
)

var ExecutionErrorHappened = false

func writeToFile(directory string, content []byte, name string, tag string) error {
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
			return err
		}
	}
	fileName = filepath.Join(path, name+".json")
	dashboardFile, err = os.Create(fileName)
	if err != nil {
		return err
	}

	defer dashboardFile.Close()

	err = ioutil.WriteFile(dashboardFile.Name(), content, os.FileMode(0755))
	if err != nil {
		return err
	}
	return nil
}
