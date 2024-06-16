package main

import (
	"os"
	"strings"

	"github.com/sirupsen/logrus"
)

func main() {
	readFile, error := os.ReadFile("/mnt/edisk/Yoma Projects/readme_target/src/main/java/org/yomabank/controller/RestController.java")

	if error != nil {
		logrus.Errorf("Error occurred at reading given directory ::: %v\n", error.Error())
	}

	fetchedFile := string(readFile)

	splittedString := strings.Split(fetchedFile, "\n")

	logrus.Infof("Fetched splitted string ::: %v\n", splittedString)

	for index, value := range splittedString {
		if value == "// OpenDocs" {
			logrus.Infof("Index of OpenDocs ::: %v\n", index)
		}
	}
}
