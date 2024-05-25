package handlers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"compress/gzip"

	"goserver/internal/customdata"
	"goserver/internal/database"
)

// HomeHandler - основной обработчик
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Received request: %s %s", r.Method, r.URL.Path)
	w.Write([]byte("Hello, World!"))
}

// FileHandler - обработчик файлов
func FileHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Received request: %s %s", r.Method, r.URL.Path)
	filePath := r.URL.Query().Get("file")
	if filePath == "" {
		WriteError(w, http.StatusBadRequest, "File parameter is missing")
		return
	}

	file, err := os.Open(filePath)
	if err != nil {
		WriteError(w, http.StatusNotFound, "File not found")
		return
	}
	defer file.Close()

	w.Header().Set("Content-Encoding", "gzip")
	gzipWriter := gzip.NewWriter(w)
	defer gzipWriter.Close()

	_, err = io.Copy(gzipWriter, file)
	if err != nil {
		WriteError(w, http.StatusInternalServerError, "Failed to send file")
	}
}

// CustomDataHandler - обработчик пользовательских данных, отправляет данные в формате JSON
func CustomDataHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Received request: %s %s", r.Method, r.URL.Path)
	key := r.URL.Query().Get("key")
	if key == "" {
		WriteError(w, http.StatusBadRequest, "Key parameter is missing")
		return
	}

	_, value, color, err := database.GetCustomData(key)
	if err != nil {
		WriteError(w, http.StatusNotFound, "Data not found")
		return
	}

	data := customdata.CustomData{
		Key:   key,
		Value: value,
		Color: color,
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(data); err != nil {
		WriteError(w, http.StatusInternalServerError, "Failed to encode data")
	}
}

// CustomDataPostHandler - обработчик для получения данных в собственном формате
func CustomDataPostHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Received request: %s %s", r.Method, r.URL.Path)
	if r.Method != http.MethodPost {
		WriteError(w, http.StatusMethodNotAllowed, "Invalid request method")
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		WriteError(w, http.StatusInternalServerError, "Failed to read request body")
		return
	}

	data, err := customdata.FromBinary(body)
	if err != nil {
		WriteError(w, http.StatusBadRequest, "Invalid data format")
		return
	}

	err = database.SaveCustomData(data.Key, data.Value, data.Color)
	if err != nil {
		WriteError(w, http.StatusInternalServerError, "Failed to save data")
		return
	}

	response, err := json.Marshal(data)
	if err != nil {
		WriteError(w, http.StatusInternalServerError, "Failed to encode response")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

// UpdateCustomDataHandler - обработчик для обновления данных
func UpdateCustomDataHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Received request: %s %s", r.Method, r.URL.Path)
	if r.Method != http.MethodPut {
		WriteError(w, http.StatusMethodNotAllowed, "Invalid request method")
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		WriteError(w, http.StatusInternalServerError, "Failed to read request body")
		return
	}

	data, err := customdata.FromBinary(body)
	if err != nil {
		WriteError(w, http.StatusBadRequest, "Invalid data format")
		return
	}

	err = database.UpdateCustomData(data.Key, data.Value, data.Color)
	if err != nil {
		WriteError(w, http.StatusInternalServerError, "Failed to update data")
		return
	}

	response, err := json.Marshal(data)
	if err != nil {
		WriteError(w, http.StatusInternalServerError, "Failed to encode response")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

// DeleteCustomDataHandler - обработчик для удаления данных
func DeleteCustomDataHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Received request: %s %s", r.Method, r.URL.Path)
	if r.Method != http.MethodDelete {
		WriteError(w, http.StatusMethodNotAllowed, "Invalid request method")
		return
	}

	key := r.URL.Query().Get("key")
	if key == "" {
		WriteError(w, http.StatusBadRequest, "Key parameter is missing")
		return
	}

	err := database.DeleteCustomData(key)
	if err != nil {
		WriteError(w, http.StatusInternalServerError, "Failed to delete data")
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// GetAllCustomDataHandler - обработчик для получения всех данных
func GetAllCustomDataHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Received request: %s %s", r.Method, r.URL.Path)
	data, err := database.GetAllCustomData()
	if err != nil {
		WriteError(w, http.StatusInternalServerError, "Failed to get data")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(data); err != nil {
		WriteError(w, http.StatusInternalServerError, "Failed to encode data")
	}
}

// CustomFileHandler - обработчик для чтения и записи файлов с расширением .fd
func CustomFileHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Received request: %s %s", r.Method, r.URL.Path)
	switch r.Method {
	case http.MethodGet:
		filePath := r.URL.Query().Get("file")
		log.Printf("Received GET request for file: %s", filePath)
		if filePath == "" {
			WriteError(w, http.StatusBadRequest, "File parameter is missing")
			return
		}

		file, err := os.Open(filePath)
		if err != nil {
			log.Printf("File not found: %s", filePath)
			WriteError(w, http.StatusNotFound, "File not found")
			return
		}
		defer file.Close()

		content, err := io.ReadAll(file)
		if err != nil {
			WriteError(w, http.StatusInternalServerError, "Failed to read file")
			return
		}

		data, err := customdata.FromBinary(content)
		if err != nil {
			WriteError(w, http.StatusBadRequest, "Invalid file format")
			return
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(data); err != nil {
			WriteError(w, http.StatusInternalServerError, "Failed to encode data")
		}

	case http.MethodPost:
		filePath := r.URL.Query().Get("file")
		log.Printf("Received POST request for file: %s", filePath)
		if filePath == "" {
			WriteError(w, http.StatusBadRequest, "File parameter is missing")
			return
		}

		body, err := io.ReadAll(r.Body)
		if err != nil {
			WriteError(w, http.StatusInternalServerError, "Failed to read request body")
			return
		}

		data, err := customdata.FromBinary(body)
		if err != nil {
			WriteError(w, http.StatusBadRequest, "Invalid data format")
			return
		}

		file, err := os.Create(filePath)
		if err != nil {
			WriteError(w, http.StatusInternalServerError, "Failed to create file")
			return
		}
		defer file.Close()

		binData, err := data.ToBinary()
		if err != nil {
			WriteError(w, http.StatusInternalServerError, "Failed to encode data")
			return
		}

		if _, err := file.Write(binData); err != nil {
			WriteError(w, http.StatusInternalServerError, "Failed to write to file")
		}

	default:
		WriteError(w, http.StatusMethodNotAllowed, "Unsupported request method")
	}
}







