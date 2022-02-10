package sysMeta

import (
	"encoding/json"
	"fmt"
	"github.com/denisbrodbeck/machineid"
	"github.com/pbnjay/memory"
	"os"
	"runtime"
)

func NumberOfCore() int {
	return runtime.NumCPU()
}

type MemoryStat struct {
	Total uint64
	Free  uint64
}

func Memories() MemoryStat {
	ms := MemoryStat{}
	ms.Total = memory.TotalMemory()
	ms.Free = memory.FreeMemory()
	return ms
}

type SysMeta struct {
	TotalMemory  uint64 `json:"total_memory"`
	FreeMemory   uint64 `json:"free_memory"`
	NumberOfCore int    `json:"number_of_core"`
	MachineId    string	`json:"machine_id"`
	HostName     string `json:"host_name"`
}

func HostName() string {
	hostname, err := os.Hostname()
	if err != nil {
		fmt.Println("Fail to extract hostname")
	}
	return hostname
}

func MachineId() string {
	id, err := machineid.ID()
	if err != nil {
		fmt.Println("Fail to extract machine id")
	}
	return id
}

func GetMeta() SysMeta {
	sm := SysMeta{}
	ms := Memories()
	sm.TotalMemory = ms.Total
	sm.FreeMemory = ms.Free
	sm.HostName = HostName()
	sm.MachineId = MachineId()
	sm.NumberOfCore = NumberOfCore()
	return sm
}

func GetMetaString() (string, error) {
	meta := GetMeta()
	metaBytes, err := json.Marshal(&meta)
	if err != nil{
		return "",err
	} else {
		return string(metaBytes), err
	}
	return "", nil
}
