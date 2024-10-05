package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	data := []float64{}
	for scanner.Scan() {
		nb, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Printf("The input %v is not an int.\n", scanner.Text())
			continue
		}
		data = append(data, float64(nb))
		if len(data) <= 1 {
			continue
		}
		max, min := Guess(data)
		fmt.Println(int(min), int(max))
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Erreur de lecture :", err)
	}
}

func Average(data []float64) float64 {
	var AVG float64
	for _, e := range data {
		AVG += e
	}
	return AVG / float64(len(data))
}

func Median(data []float64) float64 {
	var Date_sorted []float64
	Date_sorted = append(Date_sorted, data...)
	Date_sorted = Sort(Date_sorted)
	var MD float64
	if len(Date_sorted)%2 != 0 {
		MD = Date_sorted[len(data)/2]
	} else {
		MD = (Date_sorted[len(Date_sorted)/2] + Date_sorted[((len(Date_sorted)/2)-1)]) / 2
	}
	return MD
}

func Sort(sl []float64) []float64 {
	for i := 0; i < len(sl)-1; i++ {
		for j := i + 1; j < len(sl); j++ {
			if sl[i] > sl[j] {
				sl[i], sl[j] = sl[j], sl[i]
			}
		}
	}
	return sl
}

func StandardDeviation(data []float64) float64 {
	return math.Sqrt(float64(Variance(data)))
}

func Variance(data []float64) float64 {
	var VR float64
	for _, e := range data {
		x := e - Average(data)
		VR += Pow(x)
	}
	return VR / float64(len(data))
}

func Pow(n float64) float64 {
	p := 1.0
	for i := 1; i <= 2; i++ {
		p = p * n
	}
	return p
}

func Guess(data []float64) (int, int) {
	var max, min float64
	start := len(data) - 4
	if start > 0 {
		newSl := data[start:]
		max = Average(newSl) + StandardDeviation(newSl)
		min = Average(newSl) - StandardDeviation(newSl)
	} else {
		max = Average(data) + StandardDeviation(data)
		min = Average(data) - StandardDeviation(data)
	}
	if min < 0 {
		min *= -1
	}
	return int(max), int(min)
}
