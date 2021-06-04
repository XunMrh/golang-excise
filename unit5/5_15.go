package main

import (
	"errors"
	"fmt"
	"log"
)

func main5_15() {
	mlist := []int{1, 2, 3, 4}
	max_, err := max(mlist...)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("max: %d\n", max_)
	min_, err := min(mlist...)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("min: %d\n", min_)
}

func max(vars ...int) (int, error) {
	if len(vars) == 0 {
		return 0, errors.New("empty")
	}
	tempMax := vars[0]
	for _, v := range vars {
		if tempMax < v {
			tempMax = v
		}
	}
	return tempMax, nil
}
func min(vars ...int) (int, error) {
	if len(vars) == 0 {
		return 0, errors.New("empty")
	}
	tempMin := vars[0]
	for _, v := range vars {
		if tempMin > v {
			tempMin = v
		}
	}
	return tempMin, nil
}
