package coin

import "math/rand"

func New() *Coin {
	coin := new(Coin)
	coin.rand = rand.New(rand.NewSource(99))
	return coin
}

type Face int

const (
	Heads = iota
	Tails
)

func (f Face) String() string {
	if f == Heads {
		return "Heads"
	} else {
		return "Tails"
	}
}

type Coin struct {
	rand *rand.Rand
}

func (coin Coin) Flip() Face {
	i := coin.rand.Intn(2)
	if i == Heads {
		return Heads
	} else {
		return Tails
	}
}