package abstractfactory

type Nike struct{}

type NikeShoe struct {
	Shoe
}

func (a *Nike) MakeShoe() IShoe {
	return &NikeShoe{
		Shoe: Shoe{
			logo: "nike",
			size: 40,
		},
	}
}

type NikeShirt struct {
	Shirt
}

func (a *Nike) MakeShirt() IShirt {
	return &NikeShirt{
		Shirt: Shirt{
			logo: "nike",
			size: 40,
		},
	}
}
