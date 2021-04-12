package main

import (
	"fmt"
	"github.com/pkg/errors"
	"strconv"
)

func getAssetBalance(asset string) (float64, error) {
	resp, err := krakenClient.Query("Balance", map[string]string{"pair": asset})
	if err != nil {
		return 0, errors.Wrap(err, "failed to get account balance")
	}
	balanceStr, ok := resp.(map[string]interface{})[asset]

	if !ok {
		return 0, errors.New(fmt.Sprintf("you don't have %v in your balance", asset))
	}

	balance, err := strconv.ParseFloat(balanceStr.(string), 64)
	if err != nil {
		return 0, errors.New("can't parse balance")
	}

	return balance, nil
}
