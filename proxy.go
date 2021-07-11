package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"time"

	"golang.org/x/net/proxy"
)

// canConnectThroughProxy uses an internet proxy to see if my ports are open
// network should be either tcp4 or tcp6, not tcp
func canConnectThroughProxy(proxyaddr string, tcpaddr *net.TCPAddr, network string) (proxyOK bool, portOpen bool) {
	const (
		baseDialerTimeout  = 3 * time.Second
		proxyDialerTimeout = 3 * time.Second
		SOCKS5PROXY        = "deadsfu.com:60000"
	)
	if network != "tcp4" && network != "tcp6" {
		checkFatal(fmt.Errorf("network not okay"))
	}

	if proxyaddr == "" {
		proxyaddr = SOCKS5PROXY
	}

	baseDialer := &net.Dialer{
		Timeout: baseDialerTimeout,
		//Deadline: time.Time{},
		//FallbackDelay: -1,
	}

	// always get to proxy using ipv4, more reliable for this test
	dialer, err := proxy.SOCKS5(network, proxyaddr, nil, baseDialer)
	if err != nil {
		return
	}

	contextDialer, ok := dialer.(proxy.ContextDialer)
	if !ok {
		log.Println("cannot deref dialer")
		//not fatal
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), proxyDialerTimeout)
	_ = cancel

	// we ask the proxy for the given network
	conn, err := contextDialer.DialContext(ctx, network, tcpaddr.String())
	if err != nil {

		println(888, err.Error())
		println(999, errors.Is(err, os.ErrDeadlineExceeded))
		println(111)
		for xx := err; xx != nil; xx = errors.Unwrap(xx) {
			fmt.Printf("%#v\n", xx)

			var operr *net.OpError
			if errors.As(xx, &operr) {
				println(222, operr.Op == "read")
				println(333, errors.Is(xx, os.ErrDeadlineExceeded))
				println(444, operr.Err == os.ErrDeadlineExceeded)

				fmt.Println("Failed at op:", operr.Op)
			} else {
				fmt.Println(err)
			}
		}
		println(111)
		for xx := err; xx != nil; xx = errors.Unwrap(xx) {
			fmt.Printf("%+v\n", xx)
		}
		println(111)

		readtcp := strings.Contains(err.Error(), " read tcp ")
		dialtcp := strings.Contains(err.Error(), " dial tcp ")
		iotimeout := strings.Contains(err.Error(), "i/o timeout")

		if iotimeout && readtcp {
			proxyOK = true
			portOpen = false
		} else if iotimeout && dialtcp {
			proxyOK = false
			portOpen = false
		}
		// sometime else is going on, but we are going to stay silent
		// the user isn't going to know what to do

		return
	}
	conn.Close()

	// if we got here, I got through the proxy and connected to myself

	proxyOK = true
	portOpen = true
	return
}
