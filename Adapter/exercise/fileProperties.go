package main

import "os"

type FileProperties struct {
	Properties
}

func NewFileProperties() *FileProperties {
	return &FileProperties{NewProperties()}
}

func (fp *FileProperties) ReadFromFile(filename string) error {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	if err := fp.Load(f); err != nil {
		return err
	}
	return nil
}

func (fp *FileProperties) WriteToFile(filename string) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	if err := fp.Store(f, "#written by FileProperties"); err != nil {
		return err
	}
	return nil
}

func (fp *FileProperties) SetValue(key string, value string) {
	fp.Properties[key] = value
}

func (fp *FileProperties) GetValue(key string) string {
	return fp.Properties[key]
}
