package main

func main() {
	alice := NewNoSupport("Alice")
	bob := NewLimitSupport("Bob", 100)
	charlie := NewSpecialSupport("Charlie", 429)
	diana := NewLimitSupport("Diana", 200)
	elmo := NewOddSupport("Elmo")
	fred := NewLimitSupport("Fred", 300)

	alice.SetNext(bob).SetNext(charlie).SetNext(diana).SetNext(elmo).SetNext(fred)

	for i := 0; i < 500; i += 33 {
		alice.Do(alice, NewTrouble(i))
	}
}
