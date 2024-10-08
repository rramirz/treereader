package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// TreeReader recursively opens and reads files in a directory while preserving the directory hierarchy
func TreeReader(rootDir string, ignoreFilenames, selectFilenames map[string]struct{}) {
	err := filepath.Walk(rootDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Skip .git directory and its contents
		if info.IsDir() && info.Name() == ".git" {
			return filepath.SkipDir
		}

		relativePath, err := filepath.Rel(rootDir, path)
		if err != nil {
			return err
		}

		if !info.IsDir() {
			if info.Mode()&0o111 != 0 {
				return nil
			}
			if _, ok := ignoreFilenames[relativePath]; ok {
				return nil
			}
			if len(selectFilenames) > 0 {
				if _, ok := selectFilenames[relativePath]; !ok {
					return nil
				}
			}
			content, err := os.ReadFile(path)
			if err != nil {
				return err
			}
			fmt.Printf("Opening file: %s\n", path)
			fmt.Println(string(content))
		}
		return nil
	})

	if err != nil {
		log.Fatalf("Error walking the directory: %v", err)
	}
}

func main() {
	// Parse command-line arguments manually
	ignoreList := ""
	selectList := ""
	args := os.Args[1:]
	rootDir := "."

	for i, arg := range args {
		if strings.HasPrefix(arg, "-ignore=") {
			ignoreList = strings.TrimPrefix(arg, "-ignore=")
			args = append(args[:i], args[i+1:]...)
			break
		}
	}

	for i, arg := range args {
		if strings.HasPrefix(arg, "-select=") {
			selectList = strings.TrimPrefix(arg, "-select=")
			args = append(args[:i], args[i+1:]...)
			break
		}
	}

	if len(args) > 0 {
		rootDir = args[0]
	}

	// Initialize the map for ignored filenames
	ignoreFilenames := make(map[string]struct{})

	// Split the ignore list into individual filenames and add them to the map
	if ignoreList != "" {
		for _, name := range strings.Split(ignoreList, ",") {
			trimmedName := strings.TrimSpace(name)
			ignoreFilenames[trimmedName] = struct{}{}
		}
	}

	// Initialize the map for selected filenames
	selectFilenames := make(map[string]struct{})

	// Split the select list into individual filenames and add them to the map
	if selectList != "" {
		for _, name := range strings.Split(selectList, ",") {
			trimmedName := strings.TrimSpace(name)
			selectFilenames[trimmedName] = struct{}{}
		}
	}

	// Run TreeReader
	TreeReader(rootDir, ignoreFilenames, selectFilenames)
}
