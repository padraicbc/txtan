package txtan

import (
	"math"
	"runtime"
)

var cpus = runtime.GOMAXPROCS(0)

func Dot(a, b []int) float64 {
	// probably better to set a min element size but ok for now..
	if len(a) < cpus {
		return _dot(a, b)
	}
	var dp float64
	c := make(chan float64, cpus)
	for i := 0; i < cpus; i++ {
		go dotc(i*len(a)/cpus, (i+1)*len(a)/cpus, a, b, c)
	}
	for i := 0; i < cpus; i++ {
		dp += <-c

	}

	return dp
}

func dotc(i, end int, a, b []int, c chan float64) {
	var dp float64
	for ; i < end; i++ {
		dp += float64(a[i] * b[i])
	}
	c <- dp

}
func _dot(a, b []int) float64 {
	var dp float64
	for i, j := range a {
		dp += float64(j * b[i])
	}
	return dp

}

func Norm(a []int) float64 {

	if len(a) < cpus {
		return _norm(a)
	}
	c := make(chan float64, cpus)
	for i := 0; i < cpus; i++ {
		go normc(i*len(a)/cpus, (i+1)*len(a)/cpus, a, c)
	}
	var o float64

	for i := 0; i < cpus; i++ {
		o += <-c
	}

	return math.Sqrt(o)
}
func _norm(a []int) float64 {
	var o float64
	for _, i := range a {
		o += math.Pow(float64(i), 2)
	}
	return math.Sqrt(o)
}

func normc(start, end int, a []int, c chan float64) {
	var o float64
	for i := start; i < end; i++ {
		o += math.Pow(float64(a[i]), 2)
	}
	c <- o

}
