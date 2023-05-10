package piscine

type food struct {
	preptime int
}

func FoodDeliveryTime(order string) int {
	var eda food
	if order == "burger" {
		eda.preptime = 15
		return eda.preptime
	} else if order == "chips" {
		eda.preptime = 10
		return eda.preptime
	} else if order == "nuggets" {
		eda.preptime = 12
		return eda.preptime
	} else {
		return 404
	}
}
