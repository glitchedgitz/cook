package cook

import (
	"fmt"
	"os"
)

func (cook *COOK) ShowCols() {
	fmt.Fprintln(os.Stderr)
	for i, p := range cook.Pattern {
		fmt.Fprintf(os.Stderr, "Col %d: %s\n", i, p)
	}
	fmt.Fprintln(os.Stderr)
	os.Exit(0)
}
