package handler

import (
	"log"
	"net/http"
	"strconv"

	"../../../DB/Logica"
	"../Model"
	"github.com/gorilla/mux"
)

// GetFile...
func GetFiles(w http.ResponseWriter, r *http.Request) {
	files := []Model.File{}
	files = data.Read_File_DB()
	respondJSON(w, http.StatusOK, files)
}

// CreateFile...
func CreateFile(w http.ResponseWriter, r *http.Request) {
	file, handle, err := r.FormFile("UploadFile")
	if err != nil {
		log.Printf("Error loading the file %v", err)
		return
	}
	defer file.Close()
	data.Write_File_DB(file, handle.Filename, int(handle.Size))
	respondJSON(w, http.StatusCreated, "Upload successful")
}

// GetFile ...
func GetFile(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	for _, file := range data.Read_File_DB() {
		if file.FileID == id {
			respondJSON(w, http.StatusOK, file)
			return
		}
	}
	respondJSON(w, http.StatusNotFound, nil)
}

// DeleteFile...
func DeleteFile(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	for _, file := range data.Read_File_DB() {
		if file.FileID == id {
			data.DeleteFile(file.Name, id)
			respondJSON(w, http.StatusOK, "Delected successful")
			return
		}
	}
	respondJSON(w, http.StatusNotFound, nil)
}
