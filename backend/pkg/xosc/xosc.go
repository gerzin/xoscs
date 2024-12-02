package xosc

//go:generate stringer -type=OpenScenarioDataType

type OpenScenarioDataType byte

const (
	Integer OpenScenarioDataType = iota
	Double
	String
	UnsignedInt
	UnsignedShort
	Boolean
	DateTime
)

//go:generate stringer -type=OpenScenarioConstraintRule

type OpenScenarioConstraintRule byte

const (
	EqualTo OpenScenarioConstraintRule = iota
	GreaterThan
	LessThan
	GreaterOrEqual
	LessOrEqual
	NotEqualTo
)

type OpenScenarioValueConstraint struct {
	Rule  OpenScenarioConstraintRule
	Value string
}

type OpenScenarioParameter struct {
	Name            string
	Type            OpenScenarioDataType
	Value           string
	ConstraintGroup []OpenScenarioValueConstraint
}
type OpenScenario struct {
	Parameters []OpenScenarioParameter
}
