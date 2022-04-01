## Uptycs Client (Go)

A Go library for [uptycs](https://uptycs.io)


```
package main

import (
  "uptycs-client-go"
  "fmt"
  "os"
)

func main() {
  c, err := uptycs.NewClient(uptycs.UptycsConfig{
    Host: os.Getenv("UPTYCS_HOST"),
    ApiKey: os.Getenv("UPTYCS_API_KEY"),
    ApiSecret: os.Getenv("UPTYCS_API_SECRET"),
    CustomerID: os.Getenv("UPTYCS_CUSTOMER_ID"),
  })
  if err != nil {
    fmt.Println(err)
    return
  }
  rules, err := c.GetAlertRules()
  fmt.Println(rules.Limit, err)
  for _, alert := range rules.Items {
    fmt.Println(alert.Rule)
  }
}
```
