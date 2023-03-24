package api

type Cost struct {
	cost int64
}

func NewCost() *Cost {
	return &Cost{cost: -1}
}

func (c *Cost) SetCost(cost int64) {
	c.cost = cost
}

func (c *Cost) GetCost() int64 {
	return c.cost
}
