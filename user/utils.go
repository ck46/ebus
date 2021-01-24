package user

import "fmt"

const (
	UniqueConstraintEmail = "users_email_key"
)

type UserNotExistsError struct{}
type EmailNotExistsError struct{}

type EmailDuplicateError struct {
	Email string
}

func (*UserNotExistsError) Error() string {
	return "User does not exist!"
}

func (*EmailNotExistsError) Error() string {
	return "Email does not exist!"
}

func (e *EmailDuplicateError) Error() string {
	return fmt.Sprintf("Email '%s' already exists", e.Email)
}

type PasswordMismatchError struct{}

func (e *PasswordMismatchError) Error() string {
	return "password didn't match"
}
