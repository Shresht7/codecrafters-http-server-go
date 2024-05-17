package main

import "os"

// getDirectoryFromArguments extracts the directory from the command line arguments
func GetDirectoryFromArguments() string {
	// Get Command Line Arguments
	args := os.Args[1:]

	// Extract --directory flag
	directory := ""

	i := 0
	for i < len(args) {
		if args[i] == "--directory" {
			if i+1 < len(args) {
				directory = args[i+1]
				i++ // Skip the next argument as we just used it as the value
			}
		}
		i++
	}

	return directory
}
