package gatherer

import (
	"strconv"
)

type VmStat struct {
}

func (m *VmStat) Read() (DataPoints, error) {
	lines, err := ReadFileParts("/proc/vmstat")
	if err != nil {
		return nil, err
	}
	result := make([]DataPoint, 0)
	for _, parts := range lines {
		if len(parts) >= 2 {
			value, err := strconv.ParseInt(parts[1], 10, 64)
			if err != nil {
				return nil, err
			}

			result = append(result, DataPoint{
				Name:  convertSaveName(parts[0]),
				Value: float64(value),
			})
		}
	}

	return result, nil
}

func (m *VmStat) GetName() string {
	return "VmStat"
}
