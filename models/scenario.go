package models

type Scenario int

const (
	Local Scenario = iota
	Cloud
	MEC
)
