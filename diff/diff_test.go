package diff

import (
	"JSONCompare/config"
	"reflect"
	"testing"
)

func TestCompareJSON(t *testing.T) {
	tests := []struct {
		name     string
		json1    string
		json2    string
		expected *DiffResult
	}{
		{
			name: "No Differences",
			json1: `{
				"key1": "value1",
				"key2": {
					"subkey1": "subvalue1"
				}
			}`,
			json2: `{
				"key1": "value1",
				"key2": {
					"subkey1": "subvalue1"
				}
			}`,
			expected: &DiffResult{
				Added:    map[string]interface{}{},
				Deleted:  map[string]interface{}{},
				Modified: map[string]DiffDetail{},
			},
		},
		{
			name: "Added Fields",
			json1: `{
				"key1": "value1"
			}`,
			json2: `{
				"key1": "value1",
				"key2": "value2"
			}`,
			expected: &DiffResult{
				Added: map[string]interface{}{
					"key2": "value2",
				},
				Deleted:  map[string]interface{}{},
				Modified: map[string]DiffDetail{},
			},
		},
		{
			name: "Deleted Fields",
			json1: `{
				"key1": "value1",
				"key2": "value2"
			}`,
			json2: `{
				"key1": "value1"
			}`,
			expected: &DiffResult{
				Added: map[string]interface{}{},
				Deleted: map[string]interface{}{
					"key2": "value2",
				},
				Modified: map[string]DiffDetail{},
			},
		},
		{
			name: "Modified Fields",
			json1: `{
				"key1": "value1"
			}`,
			json2: `{
				"key1": "value2"
			}`,
			expected: &DiffResult{
				Added:   map[string]interface{}{},
				Deleted: map[string]interface{}{},
				Modified: map[string]DiffDetail{
					"key1": {
						OldValue: "value1",
						NewValue: "value2",
					},
				},
			},
		},
		{
			name: "Nested Objects",
			json1: `{
				"key1": {
					"subkey1": "subvalue1"
				}
			}`,
			json2: `{
				"key1": {
					"subkey1": "subvalue2"
				}
			}`,
			expected: &DiffResult{
				Added:   map[string]interface{}{},
				Deleted: map[string]interface{}{},
				Modified: map[string]DiffDetail{
					"key1.subkey1": {
						OldValue: "subvalue1",
						NewValue: "subvalue2",
					},
				},
			},
		},
		{
			name: "Arrays",
			json1: `{
				"key1": [1, 2, 3]
			}`,
			json2: `{
				"key1": [1, 2, 4]
			}`,
			expected: &DiffResult{
				Added:   map[string]interface{}{},
				Deleted: map[string]interface{}{},
				Modified: map[string]DiffDetail{
					"key1[2]": {
						OldValue: float64(3),
						NewValue: float64(4),
					},
				},
			},
		},
		{
			name: "Complex Changes with different order and different data types",
			json1: `{
				"key1": {
					"subkey2": [1, 2],
					"subkey1": "value1",
					"fine":34.0,
					"numbers":45.0002,
					"money": 125.00
				},
				"key2": "value2"
			}`,
			json2: `{
				"key1": {
					"subkey1": "value2",
					"money": 125.0,
					"fine":34,
					"numbers":45,
					"subkey2": [1, 3]
				},
				"key3": "value3"
			}`,
			expected: &DiffResult{
				Added: map[string]interface{}{
					"key3": "value3",
				},
				Deleted: map[string]interface{}{
					"key2": "value2",
				},
				Modified: map[string]DiffDetail{
					"key1.subkey1": {
						OldValue: "value1",
						NewValue: "value2",
					},
					"key1.subkey2[1]": {
						OldValue: float64(2),
						NewValue: float64(3),
					},
					"key1.numbers": {
						OldValue: float64(45.0002),
						NewValue: float64(45),
					},
				},
			},
		},
	}

	levelSet := make(map[int]struct{})

	for _, level := range config.LEVELS {
		levelSet[level] = struct{}{}
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			obj1, err := ParseJSON(tt.json1)
			if err != nil {
				t.Fatalf("Error parsing JSON 1: %v", err)
			}
			obj2, err := ParseJSON(tt.json2)
			if err != nil {
				t.Fatalf("Error parsing JSON 2: %v", err)
			}

			result := CompareJSON(obj1, obj2, levelSet)

			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("CompareJSON() = %v, want %v", result, tt.expected)
			}
		})
	}
}
