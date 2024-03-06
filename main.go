package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

func main() {
	rootDir := os.Args[1]
	fi, err := os.Stat(rootDir)
	if err != nil {
		panic(err)
	}

	if !fi.IsDir() {
		panic(fmt.Errorf("%s is not a directory", rootDir))
	}

	entries, err := os.ReadDir(rootDir)
	if err != nil {
		panic(err)
	}

	entryMap := map[string]int{}

	for _, entry := range entries {
		if entry.IsDir() {
			count, err := countFile(filepath.Join(rootDir, entry.Name()))
			if err != nil {
				panic(err)
			}
			entryMap[entry.Name()] = count
		}
	}

	for k, v := range entryMap {
		if v > 1500 {
			fmt.Println(k, v)
		}
	}

}

func countFile(dir string) (int, error) {
	count := 0

	err := filepath.Walk(dir, func(path string, info fs.FileInfo, err error) error {
		if !info.IsDir() {
			count = count + 1
		}
		return nil
	})

	if err != nil {
		return count, err
	}

	return count, nil

}
