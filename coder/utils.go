package coder

import (
	"bufio"
	"io"
	"os"
)

// GetSequence получить входные данные из файла
func GetSequence(inputPath string) ([]byte, error) {
	file, err := os.Open(inputPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	var inputSeq = make([]byte, 0)
	for {
		b, err := reader.ReadByte()
		if err == io.EOF {
			break
		}
		inputSeq = append(inputSeq, b)
	}
	return inputSeq, nil
}

// SaveSequence результат в файл
func SaveSequence(outputPath string, outputSeq []byte) error {
	file, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)

	writer.Write(outputSeq)
	writer.Flush()

	return nil
}
