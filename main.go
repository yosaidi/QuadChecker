package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// Function to capture the output of a command
func runCommand(command string, args ...string) (string, error) {
	cmd := exec.Command(command, args...)
	var out bytes.Buffer
	cmd.Stdout = &out
	var errOut bytes.Buffer
	cmd.Stderr = &errOut
	err := cmd.Run()
	if err != nil {
		return "", fmt.Errorf("error running command: %w, stderr: %s", err, errOut.String())
	}
	return out.String(), nil
}

// Function to compare outputs
func compareOutput(output, expectedOutput string) bool {
	return strings.TrimSpace(output) == strings.TrimSpace(expectedOutput)
}

// Function to extract dimensions from captured output
func extractDimensions(output string) (string, string) {
	lines := strings.Split(strings.TrimSpace(output), "\n")
	if len(lines) > 0 {
		width := len(lines[0]) // width is the length of the first line
		height := len(lines)   // height is the number of lines
		return fmt.Sprintf("%d", width), fmt.Sprintf("%d", height)
	}
	return "", ""
}

func main() {
	var outputBuffer bytes.Buffer

	// Read from stdin
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		outputBuffer.WriteString(line + "\n")
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading input:", err)
		return
	}
	capturedOutput := strings.TrimSpace(outputBuffer.String())

	// Extract dimensions
	width, height := extractDimensions(capturedOutput)
	if width == "" || height == "" {
		fmt.Println("Failed to extract dimensions from the captured output.")
		return
	}

	// List of commands to run
	commands := []struct {
		Name string
		Cmd  string
	}{
		{"quadA", "./quadA"},
		{"quadB", "./quadB"},
		{"quadC", "./quadC"},
		{"quadD", "./quadD"},
		{"quadE", "./quadE"},
	}

	var results []string

	// Run each command and compare outputs
	for _, cmd := range commands {
		output, err := runCommand(cmd.Cmd, width, height)
		if err != nil {
			fmt.Printf("Error running %s: %v\n", cmd.Name, err)
			continue
		}

		if compareOutput(output, capturedOutput) {
			results = append(results, fmt.Sprintf("[%s] [%s] [%s]", cmd.Name, width, height))
		}
	}

	if len(results) > 0 {
		fmt.Println(strings.Join(results, " || "))
	} else {
		fmt.Println("Not a quad function")
	}
}
