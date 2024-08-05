package factory

import "design_patterns/abstractfactory/transport/product"

// LandTransportFactory is a concrete factory that creates land transports.
type LandTransportFactory struct{}

func (f *LandTransportFactory) CreateTransport(transportType string) product.Transport {
	switch transportType {
	case "car":
		return &product.Car{}
	default:
		return nil
	}
}
