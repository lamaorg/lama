package tests

import (
	"lamanodrama/lama/internals/dnaProof"

	"github.com/stretchr/testify/suite"
	"time"
)

type LamaDnaProofTestSuite struct {
	suite.Suite
	params dnaProof.TemperProofParams
}

func (suite *LamaDnaProofTestSuite) SetupTest() {
	suite.params.MaxFitness = 0.5
	suite.params.ParseDuration, _ = time.ParseDuration("50000ms")
	suite.params.PopulationSize = 1000
	suite.params.MutationRate = 0.005
}
