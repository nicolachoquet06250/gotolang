package main

import "os"

func openFile(file string) string {
	dir, err := os.Getwd()
	checkError(err)

	completePathFile := buildPath(dir, file)

	data, err := os.ReadFile(completePathFile)
	checkError(err)

	return string(data)
}
