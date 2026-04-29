package main

import (
	"fmt"
	"time"
	"win-temp-monitor/internal/collector"
)

func main() {
	c := collector.NewCPUCollector()

	for {
		m, err := c.Collect()
		if err != nil {
			fmt.Println("error: ", err)
			time.Sleep(time.Second)
			continue
		}

		fmt.Printf("CPU: %.2f%% | Temp: %.2f℃\n",
			m.Usage,
			m.Temp,
		)
		time.Sleep(time.Second)
	}
}
