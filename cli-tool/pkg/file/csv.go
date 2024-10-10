package file

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"strings"
)

func CSV2Records(f io.Reader) ([][]string, error) {

	lines, err := format2RFC4180(f)
	if err != nil {
		return nil, err
	}
	r := csv.NewReader(strings.NewReader(strings.Join(lines, "\n")))
	records, err := r.ReadAll()
	if err != nil {
		fmt.Println("Error parsing CSV:", err)
		return nil, err
	}

	return records, nil
}

func format2RFC4180(f io.Reader) ([]string, error) {
	fixedLines := []string{}
	scanner := bufio.NewScanner(f)
	isHeader := true
	for scanner.Scan() {
		if isHeader {
			isHeader = false
			continue
		}
		line := scanner.Text()
		line = strings.ReplaceAll(line, "\\\"", "\"\"")
		fixedLines = append(fixedLines, line)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return nil, err
	}
	return fixedLines, nil
}
