package main

import (
	"fmt"
	//"go_design_pattern/decorator/internal/beverage"

	"github.com/rstliz/beverage"
)

func main() {
	espresso := new(beverage.Espresso)
	fmt.Printf("%s $%.2f\n", espresso.Description(), espresso.Cost())

	var darkRoast beverage.Beverage = &beverage.DarkRoast{}
	darkRoast = beverage.NewMocha(darkRoast)
	darkRoast = beverage.NewMocha(darkRoast)
	darkRoast = beverage.NewWhip(darkRoast)
	fmt.Printf("%s $%.2f\n", darkRoast.Description(), darkRoast.Cost())

	var houseBlend beverage.Beverage = &beverage.HouseBlend{}
	houseBlend = beverage.NewMilk(houseBlend)
	houseBlend = beverage.NewMocha(houseBlend)
	houseBlend = beverage.NewWhip(houseBlend)
	fmt.Printf("%s $%.2f\n", houseBlend.Description(), houseBlend.Cost())
}
