package main

import (
	"fmt"
	"os"
)

func exitWithWait(msg string) {
	fmt.Println(msg)
	fmt.Println("Press any key to countinue..")
	fmt.Scanln()
	os.Exit(0)
}
