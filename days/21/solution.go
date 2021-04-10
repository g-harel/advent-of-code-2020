package solution

import (
	"strings"

	"github.com/g-harel/advent-of-code-2020/lib"
)

type Food struct {
	allergens   []string
	ingredients []string
}

func (f Food) IsSolved() bool {
	return len(f.allergens) == 1 && len(f.ingredients) == 1
}

func Solve(foods []Food) []Food {
	for i, fa := range foods {
		for j := i + 1; j < len(foods); j++ {
			fb := foods[j]
			commonAllergens := lib.Common(fa.allergens, fb.allergens)
			commonIngredients := lib.Common(fa.ingredients, fb.ingredients)

			// Remove mislabelled ingredients from foods.
			if len(commonIngredients) == 1 && len(commonAllergens) == 0 {
				if fa.IsSolved() {
					foods[j].ingredients = lib.Remove(foods[j].ingredients, fa.ingredients[0])
					return Solve(foods)
				}
				if fb.IsSolved() {
					foods[i].ingredients = lib.Remove(foods[i].ingredients, fb.ingredients[0])
					return Solve(foods)
				}
			}

			// Nothing to do when common lists are different lengths.
			if len(commonAllergens) != len(commonIngredients) {
				continue
			}

			// Nothing to do when nothing in common.
			if len(commonAllergens) == 0 {
				continue
			}

			// Nothing to do or optimize if everything is common between a and b.
			if len(commonAllergens) == len(fa.allergens) &&
				len(commonAllergens) == len(fb.allergens) &&
				len(commonIngredients) == len(fa.ingredients) &&
				len(commonIngredients) == len(fb.ingredients) {
				continue
			}

			// Found common subset, move to own food and remove from originals.
			foods = append(foods, Food{
				allergens:   commonAllergens,
				ingredients: commonIngredients,
			})
			foods[i].allergens = lib.Remove(foods[i].allergens, commonAllergens...)
			foods[i].ingredients = lib.Remove(foods[i].ingredients, commonIngredients...)
			foods[j].allergens = lib.Remove(foods[j].allergens, commonAllergens...)
			foods[j].ingredients = lib.Remove(foods[j].ingredients, commonIngredients...)
			return Solve(foods)
		}
	}
	return foods
}

func ParseFoods(lines []string) []Food {
	foods := []Food{}
	for _, line := range lines {
		s := strings.Split(line, " (contains ")
		allergens := strings.TrimSuffix(s[1], ")")
		foods = append(foods, Food{
			ingredients: strings.Split(s[0], " "),
			allergens:   strings.Split(allergens, ", "),
		})
	}
	return foods
}

func Part1() int {
	originalFoods := ParseFoods(lib.ReadLines("input.txt"))
	solvedFoods := Solve(originalFoods)

	solvedIngredients := map[string]string{} // ingredient->allergen
	for _, food := range solvedFoods {
		if len(food.allergens) == 1 && len(food.ingredients) == 1 {
			solvedIngredients[food.ingredients[0]] = food.allergens[0]
		}
	}

	mysteryIngredients := map[string]int{}
	for _, food := range originalFoods {
		for _, ingredient := range food.ingredients {
			if _, ok := solvedIngredients[ingredient]; !ok {
				mysteryIngredients[ingredient]++
			}
		}
	}

	mysteryCount := 0
	for _, count := range mysteryIngredients {
		mysteryCount += count
	}
	return mysteryCount
}

func Part2() int {
	lib.ReadLines("input.txt")
	return -1
}
