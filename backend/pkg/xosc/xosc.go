package xosc

//go:generate stringer -type=OpenScenarioParameterType

type OpenScenarioParameterType byte

const (
	Integer OpenScenarioParameterType = iota
	Double
	String
	UnsignedInt
	UnsignedShort
	Boolean
	DateTime
)

type OpenScenarioParameter struct {
	Name  string
	Type  OpenScenarioParameterType
	Value string
}
type OpenScenario struct{}
