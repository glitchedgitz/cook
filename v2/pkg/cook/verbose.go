package cook

import (
	"fmt"
	"os"
)

func (cook *COOK) VPrint(msg string) {
	if cook.Config.Verbose {
		fmt.Fprintln(os.Stderr, msg)
	}
}
