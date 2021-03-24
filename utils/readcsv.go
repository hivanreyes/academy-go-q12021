package utils

import (
	"bufio"
	"encoding/csv"
	"io"
	"os"
)

func ReadCsv(filename string) ([][]string, error) {

	// Open CSV file
	file, err := os.Open(filename)
	if err != nil {
		return [][]string{}, err
	}
	defer file.Close()

	// Skip first row (line)
	row1, err := bufio.NewReader(file).ReadSlice('\n')
	if err != nil {
		return nil, err
	}

	_, err = file.Seek(int64(len(row1)), io.SeekStart)
	if err != nil {
		return nil, err
	}

	// Read File into a Variable
	lines, err := csv.NewReader(file).ReadAll()
	if err != nil {
		return [][]string{}, err
	}

	return lines, nil
}
