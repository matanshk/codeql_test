package main

import (
	"os"
	"os/exec"
)

func main() {
	// Get the value of the environment variable
	envVar := os.Getenv("MY_ENV_VAR")

	// Execute the cURL command with the environment variable as data in the POST request
	cmd := exec.Command("curl", "-X", "POST", "-d", envVar, "https://10.5.8.76:80/post")
	err := cmd.Run()

	if err != nil {
		panic(err)
	}
}
