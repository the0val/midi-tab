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
}

// func TestTabFromMap(t *testing.T) {
// 	var keyMap [128][]string
//
// 	keyMap[45] = []string{"1"}
// 	keyMap[46] = []string{"-1'"}
// 	keyMap[47] = []string{"-1"}
// 	keyMap[49] = []string{"2"}
// 	keyMap[52] = []string{"3"}
// 	keyMap[55] = []string{"-3"}
// 	keyMap[56] = []string{"4"}
//
// 	tab := Tabulator(keyMap)
//
// 	tab.Tabulate()
// }
