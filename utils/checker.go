package lemin

import (
	"fmt"
	"os"
)

func Checker(c Colony) {
	if c.Start == "" || c.Ants <= 0 || c.End == "" {
		fmt.Println("ERROR: invalid data format")
		os.Exit(0)
	}
	if len(Allpaths) == 0 {
		fmt.Println("ERROR: invalid data format")
		os.Exit(0)
	}
}
