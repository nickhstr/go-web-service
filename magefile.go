//+build mage

package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

const outputDir = "bin"
const coverageOut = "coverage.out"

var projectName string

// Uses 'go build' to create the application executable in the 'bin' directory.
func Build() error {
	var err error

	mg.Deps(Project)

	sh.RunV("echo", "🛠️  Building executable...")
	err = sh.Run("go", "build", "-o", filepath.Join(outputDir, projectName), "main.go")
	if err != nil {
		return err
	}

	return sh.RunV("echo", "👍 Done.")
}

// Removes build artifacts.
func Clean() error {
	var err error

	sh.RunV("echo", "🔥 Removing build artifacts...")
	err = os.RemoveAll(outputDir)
	if err != nil {
		return err
	}

	return sh.RunV("echo", "✨ Done.")
}

// Runs all tests and reports coverage.
func Coverage() error {
	var err error

	mg.Deps(CreateCoverage)
	sh.RunV("echo", "========== Coverage Summary ==========")
	err = sh.RunV("go", "tool", "cover", "-func", coverageOut)
	if err != nil {
		return err
	}

	return sh.RunV("echo", "======================================")
}

// Opens coverage report in a browser.
func CoverageHtml() error {
	mg.Deps(CreateCoverage)
	sh.RunV("echo", "🛠  Opening coverage report in browser...")

	return sh.Run("go", "tool", "cover", "-html", coverageOut)
}

// Runs all tests, and outputs a coverage report.
func CreateCoverage() error {
	sh.RunV("echo", "🏃 Running tests and creating coverage report...")
	os.Setenv(mg.VerboseEnv, "true")
	env := map[string]string{
		"GO_ENV": "test",
	}
	sh.RunWith(env, "go", "test", "-race", "-coverprofile", coverageOut, "./...")

	return sh.RunV("echo", "👍 Done.")
}

// Starts the app in dev mode.
func Dev() error {
	sh.Run("echo", "🚀 Starting dev server...")

	return sh.RunV("modd", "-f", "./modd.conf")
}

// Installs all dependencies.
func Install() error {
	var err error

	sh.RunV("echo", "🛠  Installing package dependencies...")
	err = sh.RunV("go", "mod", "download")
	if err != nil {
		return err
	}

	devDeps := []string{
		"github.com/cortesi/modd/cmd/modd",
	}

	for _, dep := range devDeps {
		err = sh.RunV("go", "install", dep)
		if err != nil {
			return err
		}
	}

	return sh.RunV("echo", "👍 Done.")
}

// Lints all files.
func Lint() error {
	var err error

	sh.RunV("echo", "🔍  Linting files...")
	err = sh.RunV("golangci-lint", "run")
	if err != nil {
		return err
	}

	return sh.RunV("echo", "👍 Done.")
}

// Gets the base path from the current working directory, used to set the project name.
func Project() {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	projectName = filepath.Base(dir)
	log.Println(projectName)
}

// Starts the service.
func Serve() error {
	mg.Deps(Project, Build)
	sh.RunV("echo", "🚀 Starting server...")

	return sh.RunV("./" + filepath.Join(outputDir, projectName))
}

// Runs all tests.
func Test() error {
	var err error

	sh.RunV("echo", "🏃 Running all Go tests...")
	os.Setenv(mg.VerboseEnv, "true")
	env := map[string]string{
		"GO_ENV": "test",
	}
	err = sh.RunWith(env, "go", "test", "-race", "./...")
	if err != nil {
		return err
	}

	return sh.RunV("echo", "👍 Done.")
}
