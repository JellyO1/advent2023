package main

import (
	"bufio"
	"os"
)

func inputFromFile01(path string) ([]string, error) {
	f, err := os.Open(path)
	defer f.Close()
	if err != nil {
		return nil, err
	}

	result := make([]string, 0)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		result = append(result, scanner.Text())
	}

	if err = scanner.Err(); err != nil {
		return result, err
	}

	return result, nil
}
