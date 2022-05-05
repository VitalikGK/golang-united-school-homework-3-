package homework

func reverse(input []int64) (result []int64) {
	for i := int64(len(input)/2 - 1); i >= 0; i-- {
		opp := int64(len(input)) - 1 - i
		input[i], input[opp] = input[opp], input[i]

		//	fmt.Printd("%d", input[])
	}
	return input
}
