package client

import (
	"design_patterns/abstractfactory/transport/factory"
	"design_patterns/abstractfactory/transport/product"
)

// GetTransport gets a transport based on the factory and type.
func GetTransport(factory factory.TransportFactory, transportType string) product.Transport {
	return factory.CreateTransport(transportType)
}
