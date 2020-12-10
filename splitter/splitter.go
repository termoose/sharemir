package splitter

import (
	"fmt"
	"github.com/corvus-ch/shamir"
	"io/ioutil"
)

type Part []byte

func Split(file string, parts, threshold int) ([]Part, error) {
	data, err := ioutil.ReadFile(file)

	if err != nil {
		return nil, fmt.Errorf("error reading file: %w", err)
	}

	fileParts, err := shamir.Split(data, parts, threshold)

	if err != nil {
		return nil, fmt.Errorf("error splitting: %w", err)
	}

	return partsToBytes(fileParts), nil
}

func PartsToFiles(parts []Part, nameTemplate string) error {
	for i, part := range parts {
		filename := fmt.Sprintf(nameTemplate, i)
		err := ioutil.WriteFile(filename, part, 0644)

		if err != nil {
			return fmt.Errorf("error writing to file: %w", err)
		}
	}

	return nil
}

func Combine(byteParts []Part) ([]byte, error) {
	parts := bytesToParts(byteParts)
	secret, err := shamir.Combine(parts)

	if err != nil {
		return nil, fmt.Errorf("error combining: %w", err)
	}

	return secret, nil
}

func bytesToParts(bytes []Part) map[byte][]byte {
	result := make(map[byte][]byte)

	for _, part := range bytes {
		result[part[0]] = part[1:]
	}

	return result
}

func partsToBytes(parts map[byte][]byte) []Part {
	var result []Part

	for key, val := range parts {
		p := append([]byte{key}, val...)
		result = append(result, p)
	}

	return result
}