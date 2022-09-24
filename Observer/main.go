package main

func main() {
	generator := NewIncrementalNumberGenerator(10, 50, 5)
	observer1 := NewDigitObserver()
	observer2 := NewGraphObserver()
	generator.AddObserver(observer1)
	generator.AddObserver(observer2)
	generator.Execute()
}
