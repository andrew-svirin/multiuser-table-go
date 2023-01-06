package system

// CPU usage for whole the system.

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// cPUUsage - CPU usage.
func cPUUsage() (idle, total uint64) {
	contents, err := os.ReadFile("/proc/stat")
	if err != nil {
		return
	}
	lines := strings.Split(string(contents), "\n")
	for _, line := range lines {
		fields := strings.Fields(line)
		if fields[0] == "cpu" {
			numFields := len(fields)
			for i := 1; i < numFields; i++ {
				val, err := strconv.ParseUint(fields[i], 10, 64)
				if err != nil {
					fmt.Println("Error: ", i, fields[i], err)
				}
				total += val // tally up all the numbers to get total ticks
				if i == 4 {  // idle is the 5th field in the cpu line
					idle = val
				}
			}
			return
		}
	}
	return
}

// CPUUsagePercents - CPU usage in percents.
func CPUUsagePercents() float64 {
	return numToPercents(cPUUsage())
}
