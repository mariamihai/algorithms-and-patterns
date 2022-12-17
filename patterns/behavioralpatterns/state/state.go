package state

type State interface {
	AddItem(int) error
	RequestItem() error
	InsertMoney(int) error
	DisperseItem() error
}
