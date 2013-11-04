package models

import (
    "errors"
    "os"
    "crypto/sha512"
    "io"
    "encoding/hex"
)

type User struct {
    Name string
    Password string
}

func (u *User) Verify() error {
    if !u.CorrectPassword() {
        return errors.New("Bad password")
    }
    if !u.CorrectUsername() {
        return errors.New("Bad username")
    }
    return nil
}

func (u *User) CorrectPassword() bool {
    hash := sha512.New()
    io.WriteString(hash, u.Password)
    return hex.EncodeToString(hash.Sum(nil)) ==
        os.Getenv("HASHED_PASSWORD")
}

func (u *User) CorrectUsername() bool {
    hash := sha512.New()
    io.WriteString(hash, u.Name)
    return hex.EncodeToString(hash.Sum(nil)) ==
        os.Getenv("HASHED_USERNAME")
}
