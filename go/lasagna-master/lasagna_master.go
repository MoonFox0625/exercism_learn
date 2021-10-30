package lasagna

func PreparationTime(layers []string, average int) int {
	if average <= 0 {
		average = 2
	}
	return average * len(layers)
}

func Quantities(layers []string) (noodles int, sauce float64) {

	for _, layer := range layers {
		if layer == "noodles" {
			noodles += 50
		} else if layer == "sauce" {
			sauce += 0.2
		}
	}
	return
}

func AddSecretIngredient(friendsList []string, myList []string) []string {
	return append(myList, friendsList[len(friendsList)-1])
}

func ScaleRecipe(quantities []float64, portions int) (scaledQuantities []float64) {
	for _, quantity := range quantities {
		scaledQuantities = append(scaledQuantities, quantity/2*float64(portions))
	}
	return
}
