package cook

import (
	"fmt"
	"os"
)

func (cook *COOK) VPrint(msg string) {
	if cook.Verbose {
		fmt.Fprintln(os.Stderr, msg)
	}
}
