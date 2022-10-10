package utils

import "os"

func OpenFile(file string) string {
	dir, err := os.Getwd()
	CheckError(err)

	completePathFile := buildPath(dir, file)

	data, err := os.ReadFile(completePathFile)
	CheckError(err)

	return string(data)
}
