package customdata

import (
	"log"
	"os"
)

// GenerateBinaryFile создает бинарный файл с данными CustomData
func GenerateBinaryFile(filename string, data CustomData) {
	binData, err := data.ToBinary()
	if err != nil {
		log.Fatalf("Failed to convert to binary: %v", err)
	}

	err = os.WriteFile(filename, binData, 0644)
	if err != nil {
		log.Fatalf("Failed to write file: %v", err)
	}

	log.Printf("Binary file '%s' created successfully.", filename)
}
