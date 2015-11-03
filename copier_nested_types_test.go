package copier

import (
	"testing"
)

type SourceNested struct {
	Age    int
	Secret int
}

type SourceType struct {
	Name   string
	Nested SourceNested
}

type DestinationNested struct {
	Age int
}

type DestinationType struct {
	Name   string
	Nested DestinationNested
}

func TestSkipFieldOfDifferentType(t *testing.T) {
	source := SourceType{
		Name: "Test",
		Nested: SourceNested{
			Age:    23,
			Secret: 2,
		},
	}

	destination := DestinationType{Nested: DestinationNested{}}

	Copy(&destination, &source)

}
