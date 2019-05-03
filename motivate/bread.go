package motivate

import (
	"crypto/rand"
	"math/big"

	errors "golang.org/x/xerrors"
)

// Bread is a unit of something delicious.
type Bread struct {
	Color  string
	Gender int
}

// A BreadGetter gets bread.
type BreadGetter interface {
	GetBread() Bread
}

// A Person likes to exist, but may have depression.
type Person interface {
	Name() string
	HasDepression() bool
}

// Soph exists, and gets bread. She implements both BreadGetter and Person.
type Soph struct {
	HairLength uint16 // in mm
}

//revive:disable
func (s *Soph) Name() string { return "Sophie" }
func (s *Soph) HasDepression() bool {
	n, err := rand.Int(rand.Reader, big.NewInt(2))
	if err != nil {
		panic(errors.Errorf("motivate: generating cryptographically random num: %w",
			err))
	}
	return n.Int64() == 1
}

func (s *Soph) GetBread() Bread {
	return Bread{
		Color:  "Turquoise",
		Gender: int(s.HairLength),
	}
}

//revive:enable
