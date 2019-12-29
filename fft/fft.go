package fft

import (
	"fmt"
	"math"
)

var wRev, wForw [][]complex128

// FFT performs fast fourier transformation
func FFT(f []float64) (F []float64) {

	zeroFill(&f)
	bitWidth := uint(bitWidth(len(f)))
	prepOmega(1 << (bitWidth - 1))

	shuffle := make([]int, len(f))
	for i := range shuffle {
		shuffle[i] = bitreverse(i, int(bitWidth))
	}

	return
}

func zeroFill(f *[]float64) {
	var origBitWidth, newBitWidth uint
	var newLength int
	var temp []float64

	/*	for ; (len(*f) >> origBitWidth) > 0; origBitWidth++ {
		}
	*/
	origBitWidth = uint(bitWidth(len(*f)))

	if len(*f) == (1 << (origBitWidth - 1)) {
	} else {
		newBitWidth = origBitWidth + 1
		newLength = 1 << (newBitWidth - 1)
		temp = make([]float64, newLength-len(*f))
		*f = append(*f, temp...)
	}

}

//Bitreverse is a function that rearranges the index for FFT
func bitreverse(x int, bitWidth int) int {
	//	ux := uint32(x)
	var k uint
	var uBitWidth = uint(bitWidth)
	var temp int

	for k = 0; k <= uBitWidth-1; k++ {
		temp |= (x & 1) << (uBitWidth - k - 1)
		x = x >> 1
	}

	return temp
}

//prepOmega prepares table for wRev and wForw
//argument "size" is literalli a "size", normally 1, 2, 4, 8..., and NOT a logarithm of them
//wForw[i] corresponds to a series of 2^i th root of unity, and contains 2^i datum (w^0 to w^(i-1))
func prepOmega(size int) {

	var numBitSize int

	/*	for ; (size >> uint(numBitSize)) > 0; numBitSize++ {
		}*/

	numBitSize = bitWidth(size)

	wForw = make([][]complex128, numBitSize)
	wRev = make([][]complex128, numBitSize)

	for i := 0; 1<<uint(i) <= size; i++ {

		thisDivision := 1 << uint(i)
		wForw[i] = make([]complex128, thisDivision)
		wRev[i] = make([]complex128, thisDivision)

		unitRad := 2.0 * math.Pi / float64(thisDivision)
		for j := 0; j <= thisDivision-1; j++ {
			wForw[i][j] = complex(math.Cos(unitRad*float64(j)), math.Sin(unitRad*float64(j)))
			wRev[i][j] = complex(math.Cos(unitRad*float64(j)), -math.Sin(unitRad*float64(j)))
		}
	}
}

//bitWidth returns bit width of an integer x. ex) bitwidth(4)==3
func bitWidth(x int) int {
	var width uint
	for width = 0; (x >> width) > 0; width++ {
	}

	return int(width)
}

//Test is an exported test field
func Test() {
	fmt.Println("bitreverse(3,4):", bitreverse(3, 4))
}

/*func Test() {
	//	f8, f10 := make([]float64, 8), make([]float64, 10)

	f8 := []float64{1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 7.0, 8.0}
	f10 := []float64{10.0, 20.0, 30.0, 40.0, 50.0, 60.0, 70.0, 80.0, 90.0, 100.0}

	fmt.Println("len(f8)", len(f8))
	fmt.Println("len(f10)", len(f10))

	zeroFill(&f8)
	zeroFill(&f10)
	fmt.Println("zerofilled")
	fmt.Println("len(f8)", len(f8))
	fmt.Println("len(f10)", len(f10))

}*/

/*func Test() {
	size := 8

	prepOmega(size)

	fmt.Println("len(wForw):", len(wForw))

	for i := 0; i <= len(wForw)-1; i++ {
		fmt.Println("i:", i)
		fmt.Println("wForw len:", len(wForw[i]), ", wRev len:", len(wRev[i]))
		for j := 0; j <= len(wForw[i])-1; j++ {
			fmt.Println(wForw[i][j], ", ", wRev[i][j])
		}
	}
}*/
