package utils

import (
	"net"
	"os"
	"reflect"
	"sync"
)

const (
	Empty = ""
)

var (
	internalIpOnce sync.Once
)

//ips
func GetInternal() string {
	internalIp := ""

	internalIpOnce.Do(func() {
		addrs, err := net.InterfaceAddrs()
		if err != nil {
			os.Stderr.WriteString("Oops:" + err.Error())
			os.Exit(1)
		}
		for _, a := range addrs {
			if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
				if ipnet.IP.To4() != nil {
					internalIp = ipnet.IP.To4().String()
				}
			}
		}
	})
	return internalIp
}

func IsNotNil(object interface{}) bool {
	return !IsNilObject(object)
}

func IsNilObject(object interface{}) bool {
	if object == nil {
		return true
	}

	value := reflect.ValueOf(object)
	kind := value.Kind()
	if kind >= reflect.Chan && kind <= reflect.Slice && value.IsNil() {
		return true
	}

	return false
}
