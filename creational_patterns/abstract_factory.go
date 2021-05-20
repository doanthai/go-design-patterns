package creational_patterns

import "fmt"

type SportFactoryMaker interface {
	makeShoe(size int) ShoeMaker
	makeShirt(size int) ShirtMaker
}

func GetSportFactoryMaker(brand string) (SportFactoryMaker, error)  {
	if brand == "adidas" {
		return &adidas{}, nil
	}
	if brand == "nike" {
		return &nike{}, nil
	}

	return nil, fmt.Errorf("Wrong brand type passed")
}

type adidas struct {
}

func (adidas *adidas) makeShoe(size int) ShoeMaker {
	return &adidasShoe{
		Shoe{
			logo: "Adidas",
			size: size,
		},
	}
}

func (adidas *adidas) makeShirt(size int) ShirtMaker {
	return &adidasShirt{
		Shirt{
			logo: "Adidas",
			size: size,
		},
	}
}

type nike struct {
}

func (nike *nike) makeShoe(size int) ShoeMaker {
	return &nikeShoe{
		Shoe{
			logo: "Nike",
			size: size,
	}}
}

func (nike *nike) makeShirt(size int) ShirtMaker {
	return &nikeShirt{
		Shirt{
			logo: "Nike",
			size: size,
		},
	}
}

type Shoe struct {
	logo string
	size int
}

type ShoeMaker interface {
	setLogo(logo string)
	setSize(size int)
	getLogo() string
	getSize() int
}

func (shoe *Shoe) setLogo(logo string) {
	shoe.logo = logo
}

func (shoe *Shoe) setSize(size int)  {
	shoe.size = size
}

func (shoe *Shoe) getLogo() string {
	return shoe.logo
}

func (shoe *Shoe) getSize() int {
	return shoe.size
}

type adidasShoe struct {
	Shoe
}

type nikeShoe struct {
	Shoe
}

type Shirt struct {
	logo string
	size int
}

type ShirtMaker interface {
	setLogo(logo string)
	setSize(size int)
	getLogo() string
	getSize() int
}

func (shirt *Shirt) setLogo(logo string) {
	shirt.logo = logo
}

func (shirt *Shirt) setSize(size int) {
	shirt.size = size
}

func (shirt *Shirt) getLogo() string {
	return shirt.logo
}

func (shirt *Shirt) getSize() int {
	return shirt.size
}

type adidasShirt struct {
	Shirt
}

type nikeShirt struct {
	Shirt
}