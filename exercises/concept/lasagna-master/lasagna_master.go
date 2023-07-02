package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, time int) int {
    var result int
    if time == 0{
    	result = 2 * len(layers)
    } else {
    result = time * len(layers)
    }
    return result
}
// TODO: define the 'Quantities()' function
func Quantities(layers []string) (int, float64) {
    noodles := 0
    sauce := 0.0

    for i:=0; i < len(layers); i++ {
        if layers[i] == "noodles" {
            noodles += 50
        }

        if layers[i] == "sauce" {
            sauce += 0.2
        }
    }

    return noodles, sauce
}
// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(layers1 []string, layers2 []string) {
    layers2[len(layers2)-1] = layers1[len(layers1)-1]
}
// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(amounts []float64, porNum int) []float64{
    scaled := make([]float64, len(amounts))
	copy(scaled, amounts)
    for i:=0; i < len(scaled); i++ {
        por := float64(porNum) / 2
        scaled[i] *= por
    }

    return scaled
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
// 
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more 
// functionality.
