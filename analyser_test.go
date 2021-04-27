package txtan_test

import (
	"fmt"
	"strings"

	"github.com/padraicbc/txtan"
)

func ExampleSetup() {
	aa, bb := `These sentences are identical so should be equal to one for both`, `These sentences are identical so should be equal to one for both`

	an := txtan.Setup(strings.Fields(aa), strings.Fields(bb))

	fmt.Printf("%.0f %.0f\n", an.CosineSimilarity(), an.JaccardSimilarity())

	aa = "Data is the new oil of the digital economy"
	bb = "Data is a new oil"
	an = txtan.Setup(strings.Fields(aa), strings.Fields(bb))

	fmt.Printf("%.2f %.2f\n", an.CosineSimilarity(), an.JaccardSimilarity())
	// Output: 1 1
	// 0.54 0.44

}
