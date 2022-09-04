package RealExample

type House struct {
	Material     string
	HasFireplace bool
	Floors       int
}

type HouseOption func(*House)

func WithConcrete() HouseOption {
	return func(h *House) {
		h.Material = "concrete"
	}
}

func WithFloors(floors int) HouseOption {
	return func(h *House) {
		h.Floors = floors
	}
}

func NewHouse(options ...HouseOption) *House {
	const (
		defaultMaterial  = "Wood"
		defaultFloors    = 4
		defaultFireplace = true
	)
	h := &House{
		Material:     defaultMaterial,
		Floors:       defaultFloors,
		HasFireplace: defaultFireplace,
	}

	// Adding HouseOption
	for _, hOption := range options {
		hOption(h)
	}
	return h
}
