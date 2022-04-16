package main

import (
	"fmt"
	"log"
	"os"
	"path"
	"strings"

	"github.com/luis10barbo/OsuBackgroundRemover/filehandler"
	"github.com/luis10barbo/OsuBackgroundRemover/logger"
	"github.com/luis10barbo/OsuBackgroundRemover/settings"

	"github.com/superhawk610/bar"
)

func main() {
	fmt.Println("Starting App...")
	
	// Open settings
	config, err := settings.LoadSettings()
	if err != nil{
		logger.FatalLog(err)
	}

	// Check if folder is actually an osu! folder to prevent accidents
	isOsuFolder := false
	osuPath , err := filehandler.ListDirectory(config.OsuPath)
	if err != nil {
		logger.FatalLog(fmt.Sprintf("OsuPath \"%s\" is not a directory! Update your settings.json file...", config.OsuPath))
	}
	for _, file := range osuPath {
			if file == "osu!.exe" {
					isOsuFolder = true
			}
	}
	if !isOsuFolder {
		logger.FatalLog("OsuPath value isn't an osu path")
	}

	// Get "Songs" folder
	songsPath := path.Join(config.OsuPath, "Songs")
	songsFiles, err := filehandler.ListDirectory(songsPath)
	if err != nil{
		logger.ErrorLog(err)
	}

	// files from songs folder
	loopProgress := bar.New(len(songsFiles))
	totalModifiedFiles := 0
	for _, songFileName := range songsFiles {
		filePath := path.Join(songsPath, songFileName)

		files, err := filehandler.ListDirectory(filePath)
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
						totalModifiedFiles++
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
					totalModifiedFiles++
				}
			}
		}
		loopProgress.Tick()
	}
	loopProgress.Done()

	logger.DesktopNotification(fmt.Sprintf("The process has been completed, %d files were modified!", totalModifiedFiles))
	fmt.Println("Ending App...")
}

