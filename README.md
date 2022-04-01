## Uptycs Client (Go)

A Go library for [uptycs](https://uptycs.io)


```
package main

import (
  "uptycs-client-go"
  "fmt"
)

func main() {
  c, err := uptycs.NewClient()
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
