package position

const DefaultBlockSize = 50

type Pool struct {
	block []Position
	off   int
}

func NewPool(blockSize int) *Pool {
	return &Pool{
		block: make([]Position, blockSize),
	}
}

func (p *Pool) Get() *Position {
	if len(p.block) == 0 {
		return nil
	}

	if len(p.block) == p.off {
		p.block = make([]Position, len(p.block))
		p.off = 0
	}

	p.off++

	return &p.block[p.off-1]
}
