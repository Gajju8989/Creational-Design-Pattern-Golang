package product

// Car is a concrete product that implements the Transport interface.
type Car struct{}

func (c *Car) GetTransportType() string {
	return "Car"
}
