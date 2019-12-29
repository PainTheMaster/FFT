package main

import (
	"PainTheMaster/FFT/fft"
)

func main() {
	/*
		var idx = make([]int, 8)

		for i := range idx {
			idx[i] = i
			fmt.Println(i, ",", idx[i])
		}

		fmt.Println()

		for i := range idx {
			fft.Bitreverse(&idx[i], 3)
			fmt.Println(i, ",", idx[i])
		}*/
	fft.Test()
}
