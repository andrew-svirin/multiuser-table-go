package system

// numToPercents - Convert numbers to percentage.
func numToPercents(part uint64, total uint64) float64 {
	return 100 * float64(total-part) / float64(total)
}

// bToKb - Format Byte to Kb
func bToKb(b uint64) uint64 {
	return b / 1024
}
