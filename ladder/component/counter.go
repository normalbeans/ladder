package component

type Counter struct {
	Originx, Originy, Width, Height int
}

func (c *Counter) GetOrigin() (int, int) {
	return c.Originx, c.Originy
}