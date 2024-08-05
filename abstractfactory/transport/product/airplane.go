package product

// Airplane is a concrete product that implements the Transport interface.
type Airplane struct{}

func (a *Airplane) GetTransportType() string {
	return "Airplane"
}
