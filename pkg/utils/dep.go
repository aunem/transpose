package utils

import "os"

// CleanConstraints HACK: needs to be fixed in dep
func CleanConstraints() error {
	name := "Gopkg.toml"
	err := os.Remove(name)
	if err != nil {
		return err
	}
	_, err = os.OpenFile(name, os.O_RDONLY|os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	return nil
}
