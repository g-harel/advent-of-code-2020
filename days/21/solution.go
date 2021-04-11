package solution

import (
	"sort"
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
			if len(commonIngredients) >= len(commonAllergens) {
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

	ingredientCounts := map[string]int{}
	for _, food := range originalFoods {
		for _, ingredient := range food.ingredients {
			ingredientCounts[ingredient]++
		}
	}

	// Wipe unsafe ingredients from the count.
	for _, food := range solvedFoods {
		if len(food.allergens) > 0 {
			for _, ingredient := range food.ingredients {
				ingredientCounts[ingredient] = 0
			}
		}
	}

	safeCount := 0
	for _, count := range ingredientCounts {
		safeCount += count
	}
	return safeCount
}

func Part2() string {
	originalFoods := ParseFoods(lib.ReadLines("input.txt"))
	solvedFoods := Solve(originalFoods)

	dangerousFoods := []string{}
	for _, food := range solvedFoods {
		if len(food.allergens) == 1 && len(food.ingredients) == 1 {
			dangerousFoods = append(dangerousFoods, food.allergens[0]+":"+food.ingredients[0])
		}
	}

	sort.Strings(dangerousFoods)
	for i, food := range dangerousFoods {
		s := strings.Split(food, ":")
		dangerousFoods[i] = s[1]
	}

	return strings.Join(dangerousFoods, ",")
}
