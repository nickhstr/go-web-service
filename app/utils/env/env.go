package env

import (
	"fmt"
	"os"
)

// Get provides a way to get the value of a supplied environment
// variable. If it is not found, the optionally supplied default
// value is returned.
func Get(envVar string, defaultVal string) string {
	if val, isSet := os.LookupEnv(envVar); isSet {
		return val
	}

	return defaultVal
}

// IsDev indicates if app is in dev env.
func IsDev() bool {
	return Get("GO_ENV", "development") == "development"
}

// IsProd indicates if app is in prod env.
func isProd() bool {
	return Get("GO_ENV", "development") == "production"
}

// GetPort returns an appropriate port for http.ListenAndServe to use.
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
