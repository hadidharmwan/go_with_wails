package sys

import (
	"math"
	"time"

	"github.com/shirou/gopsutil/cpu"
	"github.com/wailsapp/wails"
)

type Stats struct {
	log *wails.CustomLogger
}

type CPUUsage struct {
	Average int `json:"avg"`
}

func (s *Stats) WailsInit(runtime *wails.Runtime) error {
	s.log = runtime.Log.New("Stats")
	
	go func() {
		for {
			runtime.Events.Emit("cpu_usage", s.GetCPUUsage())
			time.Sleep(1 * time.Second)
		}
	}()

	return nil
}

func (s *Stats) GetCPUUsage() *CPUUsage {
	percent, err := cpu.Percent(1*time.Second, false)
	if err != nil {
		s.log.Errorf("unable to get cpu stats: %s", err.Error())
		return nil
	}

	return &CPUUsage{
		Average: int(math.Round(percent[0])),
	}
}