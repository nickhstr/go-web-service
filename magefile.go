// +build mage

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

var Default = Install

const coverageOut = "coverage.out"
const binOutput = "./bin/app"

// Builds the service's executable.
func Build() error {
	fmt.Println("🛠️  Building executable...")
	err := sh.RunV("go", "build", "-o", binOutput, "./main.go")
	fmt.Println("👍 Done.")

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

// Starts the app in dev mode.
func Dev() error {
	fmt.Println("🚀 Starting dev server...")
	return sh.RunV("modd", "--file=./internal/tools/modd.dev.conf")
}

// Installs all dependencies.
func Install() error {
	var err error

	fmt.Println("downloading dependencies")
	err = sh.RunV("go", "mod", "download")
	if err != nil {
		return err
	}

	toolDeps, err := toolsToInstall("./internal/tools/tools.go")
	if err != nil {
		return err
	}

	for _, dep := range toolDeps {
		fmt.Printf("installing %s\n", dep)
		err = sh.RunV("go", "install", dep)
		if err != nil {
			return err
		}
	}

	fmt.Println("👍 Done.")

	return err
}

// Lints all files.
func Lint() error {
	var err error

	err = sh.RunV("golangci-lint", "run")
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
	var err error

	fmt.Println("🏃 Running all Go tests...")
	// Set verbose env var to get test output
	// Needed until next release of Mage
	os.Setenv(mg.VerboseEnv, "true")
	env := map[string]string{
		"GO_ENV": "test",
	}
	err = sh.RunWith(env, "modd", "--file=./internal/tools/modd.test.conf")
	if err != nil {
		return err
	}

	fmt.Println("✅ Done.")

	return err
}

// toolsToInstall gets the Go tools to install from the tools.go file.
func toolsToInstall(toolsPath string) ([]string, error) {
	var (
		tools  []string
		err    error
		done   bool
		cutset = "\t\n_ \""
	)

	path, err := filepath.Abs(toolsPath)
	if err != nil {
		fmt.Printf("filepath.Abs error: %v\n", err)
		return tools, err
	}
	f, err := os.Open(path)
	if err != nil {
		fmt.Printf("os.Open error: %v\n", err)
		return tools, err
	}

	reader := bufio.NewReader(f)

	for !done {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			done = true
		} else if err != nil {
			fmt.Println(err)
			done = true
			break
		}

		// Target only lines with the name of the dependency, excluding Mage
		if strings.Contains(line, `_ "`) && !strings.Contains(line, "magefile/mage") {
			dep := strings.Trim(line, cutset)
			tools = append(tools, dep)
		}
	}

	return tools, err
}
