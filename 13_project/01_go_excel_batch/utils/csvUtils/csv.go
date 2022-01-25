package csvUtils

import (
	"encoding/csv"
	"io"
	"os"
)

func CsvReadAll(csvFileName string) ([]string, [][]string, error) {
	csvFile, err := os.Open(csvFileName)
	if err != nil {
		return nil, nil, err
	}
	defer csvFile.Close()

	hdr := make([]string, 0)
	lines := make([][]string, 0)

	reader := csv.NewReader(csvFile)
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, nil, err
		}

		if len(hdr) == 0 {
			hdr = record
		} else {
			lines = append(lines, record)
		}
	}

	return hdr, lines, nil
}
