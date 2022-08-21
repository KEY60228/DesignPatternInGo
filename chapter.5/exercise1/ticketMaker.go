package main

type TicketMaker struct {
	ticketMaker *TicketMaker
	ticket      int
}

func NewTicketMaker() *TicketMaker {
	tm := &TicketMaker{ticket: 1000}
	tm.ticketMaker = tm
	return tm
}

func (tm *TicketMaker) GetInstance() *TicketMaker {
	return tm.ticketMaker
}

func (tm *TicketMaker) GetNextTicketNumber() int {
	tm.ticket++
	return tm.ticket
}
