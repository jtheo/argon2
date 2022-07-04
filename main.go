package main

import (
	"fmt"
)

func main() {
	argon, config := start()

	if config.encoded != "" {
		if verifyPass(config) {
			fmt.Printf("%s Password match! %s\n", working, match)
			return
		}
		fmt.Printf("%s Password don't match! %s\n", boom, noMatch)
		return
	}

	if config.verbose {
		printHeaders(argon)
	}

	fmt.Println(encrypt(config.pass, argon))
}
