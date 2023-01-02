package factorymethod

// Concrete product

type AK47 struct {
	Gun
}

func NewAK47() IGun {
	return &AK47{
		Gun: Gun{
			name:  "Ak47",
			power: 5,
		},
	}
}
