package main

import (
	"fmt"

	"github.com/nstoker/golang_integration_test_db/app/version"
)

func main() {
	fmt.Printf("initial build %s\n", version.Version)
}
