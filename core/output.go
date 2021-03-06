package core

/*
#cgo CFLAGS: -I./c/include
#include "lwip/tcp.h"

#define AF_INET         2
#define AF_INET6        30

extern err_t output(struct pbuf *p,int t);

err_t
output_ip4(struct netif *netif, struct pbuf *p, const ip4_addr_t *ipaddr)
{
	return output(p,AF_INET);
}

err_t
output_ip6(struct netif *netif, struct pbuf *p, const ip6_addr_t *ipaddr)
{
	return output(p,AF_INET6);
}

void
set_output()
{
	if (netif_list != NULL) {
		(*netif_list).output = output_ip4;
		(*netif_list).output_ip6 = output_ip6;
	}
}
*/
import "C"
import (
	"errors"
)

var OutputFn func([]byte,int) (int, error)

func RegisterOutputFn(fn func([]byte,int) (int, error)) {
	OutputFn = fn
	C.set_output()
}

func init() {
	OutputFn = func(data []byte,tp int) (int, error) {
		return 0, errors.New("output function not set")
	}
}
