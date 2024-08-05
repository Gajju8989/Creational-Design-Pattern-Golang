package factory

import "design_patterns/abstractfactory/transport/product"

// AirTransportFactory is a concrete factory that creates air transports.
type AirTransportFactory struct{}

func (f *AirTransportFactory) CreateTransport(transportType string) product.Transport {
	switch transportType {
	case "airplane":
		return &product.Airplane{}
	default:
		return nil
	}
}
