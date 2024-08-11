package unittest

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type AddSuite struct {
	suite.Suite
}

func (s *AddSuite) TestAdd() {
	// result := add(1, 2, 3, 4, 5)
	// s.Equal(15, result)

	s.Run("add 2 positive number", func() {
		result := add(1, 2)
		s.Equal(3, result)
	})

	s.Run("add 2 negative number", func() {
		result := add(-1, -2)
		s.Equal(-32, result)
	})
}

func TestAddSuite(t *testing.T) {
	suite.Run(t, new(AddSuite))
}
