package main

import (
	"testing"
)

func BenchmarGetkGetEuclideanDistance(b *testing.B){
	a := [2]int{2,3}
	bb := [2]int{7,9}

	for i := 0; i < b.N; i++{
		geteuclideanDistance(a, bb)
	}
}