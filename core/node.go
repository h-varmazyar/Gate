package core

import "github.com/mrNobody95/Gate/strategies"

type Node struct {
	Strategy       strategies.Strategy
	Scheduler      interface{}
	Algorithm      interface{}
	NetworkManager interface{}
}

func (n *Node) Start() error {
	if err := n.Strategy.Validate(); err != nil {
		return err
	}
}
