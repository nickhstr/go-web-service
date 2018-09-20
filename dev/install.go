package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
)

// DevDeps defines the struct for unmarshalling dev
// dependencies data
type DevDeps struct {
	Dependencies map[string]string `json:"devDependencies"`
}

func main() {
	var devDeps DevDeps
	devDepsFile, err := os.Open("devdeps.json")
	if err != nil {
		fmt.Printf("failed to open dependencies file %v\n", err)
	}
	defer devDepsFile.Close()

	depsData, err := ioutil.ReadAll(devDepsFile)
	if err != nil {
		fmt.Printf("failed to read dependencies file %v\n", err)
	}

	if err = json.Unmarshal(depsData, &devDeps); err != nil {
		fmt.Printf("failed to unmarshal json %v\n", err)
	}

	cmdPath, err := exec.LookPath("go")
	if err != nil {
		fmt.Printf("Unable to run 'go' command %v\n", err)
	}

	for _, url := range devDeps.Dependencies {
		fmt.Printf("Installing %s\n", url)
		cmd := exec.Command(cmdPath, "get", url)
		cmd.Env = os.Environ()
		cmd.Env = append(cmd.Env, "GO111MODULE=off")
		if err = cmd.Run(); err != nil {
			fmt.Printf("Failed to install %s %v\n", url, err)
		}
	}
}
