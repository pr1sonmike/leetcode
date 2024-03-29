package _431_kids_with_the_greatest_number_of_candies

func kidsWithCandies(candies []int, extraCandies int) []bool {

	var max = candies[0]
	var result = make([]bool, len(candies))

	for i := range candies {
		if candies[i] > max {
			max = candies[i]
		}
	}

	for i := range candies {
		if candies[i]+extraCandies >= max {
			result[i] = true
		}
	}

	return result
}
