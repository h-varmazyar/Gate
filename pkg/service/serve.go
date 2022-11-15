package service

import (
	"fmt"
	"net"
)

/**
* Dear programmer:
* When I wrote this code, only god And I know how it worked.
* Now, only god knows it!
*
* Therefore, if you are trying to optimize this code And it fails(most surely),
* please increase this counter as a warning for the next person:
*
* total_hours_wasted_here = 0 !!!
*
* Best regards, mr-nobody
* Date: 12.11.21
* Github: https://github.com/mrNobody95
* Email: hossein.varmazyar@yahoo.com
**/

var (
	serves = make(map[uint16]ServeFunc)
)

type ServeFunc func(listener net.Listener) error

func Serve(port uint16, f ServeFunc) {
	serves[port] = f
}

func (serve ServeFunc) Listen(port uint16) error {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%v", port))
	if err != nil {
		return err
	}
	defer func() {
		_ = listener.Close()
	}()
	return serve(listener)
}
