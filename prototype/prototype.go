package prototype

import (
	"fmt"
)

type Prototype interface {
	Clone() Prototype
}

type ConcretePrototype struct {
	Value int
}

func (c *ConcretePrototype) Clone() Prototype {
	return &ConcretePrototype{Value: c.Value}
}

func (c *ConcretePrototype) String() string {
	return fmt.Sprintf("ConcretePrototype with value: %d", c.Value)
}
