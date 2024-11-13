package main

import (
	"fmt"
	"math/big"
	"strconv"
)

func main() {
	in := *big.NewInt(0b10011111011011001)
	size := 35651584 // 272
	//in := *big.NewInt(0b10000)
	//size := 20
	length := in.BitLen()
	
	fmt.Printf("In: %s\n", in.Text(2))
	for length < size {
		in, length = iterate(in, length)
	}
	fmt.Printf("Sized: (%d) %s\n", length, in.Text(2))


	in = *in.Rsh(&in, uint(length - size))
	fmt.Printf("Cropped: (%d) %s\n", length, in.Text(2))
	length = size

	in, length = checksum(in, length)
	fmt.Printf("Checksum: (%d) %s\n", length, in.Text(2))

	fmt.Printf("RESULT: [%"+strconv.Itoa(length)+"s]", in.Text(2))
}


func iterate(a big.Int, length int) (big.Int, int) {
	//fmt.Printf("a (in, %d): %s\n", length, a.Text(2))
	b := big.NewInt(0)
	for i := 0; i < length; i++ {
		b = b.SetBit(b, i, uint(-1*int(a.Bit(length-i-1)-1)))
	}
	//fmt.Printf("b (rev/tog): %s\n", b.Text(2))
	a = *a.Lsh(&a, uint(length+1))
	//fmt.Printf("a (lsh): %s\n", a.Text(2))
	a = *a.Add(&a, b)
	//fmt.Printf("a (add): %s\n", a.Text(2))
	return a, length * 2 + 1
}

func checksum(a big.Int, len int) (big.Int, int) {
	//fmt.Printf("Checksum: a: (%d/%d), %s\n", len, a.BitLen(), a.Text(2))
	sum := big.NewInt(0)
	for i := 0; i < len; i += 2 {
		if a.Bit(i) == a.Bit(i+1) {
			sum = sum.SetBit(sum, i/2, 1)
		}
	}
	//fmt.Printf("  sum: (%d), %s\n", len/2, sum.Text(2))
	if len/2 % 2 == 0 {
		return checksum(*sum, len/2)
	} else {
		return *sum, len/2
	}
}
