package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, avgPreparationTime int) int {
    if avgPreparationTime > 0 {
        return avgPreparationTime * len(layers)
    }
    return 2 * len(layers)
}

// TODO: define the 'Quantities()' function
func Quantities(layers []string) (noodles int, sauce float64) {
    for _, layer := range layers {
        if layer == "noodles" {
            noodles += 50
        }
        if layer == "sauce" {
            sauce += 0.2
        }
    }
    return
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendList []string, myList []string) {
    myList[len(myList) - 1] = friendList[len(friendList) - 1]
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(quantities []float64, portions int) []float64 {
    var scaled []float64

    for i := range quantities {
        scaled = append(scaled,(quantities[i] / 2) * float64(portions))
    }

    return scaled
}
