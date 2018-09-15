package env

import (
	"fmt"
	"os"
)

// GetGoEnv returns the GO_ENV value, if set.
// Otherwise, it returns the default value
func GetGoEnv() string {
	if goEnv, isSet := os.LookupEnv("GO_ENV"); isSet {
		return goEnv
	}

	return "development"
}

// IsDevEnv indicates if app is in dev env
func IsDevEnv() bool {
	return GetGoEnv() == "development"
}

// IsProdEnv indicates if app is in prod env
func isProdEnv() bool {
	return GetGoEnv() == "production"
}

// GetPort returns an appropriate port for http.ListenAndServe to use
func GetPort() string {
	var defaultPort = "3000"
	var isDev = IsDevEnv()

	port, isSet := os.LookupEnv("PORT")
	if !isSet {
		port = defaultPort
	}

	if isDev {
		return fmt.Sprintf("localhost:%s", port)
	}

	return fmt.Sprintf(":%s", port)
}
