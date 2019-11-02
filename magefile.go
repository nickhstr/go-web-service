// +build mage

package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

var Default = Install

const coverageOut = "coverage.out"
const binOutput = "./bin/app"

func ldflags() string {
	gc, _ := sh.Output("git", "rev-parse", "--short", "HEAD")

	flags := []string{
		fmt.Sprintf(`-X "main.GitCommit=%s"`, gc),
	}

	return strings.Join(flags, "")
}

// Builds the app's executable.
func Build() error {
	mg.Deps(Clean)

	fmt.Println("🚧 Building executable...")
	err := sh.RunV("go", "build", "-o", binOutput, "-ldflags="+ldflags(), "main.go")
	fmt.Println("✨ Done.")

	return err
}

// Removes build artifacts.
func Clean() error {
	fmt.Println("🔥 Removing build artifacts...")
	err := sh.Rm("./bin")
	fmt.Println("✨ Done.")

	return err
}

// Runs all tests and reports coverage.
func Coverage() error {
	var err error
	mg.Deps(CreateCoverage)

	fmt.Println("=============================== Coverage summary ===============================")

	err = sh.RunV("go", "tool", "cover", "-func", coverageOut)
	if err != nil {
		return err
	}

	fmt.Println("================================================================================")

	return err
}

// Opens coverage report in a browser.
func CoverageHtml() error {
	mg.Deps(CreateCoverage)

	fmt.Println("📊 Opening coverage report in browser...")

	return sh.Run("go", "tool", "cover", "-html", coverageOut)
}

// Runs all tests, and outputs a coverage report.
func CreateCoverage() error {
	fmt.Println("🏃 Running tests and creating coverage report...")
	os.Setenv(mg.VerboseEnv, "true")
	env := map[string]string{
		"GO_ENV": "test",
	}
	err := sh.RunWith(env, "go", "test", "-race", "-coverprofile", coverageOut, "./...")

	fmt.Println("✅ Done.")

	return err
}

// Starts the app and restarts on changes.
func Dev() error {
	return sh.RunV("go", "run", "vendor/github.com/cortesi/modd/cmd/modd/main.go", "--file=./internal/tools/modd.dev.conf")
}

// Installs all dependencies.
func Install() error {
	var err error

	fmt.Println("📡 Downloading dependencies...")
	err = sh.RunV("go", "mod", "download")
	if err != nil {
		return err
	}

	err = sh.RunV("go", "mod", "vendor")
	fmt.Println("✨ Done.")

	return err
}

// Lints all files.
func Lint() error {
	var err error

	fmt.Println("🔍 Linting files...")
	err = sh.RunV("go", "run", "vendor/github.com/golangci/golangci-lint/cmd/golangci-lint/main.go", "run")
	if err != nil {
		return err
	}

	fmt.Println("✨ Done.")

	return err
}

// Builds and runs the application.
func Serve() error {
	mg.Deps(Build)

	fmt.Println("🚀 Starting server...")
	err := sh.RunV(binOutput)

	return err
}

// Runs all tests.
func Test() error {
	var err error

	fmt.Println("🏃 Running all Go tests...")
	// Set verbose env var to get test output
	// Needed until next release of Mage
	os.Setenv(mg.VerboseEnv, "true")
	env := map[string]string{
		"GO_ENV": "test",
	}
	err = sh.RunWith(env, "go", "test", "-race", "./...")
	if err != nil {
		return err
	}

	fmt.Println("✅ Done.")

	return err
}

// Runs all tests, and watches for changes.
func TestDev() error {
	err := sh.RunV("go", "run", "vendor/github.com/cortesi/modd/cmd/modd/main.go", "--file=./internal/tools/modd.test.conf")

	return err
}
