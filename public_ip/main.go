package main

import (
  "fmt"
  "os"
  "github.com/vvlad/public_ip"
)


func main() {

  address, err := public_ip.IpAddress()

  if err != nil {
    fmt.Printf("%v\n", err);
    os.Exit(1)
  }

  fmt.Printf("%v\n", address);

}
