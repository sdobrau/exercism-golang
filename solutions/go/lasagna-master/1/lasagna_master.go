package lasagnamaster

// TODO: define the 'PreparationTime()' function
func PreparationTime (layers []string, perSliceTime int) int {
    if perSliceTime == 0 {
        perSliceTime = 2
    }
    return perSliceTime * len(layers)
} 
// TODO: define the 'Quantities()' function
func Quantities(layers []string) (noodles int, sauce float64) {
    for i := 0; i < len(layers); i++ {
        if layers[i] == "noodles" {
            noodles += 50
        } else if layers[i] == "sauce" {
            sauce += 0.2
        }
    }
    return
}
// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(ingredients_1 []string, ingredients_2 []string) {
    ingredients_2[len(ingredients_2)- 1] = ingredients_1[len(ingredients_1) - 1] 
}
// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(amounts []float64, nr_portions int) []float64 {
    newSlice := make([]float64, len(amounts))
    for i := 0; i < len(amounts); i++ {
        newSlice[i] = amounts[i] * (float64(nr_portions) / 2.0)
    }
    return newSlice
}
// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
