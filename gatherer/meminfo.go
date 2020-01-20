package gatherer

import (
	"strconv"
)

type MemInfo struct {
}

func (m *MemInfo) Read() (DataPoints, error) {
	lines, err := ReadFileParts("/proc/meminfo")
	if err != nil {
		return nil, err
	}
	result := make([]DataPoint, 0)
	for _, parts := range lines {
		if len(parts) >= 3 {
			value, err := strconv.ParseInt(parts[1], 10, 64)
			if err != nil {
				return nil, err
			}
			unit, err := convertUnit(parts[2])
			if err != nil {
				return nil, err
			}

			result = append(result, DataPoint{
				Name:  convertSaveName(parts[0]),
				Value: float64(value) * float64(unit),
			})
		}
	}

	return result, nil
}

func (m *MemInfo) GetName() string {
	return "MemInfo"
}
