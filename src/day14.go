package main

import (
    "bufio"
    "flag"
    "fmt"
    "log"
    "math"
    "os"
    "strconv"
    "strings"
)

type Ingredient struct {
    name     string
    quantity int
}

type Recipe struct {
    inputs []Ingredient
    output Ingredient
}

func (r *Recipe) produces(checkIngredient string) int {
    if r.output.name == checkIngredient {
        return r.output.quantity
    }
    return -1
}

var (
    inputFile = flag.String("inputFile", "inputs/day14.txt", "Input File")
    recipes   = make([]*Recipe, 0)
    oreMined  = flag.Int("oreMined", 1000000000000, "The amount of ore that was mined for the purposes of making fuel")
)

func main() {
    flag.Parse()

    processInput()

    fmt.Printf("%d ore needed to produce 1 unit of fuel (1920219 expected)\n", oreNeededToProduce(map[string]int{"FUEL": 1}))
    fmt.Printf("%d units of fuel producible with %d units of ore (1330066 expected)\n", fuelProducibleFrom(*oreMined), *oreMined)
}

func processInput() {
    file, err := os.Open(*inputFile)
    if err != nil {
        log.Fatal(err)
    }
    //noinspection GoUnhandledErrorResult
    defer file.Close()

    // Read in inputs
    scanner := bufio.NewScanner(file)
    recipeList := make([]*Recipe, 0)
    for scanner.Scan() {
        recipeList = append(recipeList, parseRecipe(scanner.Text()))
    }
    recipes = recipeList
}

func parseRecipe(inputString string) *Recipe {
    splitString := strings.Split(inputString, "=>")
    inputs, outputs := strings.Trim(splitString[0], " "), strings.Trim(splitString[1], " ")

    return &Recipe{inputs: listToIngredients(inputs), output: listToIngredients(outputs)[0]}
}

func listToIngredients(list string) []Ingredient {
    ingredients := make([]Ingredient, 0)
    ingredientStrings := strings.Split(list, ",")
    for _, ingredientString := range ingredientStrings {
        constituentParts := strings.Split(strings.Trim(ingredientString, " "), " ")
        quantity, err := strconv.Atoi(constituentParts[0])
        if err != nil {
            log.Fatalf("%s is not an integer\n", constituentParts[0])
        }
        ingredients = append(ingredients, Ingredient{quantity: quantity, name: constituentParts[1]})
    }
    return ingredients
}

func oreNeededToProduce(desiredProducts map[string]int) int {
productionPipeline:
    for {
        for product := range desiredProducts {
            if desiredProducts[product] > 0 && product != "ORE" {
                recipe := recipeThatProduces(product)
                scalar := int(math.Ceil(float64(desiredProducts[product]) / float64(recipe.output.quantity)))
                desiredProducts[product] -= recipe.output.quantity * scalar

                for _, input := range recipe.inputs {
                    desiredProducts[input.name] += input.quantity * scalar
                }
                continue productionPipeline
            }
        }
        return desiredProducts["ORE"]
    }
}

func recipeThatProduces(product string) *Recipe {
    for _, recipe := range recipes {
        if recipe.output.name == product {
            return recipe
        }
    }
    return nil
}

func fuelProducibleFrom(ore int) int {
    lowerBound := 1
    upperBound := math.MaxInt32
    for lowerBound < upperBound {
        check := (lowerBound + upperBound) / 2
        oreProduced := oreNeededToProduce(map[string]int{"FUEL": check})

        if oreProduced < ore && check != lowerBound {
            lowerBound = check
        } else if oreProduced > ore && check != upperBound {
            upperBound = check
        } else {
            return check
        }
    }
    return lowerBound
}
