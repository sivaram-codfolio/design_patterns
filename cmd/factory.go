package main

import (
	"fmt"

	"github.com/sivaram-codfolio/design_patterns/factory"
)

func main() {
	mag1, _ := factory.NewPublication("magazine", "Tyme", 50, "The Tymes")
	mag2, _ := factory.NewPublication("magazine", "Lyfe", 40, "Lyfe Inc")
	news1, _ := factory.NewPublication("newspaper", "The Herald", 60, "Heralders")
	news2, _ := factory.NewPublication("newspaper", "The Standard", 30, "Standarders")

	pubDetails(mag1)
	pubDetails(mag2)
	pubDetails(news1)
	pubDetails(news2)
}

func pubDetails(pub factory.IPublication) {
	fmt.Printf("--------------------\n")
	fmt.Printf("%s\n", pub)
	fmt.Printf("Type: %T\n", pub)
	fmt.Printf("Name: %s\n", pub.GetName())
	fmt.Printf("Pages: %d\n", pub.GetPages())
	fmt.Printf("Publisher: %s\n", pub.GetPublisher())
	fmt.Printf("--------------------\n")
}
