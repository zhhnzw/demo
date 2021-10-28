package test

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

func TestCase001TestSuite(t *testing.T) {
	suite.Run(t, new(Case001TestSuite))
}
