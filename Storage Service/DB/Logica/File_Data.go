package data

import (
	"bufio"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"

	"../../API/App/Model"
)

func Error(err error) {
	if err != nil {
		log.Println(err)
		panic(err)
	}
}

func Read_Dir_File() []os.FileInfo {
	f, err := os.Open("./Files")
	Error(err)

	files, err := f.Readdir(0)
	Error(err)

	return files
}

func Read_File_DB() []Model.File {
	Files := []Model.File{}

	Read, err := os.Open("./DB/Data/File_DB.txt")
	Error(err)
	defer Read.Close()

	scanner := bufio.NewScanner(Read)
	for scanner.Scan() {
		data := strings.Split(scanner.Text(), ",")
		fileID, _ := strconv.Atoi(data[0])
		size, _ := strconv.Atoi(data[2])
		Files = append(Files, Model.File{FileID: fileID, Name: data[1], Size: size})
	}
	return Files
}

func Write_File_DB(upload io.Reader, fileName string, fileSize int) {
	UploadFile(upload, fileName)
	Files := Read_File_DB()
	File, err := os.Create("./DB/Data/File_DB.txt")
	Error(err)
	defer File.Close()

	for _, row := range Files {
		fileID := strconv.Itoa(row.FileID)
		File.WriteString(fileID + "," + row.Name + "," + strconv.Itoa(row.Size) + "\n")
	}

	File.WriteString(strconv.Itoa(Generate_FileID()) + "," + fileName + "," + strconv.Itoa(fileSize) + "\n")
	File.Sync()
}

func Generate_FileID() int {
	Files := Read_File_DB()
	NewID := 0

	for _, row := range Files {
		if NewID < row.FileID {
			NewID = row.FileID
		}
	}
	return NewID + 1
}

func UploadFile(upload io.Reader, name string) {
	data, err := ioutil.ReadAll(upload)
	Error(err)

	err = ioutil.WriteFile("./Files/"+name, data, 0666)
	Error(err)
}

func DeleteFile(fileName string, id int) {
	os.Remove("./Files/" + fileName)

	Files := Read_File_DB()
	File, err := os.Create("./DB/Data/File_DB.txt")
	Error(err)
	defer File.Close()

	for _, row := range Files {
		fileID := strconv.Itoa(row.FileID)
		if row.FileID != id {
			File.WriteString(fileID + "," + row.Name + "," + strconv.Itoa(row.Size) + "\n")
		}
	}
	File.Sync()
}
