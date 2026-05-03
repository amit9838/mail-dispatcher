package internal

import (
	"encoding/csv"
	"os"
)

type Recipient struct {
	Name  string
	Email string
}

func LoadRecipient(filePath string, ch chan Recipient) error {
	defer close(ch)

	f, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer f.Close()

	r := csv.NewReader(f)
	records, err := r.ReadAll()

	if err != nil {
		return err
	}

	for _, record := range records[1:] {
		ch <- Recipient{
			Name:  record[0],
			Email: record[1],
		}
	}
	return nil
}
