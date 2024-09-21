package piscine

type food struct {
	preptime int
}

func FoodDeliveryTime(order string) int {
	var burger food
	burger.preptime = 15
	var chips food
	chips.preptime = 10
	var nuggets food
	nuggets.preptime = 12
	if order == "burger" {
		return burger.preptime
	} else if order == "chips" {
		return chips.preptime
	} else if order == "nuggets" {
		return nuggets.preptime
	} else {
		return 404
	}
}
