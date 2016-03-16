package edgegrid

import (
	"testing"
)

func TestComputeHmac256(t *testing.T) {
	val := ComputeHmac256("example message", "secret")
	if "yz4Ium5TwsqvwSpART/6KndMxhZOvsYkIOhu89bFn7w=" != val {
		t.Errorf("Incorrect match found: %v", val)
	}
}

func TestCompute256(t *testing.T) {
	val := Compute256("example message")
	if "rYTNCxD8Aoc4lxsHgSSuwqDnxtmGo4G+CzhvMr7oh68=" != val {
		t.Errorf("Incorrect match found: %v", val)
	}
}
