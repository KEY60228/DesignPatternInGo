package main

func main() {
	generator := NewRandomNumberGenerator(0)
	observer1 := NewDigitObserver()
	observer2 := NewGraphObserver()
	generator.AddObserver(observer1)
	generator.AddObserver(observer2)
	generator.Execute()
}
