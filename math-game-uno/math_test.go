package mathgameuno

import (
	"reflect"
	"testing"
)

func TestMathQuis(t *testing.T){
	actual,_:= ReadFile("test.csv")

	xs := reflect.TypeOf(actual)
	
}