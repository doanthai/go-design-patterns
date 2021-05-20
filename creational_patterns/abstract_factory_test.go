package creational_patterns

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAbstractFactory(t *testing.T) {
	adidasFactory, _ := GetSportFactoryMaker("adidas")
	nikeFactory, _ := GetSportFactoryMaker("nike")

	adidasShoe := adidasFactory.makeShoe(12)
	adidasShirt := adidasFactory.makeShirt(39)

	nikeShoe := nikeFactory.makeShoe(12)
	nikeShirt := nikeFactory.makeShirt(39)

	expectAdidasShoe := &Shoe{
		logo: "Adidas",
	}
	expectAdidasShoe.setSize(12)
	expectAdidasShirt := &Shirt{
		logo: "Adidas",
	}
	expectAdidasShirt.setSize(12)
	expectNikeShoe := &Shoe{
		logo: "Nike",
	}
	expectNikeShoe.setSize(12)
	expectNikeShirt := &Shirt{
		logo: "Nike",
	}
	expectNikeShirt.setSize(39)

	assert.Equal(t, adidasShoe.getLogo(), expectAdidasShoe.getLogo())
	assert.Equal(t, adidasShirt.getLogo(), expectAdidasShirt.getLogo())
	assert.Equal(t, nikeShoe.getLogo(), expectNikeShoe.getLogo())
	assert.Equal(t, nikeShirt.getLogo(), expectNikeShirt.getLogo())
}

func TestInCorrectBrandName(t *testing.T) {
	inCorrectFactory, error := GetSportFactoryMaker("InCorrect")

	assert.Nil(t, inCorrectFactory)
	assert.Errorf(t, error, "Wrong brand type passed")
}