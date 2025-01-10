package api

import (

	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/database"
)


type User struct{
	UserId 		uint64 	`json:"userID"`
	UserName 	string	`json:"name"`
	UserPhoto 	string 	`json:"userPhoto"`
}


func (u *User) FromDatabase(user database.User){
	u.UserId = user.UserId
	u.UserName = user.UserName
	u.UserPhoto = user.UserPhoto
}

func (u *User) ToDatabase() database.User{
	return database.User{
		UserId: u.UserId,
		UserPhoto: u.UserPhoto,
		UserName: u.UserName,
	}
}