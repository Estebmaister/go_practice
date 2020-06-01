package logicflow

import (
	"fmt"
	"math/rand"
	"time"
)

// Heist make a simulation of a heist with random events
func Heist() {
	rand.Seed(time.Now().UnixNano())
	var isHeistOn bool
	isHeistOn = true
	eludedGuards := rand.Intn(100)

	if eludedGuards >= 50 {
		fmt.Println("Looks like you've managed to make it past the guards. Good job, but remember, this is the first step.")
	} else {
		isHeistOn = false
		fmt.Println("Plan a better disguise next time?")
	}

	openedVault := rand.Intn(100)

	if isHeistOn && openedVault >= 70 {
		fmt.Println("Vault is open! Grab and GO!!")
	} else if isHeistOn {
		isHeistOn = false
		fmt.Println("We couldn't open the vault this time")
	}

	leftSafely := rand.Intn(5)

	if isHeistOn {
		switch leftSafely {
		case 0:
			isHeistOn = false
			fmt.Println("Looks like you tripped an alarm... run?")
		case 1:
			isHeistOn = false
			fmt.Println("Turns out vault doors don't open from the inside...")
		case 2:
			isHeistOn = false
			fmt.Println("When did they start raising dogs in vaults??")
		case 3:
			isHeistOn = false
			fmt.Println("Looks like this fingerprint scanner won’t accept any fingerprint…")
		default:
			fmt.Println("Start the getaway car!")
		}
	}

	if isHeistOn {
		amtStolen := 10000 + rand.Intn(1000000)
		fmt.Println("$", amtStolen, "not bad!")
	}

	fmt.Println("Heist is in progress:", isHeistOn)
}
