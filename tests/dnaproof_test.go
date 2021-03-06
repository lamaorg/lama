package tests

import (
	"github.com/lamaorg/lama/internals/dnaProof"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"

	"testing"
	"time"
)

func (suite *LamaDnaProofTestSuite) TestGetProof() {

	currentTime := time.Now()
	assert.NotNil(suite.T(), interface{}(dnaProof.GetProof(suite.params, currentTime)))

}

func TestLamaDnaProofTestSuite(t *testing.T) {
	suite.Run(t, new(LamaDnaProofTestSuite))
}

/*
func TestGetMediaProof(t *testing.T) {
	runtime.GOMAXPROCS(6)
	runtime.MemProfileRate = 16000 * 1024
	testImagePath := "./mocks/cat.png"
	dnaProof.GetMediaProof(testImagePath)
}*/
