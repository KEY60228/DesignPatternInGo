package main

import "fmt"

type Singleton struct {
	singleton *Singleton
}

func NewSingleton() *Singleton {
	fmt.Println("インスタンスを生成")
	s := &Singleton{}
	s.singleton = s
	return s
}

func (s *Singleton) GetInstance() *Singleton {
	return s.singleton
}
