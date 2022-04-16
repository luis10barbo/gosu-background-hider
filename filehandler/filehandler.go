package filehandler

import (
	"os"

	"github.com/luis10barbo/OsuBackgroundRemover/logger"
)

func IsDir(path string) (bool, error) {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return false, err
	}

	return fileInfo.IsDir(), nil
}

func ListDirectory(path string) ([]string, error) {
	var files []string
	pathIsDir, err := IsDir(path)
	if err != nil{
		return files, err
	}

	if !pathIsDir {
		return files, nil
	}

	osFile, err := os.Open(path)
	if err != nil {
		return files, err
	}
	defer osFile.Close()
	logger.InfoLog("Opening folder " + path)

	fileInfo, err := osFile.Readdir(0)
	if err != nil {
		return files, err
	}

	for _, file := range fileInfo {
		files = append(files, file.Name())
	}
	return files, nil
}