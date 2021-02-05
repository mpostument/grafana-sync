package grafana

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

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
	fileName = fmt.Sprintf("%s/%s.json", path, name)
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
