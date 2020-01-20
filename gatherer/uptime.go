package gatherer

import (
	"strconv"
)

type Uptime struct {
}

func (m *Uptime) Read() (DataPoints, error) {
	lines, err := ReadFileParts("/proc/uptime")
	if err != nil {
		return nil, err
	}

	result := make([]DataPoint, 0)
	for _, parts := range lines {
		if len(parts) >= 2 {
			value0, err := strconv.ParseFloat(parts[0], 64)
			if err != nil {
				return nil, err
			}
			value1, err := strconv.ParseFloat(parts[1], 64)
			if err != nil {
				return nil, err
			}

			result = append(result, DataPoint{
				Name:  "uptime",
				Value: value0,
			})
			result = append(result, DataPoint{
				Name:  "idle",
				Value: value1,
			})
		}
	}

	return result, nil
}

func (m *Uptime) GetName() string {
	return "Uptime"
}
