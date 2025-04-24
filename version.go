package main

import (
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	currentVersion := getCurrentVersion()
	version := calculateVersion(currentVersion)
	setCurrentVersion(version)
	log.Println("Version updated successfully: ", version)
}

func getCurrentVersion() string {
	file, err := os.OpenFile("version.txt", os.O_RDWR|os.O_CREATE, 0644)
	if err != nil && err != io.EOF {
		log.Fatalf("Error opening file: %v", err)
	}
	defer file.Close()

	buf := make([]byte, 1024)
	n, err := file.Read(buf)
	if err != nil && err != io.EOF {
		log.Fatalf("Error reading file: %v", err)
	}
	result := string(buf[:n])
	return result
}

func setCurrentVersion(version string) {
	file, err := os.OpenFile("version.txt", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	defer file.Close()

	_, err = file.WriteString(version)
	if err != nil {
		log.Fatalf("Error writing to file: %v", err)
	}
}

func calculateVersion(version string) string {
	res := strings.Split(version, "v")
	if len(res) < 2 {
		log.Fatal("Invalid version format")
	}

	vers := strings.Split(res[1], ".")

	if len(vers) < 3 {
		log.Fatal("Invalid version number")
	}

	major, _ := strconv.Atoi(vers[0])
	minor, _ := strconv.Atoi(vers[1])
	patch, _ := strconv.Atoi(vers[2])

	patch++

	if patch > 99 {
		patch = 0
		minor++
	}

	if minor > 99 {
		minor = 0
		major++
	}

	result := "v" + strconv.Itoa(major) + "." + strconv.Itoa(minor) + "." + strconv.Itoa(patch)

	return result
}
