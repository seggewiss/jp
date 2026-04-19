package test

import (
	"errors"
	"testing"

	"github.com/seggewiss/jp/pkg/jp"
)

// Happy Cases
func TestParseNumberValue(t *testing.T) {
	dn := "scalar.number"
	res := callParse(t, dn)

	switch rest := res[dn].(type) {
	case float64:
		if rest != float64(1337) {
			t.Errorf("Error parsing json: expected 1337, got %v", rest)
		}
	default:
		t.Errorf("Error parsing json: expected float64, got %T", rest)
	}
}

func TestParseStringValue(t *testing.T) {
	dn := "scalar.string"
	res := callParse(t, dn)

	switch rest := res[dn].(type) {
	case string:
		if rest != "foo" {
			t.Errorf("Error parsing json: expected \"foo\", got %q", rest)
		}
	default:
		t.Errorf("Error parsing json: expected string, got %T", rest)
	}
}

func TestParseNullValue(t *testing.T) {
	dn := "scalar.null"
	res := callParse(t, dn)

	switch rest := res[dn].(type) {
	case nil:
		if rest != nil {
			t.Errorf("Error parsing json: expected nil, got %v", rest)
		}
	default:
		t.Errorf("Error parsing json: expected nil, got %T", rest)
	}
}

func TestParseBooleanValue(t *testing.T) {
	dn := "scalar.boolean"
	res := callParse(t, dn)

	switch rest := res[dn].(type) {
	case bool:
		if rest != true {
			t.Errorf("Error parsing json: expected true, got %v", rest)
		}
	default:
		t.Errorf("Error parsing json: expected bool, got %T", rest)
	}
}

func TestParseScalarArray(t *testing.T) {
	dn := "scalar_array.1"
	res := callParse(t, dn)

	switch rest := res[dn].(type) {
	case string:
		if rest != "2" {
			t.Errorf("Error parsing json: expected \"2\", got %v", rest)
		}
	default:
		t.Errorf("Error parsing json: expected string, got %T", rest)
	}
}

func TestParseObjectArray(t *testing.T) {
	dn := "obj_array.1.bar"
	res := callParse(t, dn)

	switch rest := res[dn].(type) {
	case string:
		if rest != "biz" {
			t.Errorf("Error parsing json: expected \"biz\", got %v", rest)
		}
	default:
		t.Errorf("Error parsing json: expected string, got %T", rest)
	}
}

func TestParseDeepNotation(t *testing.T) {
	dn := "deep.a.0.b.0.c.d.0.e"
	res := callParse(t, dn)

	switch rest := res[dn].(type) {
	case string:
		if rest != "fuz" {
			t.Errorf("Error parsing json: expected \"fuz\", got %v", rest)
		}
	default:
		t.Errorf("Error parsing json: expected string, got %T", rest)
	}
}

// Unhappy cases
func TestTraverseJsonArrayOutOfBounds(t *testing.T) {
	json := map[string]any{
		"foo": []any{
			map[string]any{
				"bar": "foo",
			},
		},
	}

	res, err := jp.TraverseJson([]string{"foo", "10", "bar"}, json)
	if res != nil {
		t.Errorf("Array out of bound should return nil result, got %v", res)
	}

	if err == nil {
		t.Errorf("Array out of bound should return error got nil")
	}

	if _, ok := errors.AsType[*jp.ArrayOutOfBoundsError](err); !ok {
		t.Errorf("Wrong error type returned for array out of bounds %T", err)
	}
}

func TestTraverseJsonInvalidKey(t *testing.T) {
	json := map[string]any{
		"foo": []any{
			map[string]any{
				"bar": "foo",
			},
		},
	}

	res, err := jp.TraverseJson([]string{"foo", "0", "invalid"}, json)
	if res != nil {
		t.Errorf("Array out of bound should return nil result, got %v", res)
	}

	if err == nil {
		t.Errorf("Array out of bound should return error got nil")
	}

	if _, ok := errors.AsType[*jp.InvalidJsonKeyError](err); !ok {
		t.Errorf("Wrong error type returned for invalid json key %T", err)
	}
}

// callParse Helper function for calling jp ParseJSON
func callParse(t *testing.T, dn string) map[string]any {
	res, err := jp.New().ParseJSON("./testdata/happy.json", dn)
	if err != nil {
		t.Errorf("Error parsing json: %v", err)
	}
	return res
}
