package main

import (
	"fmt"
	"github.com/agirot/kraken-go-api-client"
	"github.com/pkg/errors"
	"github.com/urfave/cli/v2"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

var app *cli.App

var krakenClient *krakenapi.KrakenAPI

//Dynamic data
var previousClosedPrice float64
var currentHoldCount int

//CLI FLAGS
var testFlag string
var test = true
var tickFlag string
var tick time.Duration
var holdCountFlag string
var maxHoldCount int
var pathFlag string

//Config file
type config struct {
	Key    string
	Secret string
}

var configFile config

func init() {
	log.SetFlags(log.Ldate | log.Ltime)

	app = &cli.App{
		Name:        "Catch The Kraken !",
		Version:     "v0.1",
		Description: "Optimize your Buy or sell order",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "config",
				Value:       "config.yml",
				Usage:       "Set path of config file",
				Destination: &pathFlag,
			},
		},
		Before: func(context *cli.Context) error {
			f, err := os.OpenFile(pathFlag, os.O_RDONLY, 0444)
			if err != nil {
				log.Fatalf("failed to load %v file: %v", pathFlag, err)
			}

			b, err := ioutil.ReadAll(f)
			if err != nil {
				log.Fatalf("failed to load %v file: %v", pathFlag, err)
			}

			err = yaml.Unmarshal(b, &configFile)
			if err != nil {
				log.Fatalf("failed to load %v file: %v", pathFlag, err)
			}

			krakenClient = krakenapi.New(
				configFile.Key,
				configFile.Secret,
			)

			return nil
		},
	}
}

