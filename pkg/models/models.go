package models

type Event interface {
	display()
}

// Base struct
type Log struct {
	ID     int
	Source string
	Body   string
}

type SystemLog struct {
	Log
	Severity string
}

func (s SystemLog) display() {
	println(s.ID, s.Source, s.Body)
}
