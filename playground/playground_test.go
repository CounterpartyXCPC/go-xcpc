package playground

import (
	"fmt"
	"testing"
)

type rc4Test struct {
	key, keystream []byte
}

var examp = []rc4Test{
	{
		[]byte{0x4b, 0x65, 0x79},
		[]byte{0xeb, 0x9f, 0x77, 0x81, 0xb7, 0x34, 0xca, 0x72, 0xa7, 0x19},
	},
	{
		[]byte{0x57, 0x69, 0x6b, 0x69},
		[]byte{0x60, 0x44, 0xdb, 0x6d, 0x41, 0xb7},
	},
}

func TestExperiment(t *testing.T) {
	fmt.Println("TestExperiment")
	Experiment()
}

func TestExperiment2(t *testing.T) {
	fmt.Println("TestExperiment2")
	Experiment2()
}
func TestProtoOneof(t *testing.T) {
	fmt.Println("TestProtoOneof")
	ProtoOneof()
}
