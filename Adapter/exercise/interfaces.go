package main

type FileIO interface {
	ReadFromFile(string) error
	WriteToFile(string) error
	SetValue(string, string)
	GetValue(string) string
}
