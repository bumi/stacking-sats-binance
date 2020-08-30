# Stacking Sats on Binance(.je)

This tool automates stacking sats on Binance for you. 
Run it periodically to automatically place BTC buy orders on binance and 
convert a certain percentage of the available EUR balance to BTC.

Use this at your own risk and decide for yourself whether or not you want to run this tool!

## 🔑 Binance API Key

Obtain your Binance API Key via the [API settings page](https://www.binance.je/userCenter/createApi.html).
The key must have the following options enabled: "Read Info" and "Enable Trading"

## 💰 Run it

```sh
./stacking-sats
```

### Help
```
NAME:
   Stacking Sats on Binance - Automate market orders based on the available EUR balance

USAGE:
   stacking-sats [global options] command [command options] [arguments...]

COMMANDS:
   stack     Stacks some sats - places a new market BTC buy order
   withdraw  Withdraws sats to your wallet
   list      Lists recent orders
   help, h   Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --apikey value   Binance API Key [$BINANCE_APIKEY]
   --secret value   Binance secret [$BINANCE_SECRET]
   --baseurl value  Binance API URL (default: "https://api.binance.je")
   --help, -h       show help (default: false)
```

### 🤑 Stacking Sats

```sh
./stacking-sats --apikey=YOURAPIKEY --secret=YOURSECRET stack
```

```
OPTIONS:
   --interval value    Days since the last order. (set to 0 to ignore) (default: 7 days)
   --percentage value  Percentage of the available EUR balance (default: 25)
   --maxprice value    Max price in EUR (default: 15000.0)
   --help, -h          show help (default: false)
```

### List Orders

```sh
./stacking-sats --apikey=YOURAPIKEY --secret=YOURSECRET list
```

```
OPTIONS:
   --limit value  Lists recent orders (default: 10)
   --help, -h     show help (default: false)
```

### TODO: Withdraw BTC

```sh
./stacking-sats --apikey=YOURAPIKEY --secret=YOURSECRET withdraw
```

## ⛑ Guards

Some guards try to prevent potential errors:

### Order Volumne

The calculated order quantity must be between 0.001 BTC and 0.05 BTC. This is a fixed limit.

### Interval
Set an interval to make sure the last order is at least X days ago. 

Default: 7

Example: `./stacking-sats stack --maxprice=14`

### Max Price
You can define a max price in EUR. If BTC is above that price no order will be executed. 

Default: 15000.0

Example: `./stacking-sats stack --maxprice=20000`


## 🗓 Cron Job
Use cron to run this script periodically


## Similar Tools

* [stacking-sats-kraken](https://github.com/dennisreimann/stacking-sats-kraken) by [@dennisreimann](https://twitter.com/dennisreimann)

