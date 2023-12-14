package main

import (
	"fmt"
	"log"
	"math"
)

var (
	errStep = fmt.Errorf("the step must be a positive number")
)

func DefiningGroup(temperatures []float64, step float64) (map[int][]float64, error) {
	if step <= 0 {
		return nil, errStep
	}
	temperatureIntervals := make(map[int][]float64)
	for _, temperature := range temperatures {
		switch {
		case temperature < step && temperature >= 0:
			temperatureIntervals[1] = append(temperatureIntervals[1], temperature)
		case math.Abs(temperature) < step && temperature < 0:
			temperatureIntervals[-1] = append(temperatureIntervals[-1], temperature)
		default:
			wholePart := int(temperature/step) * int(step)
			temperatureIntervals[wholePart] = append(temperatureIntervals[wholePart], temperature)
		}
	}
	return temperatureIntervals, nil
}
func main() {
	temperatures := []float64{-25.4, -27.0, 13.0, -1.1, 1.1, 0, 19.0, 15.5, 24.5, -21.0, 32.5}
	step := 10.0
	temperatureIntervals, err := DefiningGroup(temperatures, step)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(temperatureIntervals)
}
