package jp

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Parser interface {
	ParseJSON(filePath, dotNotation string) (map[string]any, error)
}

type JSONParser struct {
}

type ArrayOutOfBoundsError struct {
	Index            int
	PathToArray      string
	AvailableIndexes int
}

type InvalidJsonKeyError struct {
	Key  string
	Path string
}

func (err ArrayOutOfBoundsError) Error() string {
	return fmt.Sprintf("array out of bounds in array \"%s\" for index %d, available indexes: %d", err.PathToArray, err.Index, err.AvailableIndexes)
}

func (err InvalidJsonKeyError) Error() string {
	return fmt.Sprintf("invalid json key \"%s\" for path \"%s\"", err.Key, err.Path)
}

// ParseJSON parses a JSON file and extracts a value using dot notation.
func (p JSONParser) ParseJSON(filePath, dotNotation string) (map[string]any, error) {
	v, err := openJsonFile(filePath)
	if err != nil {
		return nil, err
	}

	jsonPathParts := strings.Split(dotNotation, ".")

	jsonValue, err := TraverseJson(jsonPathParts, v)
	if err != nil {
		return nil, err
	}

	res := map[string]any{
		dotNotation: jsonValue,
	}

	return res, nil
}

func TraverseJson(jsonPathParts []string, currentValue interface{}) (interface{}, error) {
	currentPath := ""

	for _, part := range jsonPathParts {
		if currentPath != "" {
			currentPath = currentPath + "." + part
		} else {
			currentPath = part
		}

		switch t := currentValue.(type) {
		case map[string]interface{}:
			val, ok := t[part]
			if !ok {
				return nil, &InvalidJsonKeyError{Key: part, Path: currentPath}
			}

			currentValue = val
		case []interface{}:
			n, err := strconv.Atoi(part)
			if err != nil {
				return nil, err
			}

			arrayElementCount := len(t) - 1
			if n > arrayElementCount {
				return nil, &ArrayOutOfBoundsError{Index: n, AvailableIndexes: arrayElementCount, PathToArray: currentPath}
			}

			currentValue = t[n]
		default:
			return nil, fmt.Errorf("unknown type: %T while parsing JSON", t)
		}
	}

	return currentValue, nil
}

func openJsonFile(filePath string) (interface{}, error) {
	buf, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	// Unmarshal JSON into dynamic map/slice structure for easy traversal
	var v interface{}
	err = json.Unmarshal(buf, &v)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func New() Parser {
	return JSONParser{}
}
