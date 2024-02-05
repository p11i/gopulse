// pkg/analyzer/analyzer.go

package analyzer

import (
	"fmt"
	"log"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
)

// ExecuteCommand executes a system command and returns the output
func ExecuteCommand(command string, args ...string) (string, error) {
	cmdOutput, err := execCommand(command, args...)
	if err != nil {
		log.Printf("Error executing '%s' command: %v", command, err)
		return "", fmt.Errorf("failed to execute command: %s", command)
	}
	return cmdOutput, nil
}

// execCommand is a helper function to execute a system command and return the output
func execCommand(command string, args ...string) (string, error) {
	cmd := exec.Command(command, args...)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}
	return string(output), nil
}

// GetProcesses retrieves the list of processes
func GetProcesses() (string, error) {
	output, err := ExecuteCommand("ps", "aux")
	if err != nil {
		return "", fmt.Errorf("failed to retrieve processes: %v", err)
	}
	return output, nil
}

// AnalyzeProcesses analyzes the processes and identifies resource-intensive ones
func AnalyzeProcesses(processesOutput string) {
	lines := strings.Split(processesOutput, "\n")

	// Process information extraction using regular expressions
	processInfoRegex := regexp.MustCompile(`\s*(\d+)\s+(\S+)\s+(\S+)\s+(\S+)\s+(\S+)\s+(\S+)\s+(\S+)\s+(\S+)\s+(\S+)\s+(.+)`)

	// Header line, assuming the first line is the header
	headerLine := lines[0]
	headers := strings.Fields(headerLine)

	// Map to store resource values for each process
	processResourceMap := make(map[string]map[string]string)

	// Iterate over each line (excluding the header)
	for _, line := range lines[1:] {
		matches := processInfoRegex.FindStringSubmatch(line)
		if len(matches) != len(headers)+1 {
			continue // Skip lines that don't match the expected format
		}

		processData := make(map[string]string)

		// Extract data for each column
		for i, header := range headers {
			processData[header] = matches[i+1]
		}

		// Store data in the map
		processResourceMap[processData["COMMAND"]] = processData
	}

	// Identify and display resource-intensive processes
	fmt.Println("Resource-intensive processes:")
	for command, data := range processResourceMap {
		cpuUsage, _ := strconv.ParseFloat(data["%CPU"], 64)
		memUsage, _ := strconv.ParseFloat(data["%MEM"], 64)

		// Example: Define thresholds for considering a process resource-intensive
		cpuThreshold := 5.0
		memThreshold := 5.0

		if cpuUsage > cpuThreshold || memUsage > memThreshold {
			fmt.Printf("Command: %s\n", command)
			fmt.Printf("CPU Usage: %.2f%%\n", cpuUsage)
			fmt.Printf("Memory Usage: %.2f%%\n", memUsage)
			fmt.Println("----------")
		}
	}
}
