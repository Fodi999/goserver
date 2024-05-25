package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

func main() {
	projectName := "goserver"

	// Создание файла go.mod
	cmd := exec.Command("go", "mod", "init", projectName)
	err := cmd.Run()
	if err != nil {
		log.Fatalf("Failed to initialize go module: %v", err)
	}

	// Создание базовой структуры каталогов
	dirs := []string{
		"cmd/goserver",
		"internal/handlers",
		"internal/router",
		"internal/database",
		"internal/customdata",
	}

	for _, dir := range dirs {
		err := os.MkdirAll(dir, 0755)
		if err != nil {
			log.Fatalf("Failed to create directory %s: %v", dir, err)
		}
	}

	// Создание примера основного файла
	mainFile := filepath.Join("cmd", "goserver", "main.go")
	mainContent := `package main

import (
	"log"
	"net/http"
	"goserver/internal/router"
	"goserver/internal/database"
)

func main() {
	database.InitDB("custom_data.db")
	r := router.NewRouter()
	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}
`
	err = os.WriteFile(mainFile, []byte(mainContent), 0644)
	if err != nil {
		log.Fatalf("Failed to write main.go: %v", err)
	}

	fmt.Println("Project initialized successfully.")
}
