package builder

// Product
type House struct {
	Windows int
	Doors   int
	Rooms   int
	Garage  bool
	Pool    bool
}

// Builder interface
type HouseBuilder interface {
	SetWindows(int) HouseBuilder
	SetDoors(int) HouseBuilder
	SetRooms(int) HouseBuilder
	SetGarage(bool) HouseBuilder
	SetPool(bool) HouseBuilder
	Build() House
}

// Concrete Builder
type ConcreteHouseBuilder struct {
	house House
}

func NewConcreteHouseBuilder() *ConcreteHouseBuilder {
	return &ConcreteHouseBuilder{house: House{}}
}

func (b *ConcreteHouseBuilder) SetWindows(windows int) HouseBuilder {
	b.house.Windows = windows
	return b
}

func (b *ConcreteHouseBuilder) SetDoors(doors int) HouseBuilder {
	b.house.Doors = doors
	return b
}

func (b *ConcreteHouseBuilder) SetRooms(rooms int) HouseBuilder {
	b.house.Rooms = rooms
	return b
}

func (b *ConcreteHouseBuilder) SetGarage(hasGarage bool) HouseBuilder {
	b.house.Garage = hasGarage
	return b
}

func (b *ConcreteHouseBuilder) SetPool(hasPool bool) HouseBuilder {
	b.house.Pool = hasPool
	return b
}

func (b *ConcreteHouseBuilder) Build() House {
	return b.house
}

// Director
type HouseDirector struct {
	builder HouseBuilder
}

func NewHouseDirector(b HouseBuilder) *HouseDirector {
	return &HouseDirector{builder: b}
}

func (d *HouseDirector) ConstructBasicHouse() House {
	return d.builder.SetWindows(4).SetDoors(2).SetRooms(3).Build()
}

func (d *HouseDirector) ConstructLuxuryHouse() House {
	return d.builder.SetWindows(8).SetDoors(4).SetRooms(5).SetGarage(true).SetPool(true).Build()
}
