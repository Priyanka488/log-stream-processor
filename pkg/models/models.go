package models

import "time"

type Event interface {
	display()
}

// Base struct
type Log struct {
	ID          int
	Source      string
	Body        string
	TimeCreated time.Time
}

type SystemLog struct {
	Log
	Severity string
}

func (s SystemLog) display() {
	println(s.ID, s.Source, s.Body)
}

func (l *Log) SetTimeCreated() {
	l.TimeCreated = time.Now()
}
