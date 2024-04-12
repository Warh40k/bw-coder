package main

import (
	"bw-coding/coder"
	"fmt"
	"io"
	"io/fs"
	"math"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

var isDir bool

// TODO: подумать, как передавать размер чанка (необяз)
const CHUNK_SIZE = 4096

func main() {
	if len(os.Args) < 3 {
		fmt.Println("usage: bwencoder <input path> <output path>")
		os.Exit(1)
	}

	info, err := os.Stat(os.Args[1])
	if err != nil {
		fmt.Printf("error opening input file: %s\n", err)
		os.Exit(1)
	}
	var inFiles []string
	var dirs []string
	if info.IsDir() {
		isDir = true
		err = filepath.WalkDir(os.Args[1], func(path string, d fs.DirEntry, err error) error {
			if d.IsDir() {
				dirs = append(dirs, path)
			} else {
				inFiles = append(inFiles, path)
			}
			return nil
		})
		if err != nil {
			fmt.Printf("error getting dir files: %s\n", err)
			os.Exit(1)
		}
	} else {
		inFiles = append(inFiles, os.Args[1])
	}

	info, err = os.Stat(os.Args[2])
	if os.IsNotExist(err) {
		err = os.MkdirAll(filepath.Dir(os.Args[2]), 0777)
		if err != nil {
			fmt.Printf("error creating output directory: %s\n", err)
			os.Exit(1)
		}
	}
	for _, dir := range dirs {
		err = os.MkdirAll(filepath.Join(os.Args[2], strings.Split(dir, os.Args[1])[1]), 0777)
		if err != nil {
			fmt.Printf("error creating output directory: %s\n", err)
			os.Exit(1)
		}
	}
	wg := sync.WaitGroup{}
	for i := 0; i < len(inFiles); i++ {
		wg.Add(1)
		go processFile(inFiles[i], &wg)
	}
	wg.Wait()
}

func processFile(path string, wg *sync.WaitGroup) {
	defer wg.Done()
	var input, output *os.File
	var err error
	input, err = os.Open(path)
	if err != nil {
		fmt.Printf("error opening input file: %s\n", err)
		os.Exit(1)
	}
	defer input.Close()

	var outPath = os.Args[2]
	if isDir {
		outPath = filepath.Join(os.Args[2], strings.Split(path, os.Args[1])[1])
	}
	output, err = os.Create(outPath)
	if err != nil {
		fmt.Printf("error creating output file: %s\n", err)
		os.Exit(1)
	}
	defer output.Close()

	var chunk = make([]byte, CHUNK_SIZE) // чанк (в байтах)
	var bitCount = int(math.Ceil(math.Log2(float64(CHUNK_SIZE))))

	for {
		var n, slen int
		slen, err = input.Read(chunk)
		if err == io.EOF {
			break
		}
		var lcol = make([]byte, slen)
		n = coder.Encode(chunk, lcol, slen)
		bnum := getBin(n, bitCount)
		output.WriteString(bnum)
		output.Write(lcol)
	}
}

func getBin(num, bitCount int) string {
	var numBit = 1
	if num != 0 {
		numBit = int(math.Log2(float64(num))) + 1
	}
	zeroCount := bitCount - numBit
	return strings.Repeat("0", zeroCount) + fmt.Sprintf("%b", num)
}
