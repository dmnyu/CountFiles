package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/nyudlts/bytemath"
)

type ProblemER struct {
	Dir   string
	Count int
	Size  int64
}

var (
	ProblemERs       = []ProblemER{}
	totalCount int   = 0
	totalSize  int64 = 0
	rootDir    string
)

func main() {
	rootDir = os.Args[1]
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

	for _, entry := range entries {
		scan(entry)
	}

	fmt.Println("Total Count:", totalCount)
	fmt.Println("Total Size:", bytemath.ConvertBytesToHumanReadable(totalSize))

	if len(ProblemERs) > 0 {
		fmt.Println("Problem ERs")
		for _, er := range ProblemERs {
			fmt.Printf("%s count: %d size: %d\n", er.Dir, er.Count, er.Size, bytemath.ConvertBytesToHumanReadable(size))
		}
	} else {
		fmt.Println("No Problem ERs Found")
	}

}

func scan(entry fs.DirEntry) {

	if entry.IsDir() {
		fmt.Print("scanning ", entry.Name(), ": ")
		count, size, err := countFile(filepath.Join(rootDir, entry.Name()))
		if err != nil {
			panic(err)
		}

		totalCount = totalCount + count
		totalSize = totalSize + size

		if count >= 2000 || size > 214748364800 {
			ProblemERs = append(ProblemERs, ProblemER{entry.Name(), count, size})
		}

		fmt.Println("count:", count, "size:", bytemath.ConvertBytesToHumanReadable(size))
	}
}

func countFile(dir string) (int, int64, error) {
	count := 0
	var size int64 = 0

	err := filepath.Walk(dir, func(path string, info fs.FileInfo, err error) error {
		if !info.IsDir() {
			count = count + 1
			size = size + info.Size()

		}
		return nil
	})

	if err != nil {
		return count, size, err
	}

	return count, size, nil

}
