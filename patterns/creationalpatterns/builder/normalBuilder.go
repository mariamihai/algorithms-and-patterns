package builder

type NormalBuilder struct {
	windowType string
	doorType   string
	floor      int
}

func (n *NormalBuilder) resetBuilder() {
	*n = NormalBuilder{}
}

func NewNormalBuilder() *NormalBuilder {
	return &NormalBuilder{}
}

func (n *NormalBuilder) setWindowType() {
	n.windowType = "wooden window"
}

func (n *NormalBuilder) setDoorType() {
	n.doorType = "wooden door"
}

func (n *NormalBuilder) setNumFloor(floor int) {
	n.floor = floor
}

func (n *NormalBuilder) getHouse() House {
	return House{
		WindowType: n.windowType,
		DoorType:   n.doorType,
		Floor:      n.floor,
	}
}
