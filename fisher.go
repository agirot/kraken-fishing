package main

import (
	"fmt"
)

func sell(target, currentClosedPrice float64) bool {
	//Target reached
	if currentClosedPrice >= target {
		if previousClosedPrice == 0 {
			fmt.Println("target reached ! Now wait the best price")
		} else {
			//Check IF SELL
			if previousClosedPrice > currentClosedPrice {
				//CHECK IF HOLD
				if currentHoldCount >= maxHoldCount  {
					fmt.Println("Don't wait anymore, SELL !")
					return true
				} else {
					currentHoldCount = currentHoldCount + 1
					fmt.Println("Args ! The last price is lower than the previous one, hold on...")
				}
			} else {
				fmt.Println("Good, the current price is better than the previous, wait")
			}
		}
		previousClosedPrice = currentClosedPrice
	} else {
		currentHoldCount = 0
		previousClosedPrice = 0
	}
	return false
}

func buy(target, currentClosedPrice float64) bool {
	//Target reached
	if currentClosedPrice <= target {
		if previousClosedPrice == 0 {
			fmt.Println("target reached ! Now wait the best price")
		} else {
			//Check IF BUY
			if previousClosedPrice < currentClosedPrice {
				//CHECK IF HOLD
				if currentHoldCount >= maxHoldCount  {
					fmt.Println("Don't wait anymore, BUY !")
					return true
				} else {
					currentHoldCount = currentHoldCount + 1
					fmt.Println("Args ! The last price is higher than the previous one, hold on...")
				}
			} else {
				fmt.Println("Good, the current price is better than the previous, wait")
			}
		}
		previousClosedPrice = currentClosedPrice
	} else {
		currentHoldCount = 0
		previousClosedPrice = 0
	}
	return false
}
