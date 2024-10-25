package user

import (
	"golphin/src/utils/validator"
	"github.com/google/uuid"
)

type builder struct {
	fields 			[]string
	errorMessages 	[]string
	user 			*user
}

func NewBuilder() *builder {
	return &builder{
        fields: []string{},
        errorMessages: []string{},
        user: &user{},
    }
}

func (b *builder) SetID(id uuid.UUID) *builder {
	if !validator.IsUUIDValid(id) {
		b.fields = append(b.fields, "User ID")
		b.errorMessages = append(b.errorMessages, "Invalid User ID")
		return b
	}
	b.user.id = &id
	return b
}

func (b *builder) SetUsername(username string) *builder {
	if len(username) < 5 || len(username) > 16 {
        b.fields = append(b.fields, "Username")
        b.errorMessages = append(b.errorMessages, "Username must be between 5 and 20 characters")
        return b
    }
	if validator.IsTextBlank(username) {
        b.fields = append(b.fields, "Username")
        b.errorMessages = append(b.errorMessages, "Username cannot be blank")
        return b
    }
    b.user.username = username
    return b
}

func (b *builder) SetEmail(email string) *builder {
	if !validator.IsEmailValid(email) {
        b.fields = append(b.fields, "Email")
        b.errorMessages = append(b.errorMessages, "Invalid Email")
        return b
    }
    b.user.email = email
    return b
}

func (b *builder) SetPassword(password string) *builder {
	if len(password) < 8 || len(password) > 32 {
        b.fields = append(b.fields, "Password")
        b.errorMessages = append(b.errorMessages, "Password must be between 8 and 32 characters")
        return b
    }
    if validator.IsTextBlank(password) {
        b.fields = append(b.fields, "Password")
        b.errorMessages = append(b.errorMessages, "Password cannot be blank")
        return b
    }
    b.user.password = password
    return b
}