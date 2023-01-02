package builder

type IglooBuilder struct {
	windowType string
	doorType   string
	floor      int
}

func NewIglooBuilder() *IglooBuilder {
	return &IglooBuilder{}
}

func (i *IglooBuilder) resetBuilder() {
	*i = IglooBuilder{}
}

func (i *IglooBuilder) setWindowType() {
	i.windowType = "snow window"
}

func (i *IglooBuilder) setDoorType() {
	i.doorType = "snow door"
}

func (i *IglooBuilder) setNumFloor(floor int) {
	i.floor = floor
}

func (i *IglooBuilder) getHouse() House {
	return House{
		WindowType: i.windowType,
		DoorType:   i.doorType,
		Floor:      i.floor,
	}
}
