package position

// DefaultBlockSize is the first slab. Full slabs are replaced by a slab of
// twice the size; pointers into earlier slabs stay valid.
const DefaultBlockSize = 32

type Pool struct {
	block []Position
	off   int
}

func NewPool(blockSize int) *Pool {
	if blockSize < 1 {
		blockSize = DefaultBlockSize
	}
	return &Pool{
		block: make([]Position, blockSize),
	}
}

func (p *Pool) Get() *Position {
	if p.off == len(p.block) {
		next := len(p.block) * 2
		if next < DefaultBlockSize {
			next = DefaultBlockSize
		}
		p.block = make([]Position, next)
		p.off = 0
	}

	p.off++

	return &p.block[p.off-1]
}
