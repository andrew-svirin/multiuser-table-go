package system

// Memory usage for whole the system.

import (
	"runtime"
)

// memoryUsage- Memory usage value.
func memoryUsage() uint64 {
	var m runtime.MemStats

	runtime.ReadMemStats(&m)

	return m.Sys
}

// MemoryUsageKb - memory usage in Kb
func MemoryUsageKb() uint64 {
	return bToKb(memoryUsage())
}
