package monitoring

import (
	"github.com/shirou/gopsutil/mem"
)

type RamInfo struct {
	Total uint64 `json:"total"`
	Used  uint64 `json:"used"`
}

func Ram() RamInfo {
	raminfo := RamInfo{}
	v, err := mem.VirtualMemory()
	if err != nil {
		raminfo.Total = 0
		raminfo.Used = 0
	} else {
		raminfo.Total = v.Total
		raminfo.Used = v.Used
	}
	return raminfo
}
func Swap() RamInfo {
	swapinfo := RamInfo{}
	s, err := mem.SwapMemory()
	if err != nil {
		swapinfo.Total = 0
		swapinfo.Used = 0
	} else {
		swapinfo.Total = s.Total
		swapinfo.Used = s.Used
	}
	return swapinfo
}
