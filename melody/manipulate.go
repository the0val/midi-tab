package melody

// Tanspose 'moves' all notes up by steps.
// Use negative numbers to transpose down
func (m *Melody) Transpose(steps int) {
	for i := range *m {
		(*m)[i].Key += uint8(steps)
	}
}
