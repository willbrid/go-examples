package thing

type Thing struct {
	X, Y, Z int
}

func NewThing(x, y, z int) (*Thing, error) {
	return &Thing{X: x, Y: y, Z: z}, nil
}
