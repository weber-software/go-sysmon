package gatherer

import (
	"errors"
	"strconv"
)

type LoadAvg struct {
}

func (m *LoadAvg) Read() (DataPoints, error) {
	lines, err := ReadFileParts("/proc/loadavg")
	if err != nil {
		return nil, err
	}

	result := make([]DataPoint, 0)
	for _, parts := range lines {
		if len(parts) >= 5 {
			value0, err := strconv.ParseFloat(parts[0], 64)
			if err != nil {
				return nil, err
			}
			value1, err := strconv.ParseFloat(parts[1], 64)
			if err != nil {
				return nil, err
			}
			value2, err := strconv.ParseFloat(parts[2], 64)
			if err != nil {
				return nil, err
			}

			subParts := splitRemoveEmpty(parts[3], "/")
			if len(subParts) < 2 {
				return nil, errors.New("invalid format")
			}

			sub0, err := strconv.ParseFloat(subParts[0], 64)
			if err != nil {
				return nil, err
			}
			sub1, err := strconv.ParseFloat(subParts[1], 64)
			if err != nil {
				return nil, err
			}

			value4, err := strconv.ParseFloat(parts[4], 64)
			if err != nil {
				return nil, err
			}

			result = append(result, DataPoint{
				Name:  "LoadAvg1",
				Value: value0,
			})
			result = append(result, DataPoint{
				Name:  "LoadAvg5",
				Value: value1,
			})
			result = append(result, DataPoint{
				Name:  "LoadAvg15",
				Value: value2,
			})
			result = append(result, DataPoint{
				Name:  "RunnableSchedulingEntities",
				Value: sub0,
			})
			result = append(result, DataPoint{
				Name:  "TotalSchedulingEntities",
				Value: sub1,
			})
			result = append(result, DataPoint{
				Name:  "MaxPid",
				Value: value4,
			})
		}
	}

	return result, nil
}

func (m *LoadAvg) GetName() string {
	return "LoadAvg"
}
