package main

import (
	"os"
	"errors"
	"fmt"
)

// create a new error
func NewError(msg string, a ...interface{}) error {
	return errors.New(fmt.Sprintf(msg, a))
}

// check if a dir is existed, if not, create it
func DirExistedOrCreate(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.Mkdir(path, 0644)
		return nil
	} else {
		return err
	}
}


