package main

import (
	"fmt"
	"os/exec"
)

func verify() {
	cmd := exec.Command("go", "version")
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error executing command:", err)
		return
	}

	fmt.Println("Go is installed correctly!")
	fmt.Println(string(output))
}
