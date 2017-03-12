package public_ip

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"faunus_backend/shared/logger"
	"os"
	"time"
)

func Test_GetIpWIthEmptyIpSlice(t *testing.T) {
	assert := assert.New(t)
	ipResult := GetIP([]string{}, time.Duration(0))
	assert.Nil(ipResult.Error)
	logger.L.Debug(ipResult.Ip)
}

func Test_GetIpWithInvalidServicesSpec(t *testing.T) {
	assert := assert.New(t)
	ipResult := GetIP([]string{"google.com"},time.Duration(0))
	assert.Equal(false, ipResult.Success)
	assert.NotNil(ipResult.Error)
}

func Test_GetIpWithInvalidService(t *testing.T) {
	assert := assert.New(t)
	ipResult := GetIP([]string{"http://google.com"},time.Duration(0))
	assert.Equal(false, ipResult.Success)
	assert.NotNil(ipResult.Error)
}

func Test_GetIpWithEmptyServiceSpecs(t *testing.T) {
	assert := assert.New(t)
	ipResult := GetIP([]string{"", "", ""}, time.Duration(0))
	assert.Equal(false, ipResult.Success)
	assert.NotNil(ipResult.Error)
	logger.L.Debug(ipResult.Error.Error())
}

func Test_GeIpReturnAValidIp(t *testing.T) {
	ip := os.Getenv("IP")
	assert := assert.New(t)
	ipResult := GetIP([]string{}, time.Duration(0))
	assert.Nil(ipResult.Error)
	assert.Equal(ip, ipResult.Ip)
	logger.L.Debug(ipResult.Ip)
}