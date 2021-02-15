package melody

import (
	"regexp"
	"strconv"
	"strings"
)

type TabErr string

func (e TabErr) Error() string {
	return string(e)
}

const ErrorBadTemplate = TabErr("Could not parse provided temlate")

type Tabulator [128][]string

func (t Tabulator) Tabulate(m Melody) string {
	var outStr string
	for _, n := range m {
		if ss := t[n.Key]; len(ss) == 0 {
			outStr += "XX "
		} else {
			outStr += ss[0] + " "
		}
	}
	// trim trailing space
	return outStr[:len(outStr)-1]
}

func ParseTabTemplate(data []byte) (Tabulator, error) {
	var ret Tabulator

	newPattern := regexp.MustCompile("^([0-9]+):(.+)$")

	recentKey := 0
	for _, line := range strings.Split(string(data), "\n") {
		if newPattern.MatchString(line) {
			subMatch := newPattern.FindStringSubmatch(line)

			note := subMatch[1]
			key, err := strconv.Atoi(note)
			if err != nil {
				return Tabulator{}, ErrorBadTemplate
			}

			ret[key] = append(ret[key], "")
			line = subMatch[2]
			recentKey = key
		}

		ret[recentKey][len(ret[recentKey])-1] = ret[recentKey][len(ret[recentKey])-1] + line
	}

	return ret, nil
}

func (t Tabulator) Export() []byte {
	return []byte("")
}
