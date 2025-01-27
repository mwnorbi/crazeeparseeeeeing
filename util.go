package main

import "bytes"

func getDelimiterIndex(input []byte, delimiter string) *int {
	deli := []byte(delimiter)
	for i := range input {
		if input[i] == deli[0] {
			if bytes.Equal(input[i:i+len(deli)], deli) {
				return &i
			}
		}
	}
	return nil
}
