package colorHelper

import (
	"errors"
	"fmt"
	"math"
)

func (c rgbColor) rgbToHue() float64 {
	var s float64
	// Algorithm from Numa.nu

	// determinate the min max value
	rV := float64(c.red) / 255
	gV := float64(c.green) / 255
	bV := float64(c.blue) / 255

	colorFloatArray := []float64{rV, gV, bV}
	min, max := getMinMax(colorFloatArray)

	// If min and max is equal it mean that there's no saturation
	// Therefore no hue
	if min == max {
		return 0
	}

	luminace := (min + max) / 2
	if luminace < 0.5 {
		s = (max - min) / (max + min)
	} else {
		s = (max - min) / (2 - (max - min))
	}

	// Get the color which it's value is higher than the other
	maxColorName := getMaxColor(max, colorFloatArray)
	hue, _ := calcHue(maxColorName, colorFloatArray, max, min)
	fmt.Println(hue)
	fmt.Println(s)

	return 1

}

// GetMinMax
//		* Get the minimum value and the maximum value between an array of float64
// --> colorValue []float64
// @ float64, float64
func getMinMax(colorValue []float64) (float64, float64) {
	var (
		min float64
		max float64
	)

	// Get the min and max value from an array of float64 value
	for i := 0; i < len(colorValue); i++ {
		for j := 0; j < len(colorValue); j++ {
			tempMin := math.Min(colorValue[i], colorValue[j])
			tempMax := math.Max(colorValue[i], colorValue[j])

			if tempMin < min || min == 0 {
				min = tempMin
			}

			if tempMax > max {
				max = tempMax
			}
		}
	}

	return min, max
}

func getMaxColor(maxValue float64, colorValue []float64) string {

	for i := 0; i < len(colorValue); i++ {
		if colorValue[i] == maxValue {
			if i == 0 {
				return "red"
			} else if i == 1 {
				return "green"
			} else {
				return "blue"
			}
		}
	}

	return ""
}

// calcHue
//		* Calculate the hue value
// --> colorName string
// --> colorValue []float64
// --> max float64
// --> min float64
// @ int
func calcHue(colorName string, colorValue []float64, max float64, min float64) (int, error) {
	var hue float64
	switch colorName {
	case "red":
		hue = (colorValue[1] - colorValue[2]) / (max - min)
	case "green":
		hue = 2 + ((colorValue[2] - colorValue[0]) / (max - min))
	case "blue":
		hue = 4 + ((colorValue[0] - colorValue[1]) / (max - min))
	default:
		return 0, errors.New("no colorname provide")
	}

	fmt.Println(hue)
	hue = hue * 60

	if hue < 0 {
		hue += 360
	}

	return int(hue), nil
}
