package abstractfactory

type Adidas struct{}

type AdidasShoe struct {
	Shoe
}

func (a *Adidas) MakeShoe() IShoe {
	return &AdidasShoe{
		Shoe: Shoe{
			logo: "adidas",
			size: 40,
		},
	}
}

type AdidasShirt struct {
	Shirt
}

func (a *Adidas) MakeShirt() IShirt {
	return &AdidasShirt{
		Shirt: Shirt{
			logo: "adidas",
			size: 40,
		},
	}
}
