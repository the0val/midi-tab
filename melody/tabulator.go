package melody

type Tabulator interface {
	Tab(Melody) string
}

type mapTab [128]string

func (t mapTab) Tab(m Melody) string {
	var outStr string
	for n := range m {
		outStr += t[n] + " "
	}
	// trim trailing space
	return outStr[:len(outStr)-1]
}

func NewTabulator(mapping [128][]string) Tabulator {
	return mapTab(mapping)
}
