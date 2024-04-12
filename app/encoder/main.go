package main

import (
	"bufio"
	"bw-coding/coder"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

var isDir bool

const CHUNK_SIZE = 2 << 11

func main() {
	if len(os.Args) < 3 {
		fmt.Println("usage: encoder <input path> <output path>")
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

	for i := 0; i < len(inFiles); i++ {
		// разбить входной файл на чанки
		var input, output *os.File
		input, err = os.Open(inFiles[i])
		if err != nil {
			fmt.Printf("error opening input file: %s\n", err)
			os.Exit(1)
		}
		var outPath = os.Args[2]
		if isDir {
			outPath = filepath.Join(os.Args[2], strings.Split(inFiles[i], os.Args[1])[1])
		}
		output, err = os.Create(outPath)
		if err != nil {
			fmt.Printf("error creating output file: %s\n", err)
			os.Exit(1)
		}
		reader := bufio.NewReader(input)
		writer := bufio.NewWriter(output)
		var chunk = make([]byte, CHUNK_SIZE)
		writer.WriteString(strconv.Itoa(CHUNK_SIZE))

		for {
			var n, slen int
			slen, err = reader.Read(chunk)
			if err == io.EOF {
				break
			}
			var lcol = make([]byte, slen)
			n = coder.Encode(chunk, lcol, slen)
			writer.WriteString(strconv.Itoa(n))
			writer.Write(lcol)
		}
		writer.Flush()
		input.Close()
		output.Close()
	}
}
