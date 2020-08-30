# ⚡ Stacking Sats on Binance(.je) ⚡

This tool automates stacking sats on [Binance](https://www.binance.je/en) for you. 
Run it periodically to automatically place BTC buy orders and 
convert a certain percentage of the available EUR balance to BTC. (Dollar-cost averaging your Bitcoin savings)

## What is stacking sats? What is Dollar-cost averaging?

"The act of accumulating satoshis ("sats"), the penny of bitcoin over time. This term emphasizes that even really small accumulations of bitcoin are useful because of the value they will have in the future." ([urban dictionary](https://www.urbandictionary.com/define.php?term=stacking%20sats))

It is a another way of saying "saving and accumulating Bitcoin".


Dollar-cost averaging is an investment strategy that aims to reduce the impact of volatility on the purchase of assets. It involves buying equal amounts of the asset at regular intervals.

Have a look at the DCA Investment calculator for Bitcoin: [dcabtc.com](https://dcabtc.com/)



## 💾 1. Download it

Simply download the [latest release](https://github.com/bumi/stacking-sats-binance/releases). No additional dependencies required.


## 🔑 2. Get your Binance API Key
To connect your Binance account you need to configure your API Key:

You can obtain the API Key via the [API settings page](https://www.binance.je/userCenter/createApi.html).
The key must have the following options enabled: "Read Info" and "Enable Trading"


## 💰 3. Run it

```sh
./stacking-sats
```

### Help
```
./stacking-sats --help

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

### 🌅 Stacking Sats

```sh
./stacking-sats --apikey=YOURAPIKEY --secret=YOURSECRET stack
```

#### Configuration
Provide the API Key and Secret as global parameter or as environment variables. 
All other options can be set as parameter.

##### Percentage
The percentage of the available EUR balance you want to convert to BTC. Default value: 25%

##### Interval
Guard to prevent too many accidental orders. Checks that the last order is at least X days ago. Default value: 7days

#### Example

```sh
./stacking-sats --apikey=YOURAPIKEY --secret=YOURSECRET stack --percentage=50 --interval=14
```

#### Command Help
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
Use [cron](https://en.wikipedia.org/wiki/Cron) to run stacking-sats periodically:

```sh
crontab -e
```

### Example:

Every Sunday 8:00pm 
```
0 20 * * 0 /home/bitcoin/stacking-sats --apikey=YOURAPIKEY --secret=YOURSECRET stack >> /home/bitcoin/stacking-sats.log 2>&1
````

Note: adjust the path to your `stacking-sats` file and check the logs.


## Similar Tools

* [stacking-sats-kraken](https://github.com/dennisreimann/stacking-sats-kraken) by [@dennisreimann](https://twitter.com/dennisreimann)

## Disclaimer

Use this at your own risk and decide for yourself whether or not you want to run this tool!
Audit the code and check the dependencies ([adshao/go-binance](https://github.com/adshao/go-binance), [urfave/cli/](https://github.com/urfave/cli/)) yourself.

If you have questions or problems, let me know or open an issue.

