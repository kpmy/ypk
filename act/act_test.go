package act

import (
	"testing"
)

func TestAct(t *testing.T) {
	Seq().Do(func(x interface{}) (interface{}, error) {
		return x.(int) + 1, nil
	}).Do(func(x interface{}) (interface{}, error) {
		t.Log(x)
		return nil, nil
	}).Run(0)
}
