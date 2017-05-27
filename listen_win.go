// +build windows

package http

import (
	"errors"
	"net"
	"strings"
)

func (srv *App) listen() (l net.Listener, err error) {
	switch addr := srv.address; {
	case strings.HasPrefix(addr, "http://"):
		addr = strings.TrimPrefix(addr, "http://")
		fallthrough
	default:
		l, err = net.Listen("tcp", addr)
	}
	return
}
