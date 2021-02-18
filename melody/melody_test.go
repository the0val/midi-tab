package melody

import (
	"os"
	"reflect"
	"testing"
)

func TestStringer(t *testing.T) {
	melody := Melody{
		{45, 1},
		{46, 1},
		{47, 1},
		{48, 1},
		{49, 1},
		{50, 1},
		{51, 1},
		{52, 1},
		{53, 1},
		{54, 1},
		{55, 1},
		{56, 1},
	}

	want := "A A# B C C# D D# E F F# G G#"
	got := melody.String()
	if got != want {
		t.Errorf("Expected string %q, got %q", want, got)
	}
}

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
	t.Run("ReadFile", func(t *testing.T) {
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
	})
	t.Run("Reader", func(t *testing.T) {
		file, err := os.Open(testfile)

		if err != nil {
			t.Errorf("expected no error, got %q", err)
		}

		got, err := ReadReader(file)

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v,\nexpected %v", got, want)
		}
	})
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
