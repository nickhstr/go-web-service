package utils

import (
	"bufio"
	"errors"
	"io"
	"log"
	"os"
	"regexp"
	"strings"
)

// AppName returns the application's name.
// The application name can be set as an environment variable,
// or it can be read from the go.mod file.
func AppName() string {
	if appName, isSet := os.LookupEnv("APP_NAME"); isSet {
		return appName
	}

	name, err := getAppName()
	if err != nil {
		name = "web-service"
	}

	if err = os.Setenv("APP_NAME", name); err != nil {
		log.Println("Unable to set env var 'APP_NAME'")
	}

	return name
}

func getAppName() (string, error) {
	var name string

	modFile, err := os.Open("./go.mod")
	if err != nil {
		return "", err
	}
	defer modFile.Close()

	matcher, err := regexp.Compile("mod")
	if err != nil {
		return "", err
	}

	reader := bufio.NewReader(modFile)

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}

			return "", err
		}

		if matcher.MatchString(line) {
			split := strings.Split(line, " ")
			if len(split) == 0 {
				return "", errors.New("Cannot find module declaration")
			}

			moduleName := split[1]
			names := strings.Split(moduleName, "/")
			if len(names) == 0 {
				return "", errors.New("Cannot find module name")
			}

			appName := strings.TrimSpace(names[len(names)-1])

			return appName, nil
		}
	}

	return name, nil
}
