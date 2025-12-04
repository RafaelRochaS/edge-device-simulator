package models

type Scenario int

const (
	Local Scenario = iota
	Cloud
	MEC
)

func (s Scenario) String() string {
	return [...]string{"Local", "Cloud", "MEC"}[s]
}
