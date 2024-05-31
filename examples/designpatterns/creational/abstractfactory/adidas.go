package main

type Adidas struct {
}

func (a *Adidas) makeShoes() IShoe {
	return &AdidasShoes{
		Shoe: Shoe{
			logo: "adidas",
			size: 13,
		},
	}
}

func (a *Adidas) makeShirt() IShirt {
	return &AdidasShirt{
		Shirt: Shirt{
			logo: "adidas",
			size: 23,
		},
	}
}
