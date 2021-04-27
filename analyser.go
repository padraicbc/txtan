package txtan

import (
	"sync"
)

// tokens are passed, the rest are calculated
type Analyser struct {
	tokensA []string
	tokensB []string
	countsA map[string]int
	countsB map[string]int
	union   map[string]bool
}

func cosineSimilarity(l1, l2 []int) float64 {
	return Dot(l1, l2) / (Norm(l1) * Norm(l2))
}

func (a *Analyser) CosineSimilarity() float64 {
	l1, l2 := []int{}, []int{}
	// get the 2 slices of intersecting/non-intersectig words
	// missing key will be zero value which is 0.
	for k := range a.union {
		l1 = append(l1, a.countsA[k])
		l2 = append(l2, a.countsB[k])

	}

	return cosineSimilarity(l1, l2)
}

func jaccardSimilarity(interCnt, unionCnt float64) float64 {
	return interCnt / unionCnt
}

// calculates JaccardSimilarity of passed in tokens
func (a *Analyser) JaccardSimilarity() float64 {
	return jaccardSimilarity(a.intersect(), float64(len(a.union)))
}

func (a *Analyser) intersect() float64 {
	var in float64
	for k := range a.countsA {
		if a.countsB[k] > 0 {
			in++
		}
	}
	return in
}

// count of union terms
func (a *Analyser) unions() {
	un := map[string]bool{}
	for _, mp := range []map[string]int{a.countsA, a.countsB} {
		for k := range mp {
			un[k] = true
		}
	}
	a.union = un
}

// do each count in separate go routines
func (a *Analyser) counts() {
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		cn := map[string]int{}
		for _, wo := range a.tokensA {
			cn[wo]++
		}

		a.countsA = cn
		wg.Done()
	}()
	go func() {
		cn := map[string]int{}
		for _, wo := range a.tokensB {
			cn[wo]++
		}
		a.countsB = cn
		wg.Done()
	}()
	wg.Wait()

}

func Setup(a, b []string) *Analyser {

	an := &Analyser{
		tokensA: a,
		tokensB: b}
	an.counts()
	an.unions()
	return an

}
