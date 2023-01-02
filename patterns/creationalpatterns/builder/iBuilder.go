package builder

import "fmt"

type IBuilder interface {
	resetBuilder()
	setWindowType()
	setDoorType()
	setNumFloor(floor int)
	getHouse() House
}

func GetBuilder(builderType string) (IBuilder, error) {
	if builderType == "normal" {
		return NewNormalBuilder(), nil
	}

	if builderType == "igloo" {
		return NewIglooBuilder(), nil
	}

	return nil, fmt.Errorf("invalid builder type")
}
