package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	args := os.Args[1:]

	if len(args) < 2 {
		log.Fatalf("not enough args provided, check usage")
	}

	referral_id := ResolveReferralLink(args[0])
	amount, _ := strconv.Atoi(args[1])

	fmt.Println("starting to add", amount)

	for range amount {
		registerClient := WarpClient{}

		data := registerClient.Register()
		success := registerClient.PatchReferrer(referral_id)

		if success {
			fmt.Printf("[$] added referral -> %s\n", data.ID)
		} else {
			fmt.Printf("failed to add -> %s, skipping...\n", data.ID)
		}
	}

	fmt.Println("thanks for using warper!!")
}
