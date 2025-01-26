package filestore

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"task-tracker-cli/constant"
	"task-tracker-cli/entity"
)


type FileStore struct {
	filePath string
	serialiazatinMode string
}

func New(path, serialiazatinMode string)FileStore {
	return FileStore{
		filePath: path,serialiazatinMode: serialiazatinMode}
}

func (f FileStore)Save(u entity.User){
	f.writeUserToFile(u)
}

func (f FileStore)Load() []entity.User {
	var uStorage []entity.User
	file, err := os.Open(f.filePath)
	
	if err != nil {
		fmt.Println("there is no file", err)
	}

	var data = make([]byte, 1024)
	_, oErr := file.Read(data)

	if oErr != nil {
		fmt.Println("can't read from ", oErr)
	}

	var dataString = string(data)
	dataString = strings.Trim(dataString, "\n")
	userSlice := strings.Split(dataString, "\n")

	for _, u := range userSlice {
		
		var userStruct = entity.User{}

		switch f.serialiazatinMode {
		case constant.OldOneSerilizationMode:
			var dErr error
			userStruct, dErr = deSerilizedOldOne(u)

			if dErr != nil {
				fmt.Println("cant desrilized user record to user struct", dErr)
				return nil
			}

		case constant.JsonSerializationMode:
			if u[0] != '{' && u[len(u)-1] != '}' {
				continue
			}

			uErr := json.Unmarshal([]byte(u), &userStruct)
			if uErr != nil {
				fmt.Println("cant desrilized user record to user struct from json mode", uErr)
				return nil
			}
		}
		uStorage = append(uStorage, userStruct)

	}

	return uStorage

}

func (f FileStore)writeUserToFile(user entity.User) {
	var file *os.File

	file, err := os.OpenFile(f.filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		fmt.Println("can't write file", err)
	}
	defer file.Close()
	// Serialized the user struct
	var data []byte
	if f.serialiazatinMode == constant.OldOneSerilizationMode {
		data = []byte(fmt.Sprintf("id: %d, name: %s, email: %s, password: %s\n",
			user.ID, user.Name, user.Email, user.Password))

	} else if f.serialiazatinMode == constant.JsonSerializationMode {
		var jErr error
		data, jErr = json.Marshal(user)
		data = append(data, []byte("\n")...)
		if err != nil {
			fmt.Println("can't mashal user struct to json", jErr)
			return
		}
	} else {
		fmt.Println("invalid serilization mode")

		return
	}

	_, wErr := file.Write([]byte(data))
	if wErr != nil {
		fmt.Println("can't write file", wErr)
	}

}

func deSerilizedOldOne(userStr string) (entity.User, error) {
	if userStr == "" {
		return entity.User{}, errors.New("use string is empty ")
	}
	userFields := strings.Split(userStr, ",")
	var user = entity.User{}
	for _, field := range userFields {
		values := strings.Split(field, ": ")
		if len(values) != 2 {
			fmt.Printf("invalid field format: %v\n", field)
			continue
		}
		fieldName := strings.ReplaceAll(values[0], " ", "")
		fieldValue := values[1]

		switch fieldName {
		case "id":
			id, err := strconv.Atoi(fieldValue)
			if err != nil {
				fmt.Println("error in ")

				return entity.User{}, errors.New("strconv error")
			}
			user.ID = id

		case "name":
			user.Name = fieldValue
		case "email":
			user.Email = fieldValue
		case "password":
			user.Password = fieldValue
		}
	}
	return user, nil

}

func (f FileStore)loadUserFromStorage()[]entity.User{
	return f.Load()

}