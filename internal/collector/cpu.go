package collector

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"win-temp-monitor/internal/domain"
)

type CPUCollector struct{}

func NewCPUCollector() *CPUCollector {
	return &CPUCollector{}
}

func (c *CPUCollector) Collect() (domain.CPUMetrics, error) {
	resp, err := http.Get("http://localhost:8085/data.json")
	if err != nil {
		return domain.CPUMetrics{}, err
	}
	defer resp.Body.Close()

	var root Node

	if err := json.NewDecoder(resp.Body).Decode(&root); err != nil {
		return domain.CPUMetrics{}, err
	}

	var usage, temp float64

	parseNode(&root, &usage, &temp)

	return domain.CPUMetrics{
		Usage: usage,
		Temp:  temp,
	}, nil
}

type Node struct {
	Text     string `json:"Text"`
	Value    string `json:"Value"`
	Children []Node `json:"Children"`
}

func parseNode(n *Node, usage *float64, temp *float64) {
	text := strings.ToLower(n.Text)

	// CPU 使用率
	if strings.Contains(text, "cpu total") && strings.Contains(n.Value, "%") {
		*usage = parseValue(n.Value)
	}

	// CPU 温度
	if strings.Contains(text, "tctl") && strings.Contains(text, "tdie") {
		*temp = parseValue(n.Value)
	}

	for i := range n.Children {
		parseNode(&n.Children[i], usage, temp)
	}
}

func parseValue(v string) float64 {
	v = strings.ReplaceAll(v, "C", "")
	v = strings.ReplaceAll(v, "%", "")
	v = strings.ReplaceAll(v, " \\u00B0C", "")
	v = strings.TrimSpace(v)
	var b strings.Builder
	for _, r := range v {
		if (r >= '0' && r <= '9') || r == '.' || r == '-' {
			b.WriteRune(r)
		}
	}

	f, err := strconv.ParseFloat(v, 64)
	if err != nil {
		return 0
	}
	return f
}
