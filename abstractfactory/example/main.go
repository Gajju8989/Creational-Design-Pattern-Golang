package main

import (
	"design_patterns/abstractfactory/transport/client"
	"design_patterns/abstractfactory/transport/factory"
	"fmt"
)

func main() {
	var landFactory factory.TransportFactory = &factory.LandTransportFactory{}
	var airFactory factory.TransportFactory = &factory.AirTransportFactory{}

	landTransport := client.GetTransport(landFactory, "car")
	airTransport := client.GetTransport(airFactory, "airplane")

	fmt.Println("Land Transport: ", landTransport.GetTransportType())
	fmt.Println("Air Transport: ", airTransport.GetTransportType())
}
