package api

type Cost struct {
	Enabled bool
	cost    int64
}

func NewCost() *Cost {
	return &Cost{}
}

func NewCostWith(cost int64) *Cost {
	return &Cost{cost: cost, Enabled: (cost > 0)}
}

func (c *Cost) UpdateCost(cost int64) int64 {
	c.cost += cost
	return c.cost
}

func (c *Cost) GetCost() int64 {
	return c.cost
}
