package lasagnamaster

// PreparationTime estimates the total preparation time based on the number of layers.
// If the average preparation time is 0, it defaults to 2 minutes.
func PreparationTime(layers []string, avgPrepTime int) int {
	if avgPrepTime == 0 {
		avgPrepTime = 2
	}
	return len(layers) * avgPrepTime
}

// Quantities determines the amount of noodles (in grams) and sauce (in liters) needed.
// Each noodle layer requires 50g, and each sauce layer requires 0.2L.
func Quantities(layers []string) (int, float64) {
	noodles := 0
	sauce := 0.0

	for _, layer := range layers {
		if layer == "noodles" {
			noodles += 50
		} else if layer == "sauce" {
			sauce += 0.2
		}
	}
	return noodles, sauce
}

// AddSecretIngredient copies the last ingredient from your friend's recipe
// and uses it to replace the last element of your ingredient list.
func AddSecretIngredient(friendsList []string, myList []string) {
	secretIngredient := friendsList[len(friendsList)-1]
	myList[len(myList)-1] = secretIngredient
}

// ScaleRecipe calculates the required amounts of ingredients for the desired number of portions.
// The input amounts slice assumes the recipe is scaled for exactly 2 portions.
func ScaleRecipe(amounts []float64, portions int) []float64 {
	scaled := make([]float64, len(amounts))
	factor := float64(portions) / 2.0

	for i, amount := range amounts {
		scaled[i] = amount * factor
	}
	return scaled
}
