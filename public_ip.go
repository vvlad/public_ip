package public_ip

import (
  "net/http"
  "math/rand"
  "io/ioutil"
  "fmt"
  "strings"
  "errors"
)

func IpAddress() (string, error){

  services := []string{
    "http://echoip.com",
    "http://icanhazip.com",
  }

  for len(services) > 0 {

    index := rand.Intn(len(services))

    service := services[index]
    services = append(services[:index], services[index+1:]...)
    resp, err := http.Get(service)

    if err != nil { continue }

    defer resp.Body.Close()

    address, err := ioutil.ReadAll(resp.Body)

    if err != nil { continue }

    return fmt.Sprintf("%s", strings.TrimSpace(string(address))), nil

  }

  return "", errors.New("Unable to talk with a service")
}
