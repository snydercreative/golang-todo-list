package util

import "log"

// Check automatically fails errors. Syntactic sugar.
func Check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
