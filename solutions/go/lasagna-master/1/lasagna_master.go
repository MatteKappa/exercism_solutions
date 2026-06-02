package lasagnamaster

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, avgMins int) int {
	if avgMins == 0 {
		return len(layers) * 2
	}
	return len(layers) * avgMins
}
// TODO: define the 'Quantities()' function
func Quantities(layers []string) (noodles int, sauce float64) {
	for i:=0; i<len(layers); i++ {
		if layers[i]  == "noodles" {
			noodles += 50
		}
		if layers[i]  == "sauce" {
			sauce += 0.2
		}
	}
	return
}
// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendsList, myList []string) {
	lastIdxF := len(friendsList) - 1
	lastIdxM := len(myList) - 1
	myList[lastIdxM] = friendsList[lastIdxF]
}
// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(quantities []float64, numPortions int) []float64 {
        var scaledQuantities []float64
	for i:=0;i<len(quantities);i++ {
		scaledQ := (float64(numPortions) * quantities[i]) / 2
		scaledQuantities = append(scaledQuantities, scaledQ)
	}
	return scaledQuantities
}
// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
