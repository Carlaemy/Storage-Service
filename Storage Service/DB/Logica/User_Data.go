package data

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"

	"../../API/App/Model"
)

func Read_User_DB() []Model.User {
	user := []Model.User{}
	file, err := os.Open("./DB/Data/User_DB.txt")
	Error(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		data := strings.Split(scanner.Text(), ",")
		userID, _ := strconv.Atoi(data[0])
		user = append(user, Model.User{UserID: userID, Name: data[1], LastName: data[2], Email: data[3], Password: data[4]})
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return user
}

func Write_User_DB(user Model.User, isDelected bool) {
	users := Read_User_DB()

	file, err := os.Create("./DB/Data/User_DB.txt")
	Error(err)
	defer file.Close()

	for _, row := range users {
		userID := strconv.Itoa(row.UserID)
		if isDelected {
			if row.UserID != user.UserID {
				file.WriteString(userID + "," + row.Name + "," + row.LastName + "," + row.Email + "," + row.Password + "\n")
			}
		} else {
			file.WriteString(userID + "," + row.Name + "," + row.LastName + "," + row.Email + "," + row.Password + "\n")
		}
	}

	if !isDelected {
		file.WriteString(strconv.Itoa(Generate_UserID()) + "," + user.Name + "," + user.LastName + "," + user.Email + "," + user.Password + "\n")
	}
	file.Sync()
}

func Generate_UserID() int {
	NewID := 0
	users := Read_User_DB()

	for _, row := range users {
		if NewID < row.UserID {
			NewID = row.UserID
		}
	}
	return NewID + 1
}

func Login(email string, password string) string {
	for _, user := range Read_User_DB() {
		if user.Email == email && user.Password == password {
			return "Success"
		}
	}
	return "Invalid"
}
