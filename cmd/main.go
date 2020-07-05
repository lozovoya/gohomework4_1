package main

import (
	"fmt"
	"github.com/lozovoya/gohomework4_1/pkg/card"
)

func main() {

	banks := make([]card.Service, 0, 5)
	banks = append(banks, *card.NewService("Penguin Bank"))
	banks = append(banks, *card.NewService("Zebra Bank"))
	banks = append(banks, *card.NewService("Cucumber Bank"))

	banks[0].IssueCard("master", 100_000_00, "0000 0000 0000 0000", "rub")
	banks[0].IssueCard("visa", 100_000_00, "1111 1111 1111 1111", "rub")

	banks[1].IssueCard("master", 10_000_00, "2222 2222 2222 2222", "rub")
	banks[1].IssueCard("visa", 15_000_00, "3333 3333 3333 3333", "rub")

	banks[2].IssueCard("master", 50_000_00, "4444 4444 4444 4444", "rub")
	banks[2].IssueCard("visa", 60_000_00, "5555 5555 5555 5555", "rub")

	fmt.Println(banks[2].Cards[1])
}
