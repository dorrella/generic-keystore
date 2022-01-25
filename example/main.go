package main

import (
	"fmt"

	keystore "github.com/dorrella/generic-keystore"
)

func main() {
	ks := keystore.NewKeyStore[uint32, string]()

	people := map[uint32]string{
		0: "bob baker",
		1: "drew carey",
		2: "alex trebeck",
	}
	for k, v := range people {
		ks.Put(k, v)
	}

	for k, v := range people {
		val, ok := ks.Get(k)
		fmt.Printf("%s %t %d %s\n", val, ok, k, v)
	}
}
