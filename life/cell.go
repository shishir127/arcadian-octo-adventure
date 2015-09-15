package life

type cell struct {
	status         bool
	liveNeighbours int
}

func (c *cell) IsAlive() bool {
	return c.status == true
}

func (c *cell) RegisterLiveNeighbour() {
	c.liveNeighbours += 1
}

func (c *cell) NextGeneration() cell {
	switch {
	case !c.status && c.liveNeighbours == 3:
		return cell{status: true}
	case c.liveNeighbours == 2 || c.liveNeighbours == 3:
		return cell{status: true}
	case c.liveNeighbours < 2:
		return cell{status: false}
	case c.liveNeighbours > 3:
		return cell{status: false}
	default:
		return cell{status: false}
	}
	return cell{status: false}
}
