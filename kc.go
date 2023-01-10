package main

import (
	"fmt"
	"log"
	"os"

	"github.com/atotto/clipboard"
)

func main() {
	clusterDir := getClusterDirectory()
	files := getFilesFromDirectory(clusterDir)

	for k, file := range files {
		fmt.Printf("%d: %v\n", k+1, file)
	}

	var fileKey int
	fmt.Print("Choose a file: ")
	fmt.Scanln(&fileKey)

	command := "export KUBECONFIG=" + clusterDir + files[fileKey-1]
	fmt.Println(command)
	clipboard.WriteAll(command)
}

func getFilesFromDirectory(dir string) []string {
	filesInDir, err := os.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}

	var files []string

	for _, file := range filesInDir {
		if !file.IsDir() {
			files = append(files, file.Name())
		}
	}

	return files
}

func getClusterDirectory() string {
	dirname, err := (os.UserHomeDir())
	if err != nil {
		log.Fatal(err)
	}

	return dirname + "/.kube/kc/"
}
