package main

import (
	"fmt"
	"net/http"
	"os/exec"
)

func machineInfoHandler(w http.ResponseWriter, r *http.Request) {
	// Define the command to execute
	cmd := exec.Command("uname", "-a")

	// Run the command and capture its output
	output, err := cmd.Output()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Fprintf(w, "machine-info: %s\n", output)
}

func main() {
	http.HandleFunc("/machine-info", machineInfoHandler)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
