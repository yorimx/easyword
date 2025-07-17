package users

import "github.com/gogf/gf/crypto/gmd5"

type Users struct {
}

func New() *Users {
	return &Users{}
}

func (u *Users) encryptPassword(password string) string {
	// Implement password encryption logic here
	return gmd5.MustEncryptString(password) // Placeholder, replace with actual encryption

}
