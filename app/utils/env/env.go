package env

import (
	"fmt"
	"os"
)

// Get provides a way to get the value of a supplied environment
// variable. If it is not found, the optionally supplied default
// value is returned.
func Get(envVar string, defaultVal ...string) string {
	if val, isSet := os.LookupEnv(envVar); isSet {
		return val
	}

	if len(defaultVal) > 0 {
		return defaultVal[0]
	}

	return ""
}

// IsDev indicates if app is in dev env.
func IsDev() bool {
	return Get("GO_ENV", "development") == "development"
}

// IsProd indicates if app is in prod env.
func IsProd() bool {
	return Get("GO_ENV", "development") == "production"
}

// GetAddr returns an appropriate address for http.ListenAndServe to use.
func GetAddr() string {
	isDev := IsDev()

	port := Get("PORT", "3000")

	if isDev {
		return fmt.Sprintf("localhost:%s", port)
	}

	return fmt.Sprintf("0.0.0.0:%s", port)
}
