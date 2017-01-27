package public_ip

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
	"regexp"
)

var defaultServices  = []string{
	"http://echoip.com",
	"http://ifconfig.co/x-real-ip",
	"http://icanhazip.com/",
	"http://ifconfig.io/ip",
	"http://ip.appspot.com/",
	"http://curlmyip.com/",
	"http://ident.me/",
	"http://whatismyip.akamai.com/",
	"http://tnx.nl/ip",
	"http://myip.dnsomatic.com/",
	"http://ipecho.net/plain",
	"http://diagnostic.opendns.com/myip",
}

type IpResult struct {
	Success bool
	Ip string
	Error error
}

func GetIP(services []string) *IpResult {
	if services == nil || len(services) == 0 {
		services = defaultServices
	}
	done := make(chan *IpResult)
	for k := range services {
		go ipAddress(services[k], done)
	}
	for ;; {
		select{
		case result := <-done:
			if result.Success {
				return result
			}
			continue
		case <-time.After(time.Second * 30):
			return &IpResult{false, "", errors.New("Timed out")}
		}
	}
}

func ipAddress(service string, done chan<- *IpResult) {

	resp, err := http.Get(service)
	defer resp.Body.Close()

	if err == nil {

		address, err := ioutil.ReadAll(resp.Body)
		ip := fmt.Sprintf("%s", strings.TrimSpace(string(address)))
		if err== nil && checkIp(ip) {
			sendResult(&IpResult{true, ip, nil}, done)
			return
		}
	}
	sendResult(&IpResult{false, "", errors.New("Unable to talk with a service")}, done)
}

func sendResult(result *IpResult, done chan<- *IpResult) {
	select {
	case done <- result:
		return
	default:
		return
	}
}

func checkIp(ip string) bool {
	match, _ := regexp.MatchString(`^(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])$`, ip)
	if match {
		return true
	}
	return false
}
