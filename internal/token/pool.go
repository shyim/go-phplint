package token

// DefaultBlockSize is the first slab. Later slabs double so small files do
// not reserve a thousand tokens and large files do not allocate fixed slabs.
const DefaultBlockSize = 32

type Pool struct {
	block []Token
	off   int
}

func NewPool(blockSize int) *Pool {
	if blockSize < 1 {
		blockSize = DefaultBlockSize
	}
	return &Pool{
		block: make([]Token, blockSize),
	}
}

func (p *Pool) Get() *Token {
	if p.off == len(p.block) {
		next := len(p.block) * 2
		if next < DefaultBlockSize {
			next = DefaultBlockSize
		}
		p.block = make([]Token, next)
		p.off = 0
	}

	p.off++

	return &p.block[p.off-1]
}
