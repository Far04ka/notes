package constants

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type UsrInf struct {
	UserName  string
	Id        int
	Password  string
	ListNotes *[]Note
}

type Note struct {
	Id     int
	Text   string
	Date   string
	Header string
	UserId int
}

func (u UsrInf) RegValidate() error {
	return validation.ValidateStruct(&u,
		// Street cannot be empty, and the length must between 5 and 50
		validation.Field(&u.UserName, validation.Required, validation.Length(5, 50)),
		// City cannot be empty, and the length must between 5 and 50
		validation.Field(&u.Password, validation.Required, validation.Length(8, 50)),
	)
}
