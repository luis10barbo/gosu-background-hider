package main

import (
	"fmt"
	"log"
	"os"
	"path"
	"strings"

	"github.com/luis10barbo/OsuBackgroundRemover/logger"
	"github.com/luis10barbo/OsuBackgroundRemover/settings"
)

func ListDirectory(osuPath string) ([]string, error){
	var files []string

	osFile, err := os.Open(osuPath)
	if err != nil {
		return files, err
	} 
	defer osFile.Close()

	fileInfo, err := osFile.Readdir(0)
	if err != nil {
		return files, err
	}

	for _, file := range fileInfo {
		files = append(files, file.Name())
	}
	return files, nil
}

func main() {
	fmt.Println("Starting App...")
	
	configFilePath := "settings.json"

	config, err := settings.LoadSettings(configFilePath)
	if err != nil{
		logger.ErrorLog(err)
	}

	// Check if folder is actually an osu! folder to prevent accidents
	isOsuFolder := false
	osuPath , err := ListDirectory(config.OsuPath)
	if err != nil {
		logger.ErrorLog(err)
	}
	for _, file := range osuPath {
			if file == "osu!.exe" {
					isOsuFolder = true
			}
	}
	if !isOsuFolder {
			logger.ErrorLog(fmt.Errorf("OsuPath value isn't an osu path"))
	}

	// Get "Songs" folder
	songsPath := path.Join(config.OsuPath, "Songs")
	songsFiles, err := ListDirectory(songsPath)
	if err != nil{
		logger.ErrorLog(err)
	}

	// files from songs folder
	for _, songFileName := range songsFiles {
		filePath := path.Join(songsPath, songFileName)
		logger.InfoLog("Opening folder " + filePath)
		files, err := ListDirectory(filePath)

		if err != nil {
			log.Println(err)
		}

		if config.RemoveBackgrounds == 1 {
			imageFormats := []string {
				".jpg",
				".jpeg",
				".png",
			}
			// Loop through files
			for _, fileName := range files {
				// Check if files end at any .imageFormat
				for _, imageFormat := range imageFormats{
					if strings.HasSuffix(fileName, imageFormat) {
						oldFilePath := path.Join(filePath, fileName)
						newFilePath := path.Join(filePath, fileName + "removed")
						
						os.Rename(oldFilePath, newFilePath)
						logger.InfoLog(oldFilePath, "has been changed to", newFilePath)
						fmt.Println()
					}
				}
			}
		} else {
			// Loop through files
			for _, fileName := range files {
				// Check if files ends with removed
				if strings.HasSuffix(fileName, "removed") {
					oldFilePath := path.Join(filePath, fileName)
					newFilePath := path.Join(filePath, strings.ReplaceAll(fileName, "removed", ""))

					os.Rename(oldFilePath, newFilePath)
					logger.InfoLog(oldFilePath, "has been changed to", newFilePath)
				}
			}
		}
	}

	fmt.Println("Ending App...")
}

func SaveSettings() {
	panic("unimplemented")
}
