package main

type EventName string

const (
	EventDataUpdate EventName = "data:update"
)

var Events = []struct {
	Value  EventName
	TSName string
}{
	{EventDataUpdate, "DATA_UPDATE"},
}