func main() {
	app.Commands = []*cli.Command{
		{
			Name: "sell",
			Usage: `sell <pair> <volume> <target>
					pair: asset pair to sell
					volume: volume of order (you can set 'all')
					target: price to sell
				`,
			HelpName: "sell with a minimal target",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:        "test",
					Value:       "yes",
					Usage:       "Set yes to achieve real orders",
					Destination: &testFlag,
				},
				&cli.StringFlag{
					Name:        "tick",
					Value:       "15m",
					DefaultText: "15m",
					Usage:       "Set 30/15m/1h to set the price check interval",
					Destination: &tickFlag,
				},
				&cli.StringFlag{
					Name:        "hold",
					Value:       "2",
					Usage:       "Set the count of negative evolution before sell/buy",
					Destination: &holdCountFlag,
				},
			},
			Before: func(context *cli.Context) error {
				if testFlag == "no" {
					test = false
				}

				var err error
				tick, err = time.ParseDuration(tickFlag)
				if err != nil {
					log.Fatal("invalid tick flag value")
				}

				maxHoldCount, err = strconv.Atoi(holdCountFlag)
				if err != nil {
					log.Fatal("invalid tick flag value")
				}
				return nil
			},
			Action: func(c *cli.Context) error {
				if c.Args().Get(0) == "" || c.Args().Get(1) == "" || c.Args().Get(2) == "" {
					return errors.New("you must set 3 args, see help command")
				}

				target, err := strconv.ParseFloat(c.Args().Get(2), 64)
				if err != nil {
					return errors.New("target must be a float")
				}

				pairName := c.Args().Get(0)
				pair, err := splitCryptoNameInPair(pairName)
				if err != nil {
					return err
				}

				volume := c.Args().Get(1)
				balance, err := getAssetBalance(pair.Origin)
				if err != nil {
					return err
				}

				if volume == "all" {
					volume = strconv.FormatFloat(balance, 'f', 8, 64)
				} else {
					volumeFl, err := strconv.ParseFloat(volume, 64)
					if err != nil {
						return errors.New("can't parse volume arg")
					}

					if volumeFl > balance {
						return errors.New(fmt.Sprintf("you can't sell %v %v you have only %f", volumeFl, pair.Origin, balance))
					}

				}

				ticker := time.NewTicker(tick)
				fmt.Printf("Start to fishing %v ! target to %v with %v quantities (hold %v times and check every %v)\n",
					pairName, target, volume, maxHoldCount, tick)
				fmt.Printf("DEMO mode: %v\n", test)
				for _ = range ticker.C {
					resp, err := krakenClient.Ticker(pairName)
					if err != nil {
						return err
					}

					currentClosedPrice, err := strconv.ParseFloat(resp.FLOWEUR.Close[0], 64)
					if err != nil {
						log.Println(err)
						continue
					}
					sellAction := sell(target, currentClosedPrice)
					if sellAction {
						if !test {
							order, err := krakenClient.AddOrder(pairName, "sell", "market", c.Args().Get(1), nil)
							if err != nil {
								return err
							}
							log.Printf("SELL %v\n", order.Description.PrimaryPrice)
							return nil
						} else {
							log.Printf("SELL simulation %v\n", currentClosedPrice)
						}
						return nil
					}
				}
				return nil
			},
		},
		{
			Name: "buy",
			Usage: `buy <pair> <volume> <target>
					pair: asset pair to buy
					volume: volume of order
					target: price to buy
				`,
			HelpName: "buy with a maximal target",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:        "test",
					Value:       "yes",
					Usage:       "Set yes to achieve real orders",
					Destination: &testFlag,
				},
				&cli.StringFlag{
					Name:        "tick",
					Value:       "5",
					DefaultText: "5",
					Usage:       "Set 30/15m/1h to set the price check interval",
					Destination: &tickFlag,
				},
				&cli.StringFlag{
					Name:        "hold",
					Value:       "2",
					Usage:       "Set --hold=<count> to set the count of negative value before sell/buy",
					Destination: &holdCountFlag,
				},
			},
			Before: func(context *cli.Context) error {
				if testFlag == "no" {
					test = false
				}

				var err error
				tick, err = time.ParseDuration(tickFlag)
				if err != nil {
					log.Fatal("invalid tick flag value")
				}

				maxHoldCount, err = strconv.Atoi(holdCountFlag)
				if err != nil {
					log.Fatal("invalid tick flag value")
				}
				return nil
			},
			Action: func(c *cli.Context) error {
				if c.Args().Get(0) == "" || c.Args().Get(1) == "" || c.Args().Get(2) == "" {
					return errors.New("you must set 3 args, see help command")
				}

				target, err := strconv.ParseFloat(c.Args().Get(2), 64)
				if err != nil {
					return errors.New("target must be a float")
				}

				ticker := time.NewTicker(tick)
				pair := c.Args().Get(0)
				fmt.Printf("Start to fishing %v ! target to %v with %v quantities (hold %v times and check every %v)\n",
					pair, target, c.Args().Get(1), maxHoldCount, tick)
				fmt.Printf("DEMO mode: %v\n", test)

				for _ = range ticker.C {
					resp, err := krakenClient.Ticker(pair)
					if err != nil {
						log.Println(err)
						continue
					}

					currentClosedPrice, err := strconv.ParseFloat(resp.FLOWEUR.Close[0], 64)
					if err != nil {
						return err
					}

					buyAction := buy(target, currentClosedPrice)
					if buyAction {
						if !test {
							order, err := krakenClient.AddOrder(pair, "buy", "market", c.Args().Get(1), nil)
							if err != nil {
								return err
							}
							log.Printf("BUY %v\n", order.Description.PrimaryPrice)
							return nil
						} else {
							log.Printf("BUY simulation %v\n", currentClosedPrice)
						}
						return nil
					}
				}
				return nil
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

type Pair struct {
	Origin string
	Target string
}

func splitCryptoNameInPair(pair string) (Pair, error) {
	tests := []string{"EUR", "USD", "XBT", "GBP", "JPY", "ETH"}
	for _, test := range tests {
		split := strings.Split(pair, test)
		if len(split) > 1 {
			return Pair{
				Origin: split[0],
				Target: split[1],
			}, nil
		}
	}

	return Pair{}, errors.New("failed to split pair name")
}
