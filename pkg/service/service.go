package service

import (
	"github.com/mrNobody95/Gate/pkg/dispose"
	log "github.com/sirupsen/logrus"
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

func Start(name, version string) {
	var exposes []uint16
	interrupt := make(chan error)
	for port, serve := range serves {
		exposes = append(exposes, port)
		go func(port uint16, serve ServeFunc) {
			interrupt <- serve.Listen(port)
		}(port, serve)
	}

	data := log.Fields{
		"service": name,
		"version": version,
		"exposes": exposes,
	}
	log.WithFields(data).Info("service is running")

	interruptErr := <-interrupt
	if err := dispose.Close(); err != nil {
		log.WithError(err).Error("can not dispose")
	}
	if interruptErr == nil {
		log.WithFields(data).Panic("service interrupted")
	} else {
		log.WithFields(data).WithError(interruptErr).Fatal("service interrupted")
	}
}
