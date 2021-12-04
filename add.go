package main

func Add(numbers [5]int) (sum int) {

	for _, number := range numbers {
		sum += number
	}

	return
}
