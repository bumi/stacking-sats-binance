package main

import (
  "context"
  "fmt"
  "os"
  "log"
  "time"
  "net/http"
  "strconv"
  "math"
  "errors"
  "github.com/urfave/cli/v2"
  "github.com/adshao/go-binance"
)

func newBinanceClient(apiKey, secretKey, baseUrl string) binance.Client {
  return binance.Client{
    APIKey:     apiKey,
    SecretKey:  secretKey,
    BaseURL:    baseUrl,
    UserAgent:  "Binance/golang/stacking-sats",
    HTTPClient: http.DefaultClient,
    Logger:     log.New(os.Stderr, "stacking sats ", log.LstdFlags),
  }
}

func withdraw(c *cli.Context) error {
  return errors.New("Not yet implemented")
}

func list(c *cli.Context) error {
  fmt.Println(fmt.Sprintf("📄 Recent orders 📅 %s\n", time.Now()))
  client := newBinanceClient(c.String("apikey"), c.String("secret"), c.String("baseurl"));

  orders, err := client.NewListOrdersService().Symbol("BTCEUR").Limit(c.Int("limit")).
    Do(context.Background())
  if err != nil {
    return err
  }
  for _, o := range orders {
    fmt.Println(fmt.Sprintf("⚡ OrderID=%d Status=%s Price=%s ExecutedQuantity=%s CummulativeQuoteQuantity=%s Type=%s Symbol=%s", o.OrderID, o.Status, o.Price, o.ExecutedQuantity, o.CummulativeQuoteQuantity, o.Type, o.Symbol))
  }
  return nil
}

func stack(c *cli.Context) error {
  fmt.Println(fmt.Sprintf("\n🎉 Stacking sats! 📅 %s", time.Now()))

  client := newBinanceClient(c.String("apikey"), c.String("secret"), c.String("baseurl"));

  lastOrders, err := client.NewListOrdersService().Symbol("BTCEUR").Limit(1).
    Do(context.Background())
  if err != nil {
    return err
  }
  if len(lastOrders) != 0 {
    lastOrderAt := lastOrders[0].Time / 1000 // timstamp in milliseconds
    fmt.Println(fmt.Sprintf("⏱  Last order: %s", time.Unix(lastOrderAt, 0)))
    if lastOrderAt > time.Now().Unix() - 60*60*24 * c.Int64("interval") {
      return errors.New(fmt.Sprintf("🚨 Last order is less than %d days ago", c.Int64("interval")))
    }
  }

  account, err := client.NewGetAccountService().Do(context.Background())
  if err != nil {
    return err
  }

  var amount float64
  for i := range account.Balances {
    if account.Balances[i].Asset == "EUR" {
      freeAmount, _ := strconv.ParseFloat(account.Balances[i].Free, 10)
      amount = freeAmount * (c.Float64("percentage")/100.0)
      break
    }
  }
  fmt.Println(fmt.Sprintf("💰 Current EUR account balance: %f EUR", amount))

  prices, err := client.NewListBookTickersService().Symbol("BTCEUR").
        Do(context.Background())
  if err != nil {
    return err
  }
  price, _ := strconv.ParseFloat(prices[0].BidPrice, 10)
  fmt.Println(fmt.Sprintf("📈 Current BTC price: %fEUR", price))
  if price > c.Float64("maxprice") {
    return errors.New(fmt.Sprintf("🚨 Price > %fEUR", c.Float64("maxprice")))
  }

  orderQuantity := math.Round((amount / price)*100000)/100000
  if orderQuantity > 0.05 || orderQuantity < 0.001 {
    return errors.New(fmt.Sprintf("🚨 Invalid orderQuantity: %f", orderQuantity))
  }

  fmt.Println(fmt.Sprintf("💸 Ordering: %fBTC at market price (~%fEUR)", orderQuantity, amount))

  order, err := client.NewCreateOrderService().Symbol("BTCEUR").
        Side(binance.SideTypeBuy).Type(binance.OrderTypeMarket).
        Quantity(fmt.Sprintf("%f", orderQuantity)).
        Do(context.Background())
  if err != nil {
    fmt.Println("🚨 Order failed")
    return err
  }
  fmt.Println(fmt.Sprintf("⚡ Order placed: OrderId=%d Status=%s", order.OrderID, order.Status))
  return nil
}

func main() {

  app := cli.NewApp()
  app.Name = "Stacking Sats on Binance"
  app.Usage = "Automate market orders based on the available EUR balance"
  app.Flags = []cli.Flag{
    &cli.StringFlag{
      Name: "apikey",
      Usage: "Binance API Key",
      EnvVars: []string{"BINANCE_APIKEY"},
      Required: true,
    },
    &cli.StringFlag{
      Name: "secret",
      Usage: "Binance secret",
      EnvVars: []string{"BINANCE_SECRET"},
      Required: true,
    },
    &cli.StringFlag{
      Name: "baseurl",
      Usage: "Binance API URL",
      Value: "https://api.binance.je",
    },
  }
  app.Commands = []*cli.Command{
    {
      Name:  "stack",
      Usage: "Stacks some sats - places a new market BTC buy order",
      Action: stack,
      Flags: []cli.Flag{
        &cli.Int64Flag{
          Name: "interval",
          Usage: "days since the last order (set to 0 to ignore)",
          Value: 7,
          DefaultText: "7 days",
        },
        &cli.Float64Flag{
          Name: "percentage,p",
          Usage: "percentage of the available EUR balance to buy",
          Value: 25,
          DefaultText: "25",
        },
        &cli.Float64Flag{
          Name: "maxprice",
          Usage: "max BTC price in EUR",
          Value: 15000.0,
          DefaultText: "15000.0",
        },
      },
    },
    {
      Name: "withdraw",
      Usage: "Withdraws sats to your wallet",
      Action: withdraw,
    },
    {
      Name: "list",
      Usage: "Lists recent orders",
      Action: list,
      Flags: []cli.Flag{
        &cli.Int64Flag{
          Name: "limit",
          Usage: "Number of orders to list",
          Value: 10,
          DefaultText: "10",
        },
      },
    },
  }

  err := app.Run(os.Args)
  if err != nil {
    log.Fatal(err)
  }
}
