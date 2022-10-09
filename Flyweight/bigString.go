package main

type BigString struct {
	bigChars []*BigChar
}

func NewBigString(factory *BigCharFactory, str string) *BigString {
	f := factory.GetInstance()
	bigChars := make([]*BigChar, len(str))
	for i := 0; i < len(bigChars); i++ {
		bigChars[i] = f.GetBigChar(str[i : i+1])
	}
	return &BigString{bigChars: bigChars}
}

func (bs *BigString) Print() {
	for _, bc := range bs.bigChars {
		bc.Print()
	}
}
