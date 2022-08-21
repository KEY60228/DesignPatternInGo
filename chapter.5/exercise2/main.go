package main

import "fmt"

func main() {
	triple := NewTriple()
	alpha1 := triple.GetInstance("alpha")
	alpha2 := triple.GetInstance("alpha")
	beta := triple.GetInstance("beta")
	gamma := triple.GetInstance("gamma")

	fmt.Println(alpha1 == alpha2)
	fmt.Println(alpha1 == beta)
	fmt.Println(alpha1 == gamma)
	fmt.Println(alpha2 == beta)
	fmt.Println(alpha2 == gamma)
	fmt.Println(beta == gamma)
}
