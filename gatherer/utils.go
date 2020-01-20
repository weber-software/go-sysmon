package gatherer

import (
	"errors"
	"io/ioutil"
	"strings"
)

type DataPoint struct {
	Name  string
	Value float64
}
type DataPoints []DataPoint

type Gatherer interface {
	Read() (DataPoints, error)
	GetName() string
}

func splitRemoveEmpty(str string, seperator string) []string {
	parts := strings.Split(str, seperator)
	result := make([]string, 0)
	for _, part := range parts {
		if len(part) > 0 {
			result = append(result, part)
		}
	}
	return result
}

func convertUnit(unit string) (int, error) {
	switch unit {
	case "kB":
		return 1024, nil
	default:
		return 0, errors.New("unknown unit: " + unit)
	}
}

func convertSaveName(str string) string {
	str = strings.ReplaceAll(str, ":", "")
	str = strings.ReplaceAll(str, "(", "_")
	str = strings.ReplaceAll(str, ")", "_")
	return str
}

func ReadFileParts(fileName string) ([][]string, error) {
	dat, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}

	result := make([][]string, 0)
	lines := strings.Split(string(dat), "\n")
	for _, line := range lines {
		parts := splitRemoveEmpty(line, " ")
		result = append(result, parts)
	}

	return result, nil
}