package main

import (
	"fmt"
	"log"
)

func sell(target, currentClosedPrice float64) bool {
	//Target reached
	if currentClosedPrice >= target {
		if previousClosedPrice == 0 {
			log.Println("target reached ! Now wait the best price")
		} else {
			//Check IF SELL
			if previousClosedPrice > currentClosedPrice {
				//CHECK IF HOLD
				if currentHoldCount >= maxHoldCount  {
					log.Println("Don't wait anymore, SELL !")
					return true
				} else {
					currentHoldCount = currentHoldCount + 1
					log.Println(fmt.Sprintf("Args ! The current price (%v) is lower than the previous one (%v), hold on...", currentClosedPrice, previousClosedPrice))
				}
			} else {
				log.Println(fmt.Sprintf("Good, the current price (%v) is better than the previous one (%v), wait", currentClosedPrice, previousClosedPrice))
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
			log.Println("target reached ! Now wait the best price")
		} else {
			//Check IF BUY
			if previousClosedPrice < currentClosedPrice {
				//CHECK IF HOLD
				if currentHoldCount >= maxHoldCount  {
					log.Println("Don't wait anymore, BUY !")
					return true
				} else {
					currentHoldCount = currentHoldCount + 1
					log.Println(fmt.Sprintf("Args ! The current price (%v) is higher than the previous one (%v), hold on...", currentClosedPrice, previousClosedPrice))
				}
			} else {
				log.Println(fmt.Sprintf("Good, the current price (%v) is better than the previous (%v), wait", currentClosedPrice, previousClosedPrice))
			}
		}
		previousClosedPrice = currentClosedPrice
	} else {
		currentHoldCount = 0
		previousClosedPrice = 0
	}
	return false
}
