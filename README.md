# WIP - kraken-fishing
## Optimize your Buy or sell order on kraken plateform 

```
NAME:
   Catch The Kraken !

VERSION:
   v0.1

DESCRIPTION:
   Optimize your Buy or sell order

COMMANDS:
   sell     sell with a minimal target - sell <pair> <volume> <target>
                pair: asset pair to sell
                volume: volume of order (you can set 'all')
                target: price to sell
                OPTIONS:
                  --test value  Set yes to achieve real orders (default: "yes")
                  --tick value  Set 30/15m/1h to set the price check interval (default: 15m)
                  --hold value  Set the count of negative evolution before sell/buy (default: "2")
                  --help, -h    show help (default: false)

   buy      buy with a maximal target - buy <pair> <volume> <target>
                pair: asset pair to buy
                volume: volume of order
                target: price to buy
                OPTIONS:
                  --test value  Set yes to achieve real orders (default: "yes")
                  --tick value  Set 30/15m/1h to set the price check interval (default: 5)
                  --hold value  Set --hold=<count> to set the count of negative value before sell/buy (default: "2")
                  --help, -h    show help (default: false)
   GLOBAL OPTIONS:
      --config value  Set path of config file (default: "config.yml")
      --help, -h      show help (default: false)
      --version, -v   print the version (default: false)
   ```
