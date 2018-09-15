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

// IsDev indicates if app is in dev env
func IsDev() bool {
	return GetGoEnv() == "development"
}

// IsProd indicates if app is in prod env
func isProd() bool {
	return GetGoEnv() == "production"
}

// GetPort returns an appropriate port for http.ListenAndServe to use
func GetPort() string {
	defaultPort := "3000"
	isDev := IsDev()

	port, isSet := os.LookupEnv("PORT")
	if !isSet {
		port = defaultPort
	}

	if isDev {
		return fmt.Sprintf("localhost:%s", port)
	}

	return fmt.Sprintf(":%s", port)
}
