package dispose

import (
	"fmt"
	"strings"
	"sync"
)

type closeFunc func() error

var (
	disposes []closeFunc
	locker   sync.Mutex
)

func Add(c closeFunc) {
	locker.Lock()
	defer locker.Unlock()
	disposes = append(disposes, c)
}

func Close() error {
	var msg []string
	locker.Lock()
	defer locker.Unlock()
	for _, c := range disposes {
		if err := c(); err != nil {
			msg = append(msg, err.Error())
		}
	}
	if len(msg) > 0 {
		return fmt.Errorf(strings.Join(msg, "\n"))
	}
	return nil
}
