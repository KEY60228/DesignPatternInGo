package main

import "log"

type Triple struct {
	alpha *Triple
	beta  *Triple
	gamma *Triple
}

func NewTriple() *Triple {
	alpha := &Triple{}
	beta := &Triple{}
	gamma := &Triple{}
	alpha.alpha = alpha
	alpha.beta = beta
	alpha.gamma = gamma
	return alpha
}

func (t *Triple) GetInstance(name string) *Triple {
	switch name {
	case "alpha":
		return t.alpha
	case "beta":
		return t.beta
	case "gamma":
		return t.gamma
	default:
		log.Fatal("not match")
	}
	return nil
}
