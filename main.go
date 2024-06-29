package main

import (
	"fmt"
	"github.com/AashishNandakumar/cryptit/decrypt"
	"github.com/AashishNandakumar/cryptit/encrypt" // importing a package inside a module
)

func main() {
	fmt.Println(encrypt.Nimbus("AashishNandakumar"))
	fmt.Println(decrypt.Nimbus("DdvklvkQdqgdnxpdu"))
}
