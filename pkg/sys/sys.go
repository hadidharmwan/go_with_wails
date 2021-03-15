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

