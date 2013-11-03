package models

import (
    "errors"
    "os"
    "crypto/sha512"
    "io"
)

type User struct {
    Name string
    Password string
}

func (u *User) Verify() error {
    if !u.CorrectPassword() {
        return errors.New("Bad password")
    }
    return nil
}

func (u *User) CorrectPassword() bool {
    hash := sha512.New()
    io.WriteString(hash, u.Password)
    return string(hash.Sum(nil)) ==
        os.Getenv("HASHED_PASSWORD")
}
