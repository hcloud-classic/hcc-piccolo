package userTool

import "golang.org/x/crypto/bcrypt"

// GetBcryptPassword : Get bcrypt hashed password from provided password
func GetBcryptPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}