package user

import "github.com/google/uuid"


type user struct {
	id 			*uuid.UUID 
	username 	string 
	email 		string 
	password 	string 
}

func (u *user) ID() *uuid.UUID { 
	return u.id
}

func (u *user) Username() string { 
    return u.username
}

func (u *user) Email() string { 
    return u.email
}

func (u *user) Password() string { 
    return u.password
}

func (u *user) SetID(id *uuid.UUID) error { 
    u.id = id
    return nil
}

func (u *user) SetUsername(username string) error { 
    u.username = username
    return nil
}

func (u *user) SetEmail(email string) error { 
    u.email = email
    return nil
}

func (u *user) SetPassword(password string) error { 
    u.password = password
    return nil
}
