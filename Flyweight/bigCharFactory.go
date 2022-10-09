package main

type BigCharFactory struct {
	factory *BigCharFactory
	pool    map[string]*BigChar
}

func NewBigCharFactory() *BigCharFactory {
	f := &BigCharFactory{}
	f.pool = make(map[string]*BigChar)
	f.factory = f
	return f
}

func (f *BigCharFactory) GetInstance() *BigCharFactory {
	return f.factory
}

func (f *BigCharFactory) GetBigChar(charname string) *BigChar {
	bc, ok := f.pool[charname]
	if !ok {
		bc = NewBigChar(charname)
		f.pool[charname] = bc
	}
	return bc
}
