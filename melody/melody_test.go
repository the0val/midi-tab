package melody

import (
	"reflect"
	"testing"
)

func TestRead(t *testing.T) {
	const testfile = "testdata/testmelody.mid"

	want := Melody{
		{60, 1},
		{60, 1},
		{67, 1},
		{67, 1},
		{69, 1},
		{69, 1},
		{67, 1},
		{65, 1},
		{65, 1},
		{64, 1},
		{64, 1},
		{62, 1},
		{62, 1},
		{60, 1},
	}
	got, err := ReadFile(testfile)

	if err != nil {
		t.Errorf("expected no error, got %q", err)
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v,\nexpected %v", got, want)
	}

	wantString := "C C G G A A G F F E E D D C"

	if got.String() != wantString {
		t.Errorf("expected string %q, got %q", wantString, got.String())
	}
}
