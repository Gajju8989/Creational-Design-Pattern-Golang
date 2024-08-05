package factory

import "design_patterns/abstractfactory/transport/product"

// TransportFactory defines the abstract factory interface.
type TransportFactory interface {
	CreateTransport(transportType string) product.Transport
}
