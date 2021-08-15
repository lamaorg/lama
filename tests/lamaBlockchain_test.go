package tests

import (
	"github.com/lamaorg/lama/internals/primitives"
	"log"
	"testing"
)

func TestGenerateChainID(t *testing.T) {
	c := primitives.GenerateChainID()
	log.Printf("chain id: %v", c)
}
