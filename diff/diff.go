package diff

import (
	"encoding/json"
	"fmt"
	"reflect"
	"sync"
)

func contains(intSet map[int]struct{}, value int) bool {
	_, exists := intSet[value]
	return exists
}

// DiffDetail holds old and new values for modified fields
type DiffDetail struct {
	OldValue interface{} `json:"oldValue"`
	NewValue interface{} `json:"newValue"`
}

// DiffResult holds the result of the comparison
type DiffResult struct {
	Added    map[string]interface{} `json:"added"`
	Deleted  map[string]interface{} `json:"deleted"`
	Modified map[string]DiffDetail  `json:"modified"`
	addMutex sync.Mutex
	delMutex sync.Mutex
	modMutex sync.Mutex
}

func (d *DiffResult) AddAdded(path string, value interface{}) {
	d.addMutex.Lock()
	defer d.addMutex.Unlock()
	d.Added[path] = value
}

func (d *DiffResult) AddDeleted(path string, value interface{}) {
	d.delMutex.Lock()
	defer d.delMutex.Unlock()
	d.Deleted[path] = value
}

func (d *DiffResult) AddModified(path string, oldValue, newValue interface{}) {
	d.modMutex.Lock()
	defer d.modMutex.Unlock()
	d.Modified[path] = DiffDetail{OldValue: oldValue, NewValue: newValue}
}

// CompareJSON recursively compares two JSON objects and returns a diff
func CompareJSON(json1, json2 map[string]interface{}, levelSet map[int]struct{}) *DiffResult {
	diff := &DiffResult{
		Added:    make(map[string]interface{}),
		Deleted:  make(map[string]interface{}),
		Modified: make(map[string]DiffDetail),
	}

	var wg sync.WaitGroup

	compareObjects("", json1, json2, diff, &wg, 1, levelSet)

	wg.Wait()

	return diff
}

// compareObjects compares two maps recursively
func compareObjects(path string, json1, json2 map[string]interface{}, diff *DiffResult, wg *sync.WaitGroup, level int, levelSet map[int]struct{}) {
	for key, val1 := range json1 {
		fullPath := path + key
		if val2, ok := json2[key]; !ok {
			diff.AddDeleted(fullPath, val1)
		} else {
			val := val1
			// if this is the deepest object level, parallellize
			if contains(levelSet, level) {
				wg.Add(1)
				go func(val1 interface{}) {
					defer wg.Done()
					compareValues(fullPath, val1, val2, diff, wg, level+1, levelSet)
				}(val)
			} else {
				compareValues(fullPath, val1, val2, diff, wg, level+1, levelSet)
			}
		}
	}

	for key, val2 := range json2 {
		fullPath := path + key
		if _, ok := json1[key]; !ok {
			diff.AddAdded(fullPath, val2)
		}
	}
}

func areValuesEqual(val1, val2 interface{}) bool {
	switch v1 := val1.(type) {
	case int:
		if v2, ok := val2.(int); ok {
			return v1 == v2
		}
	case float64:
		if v2, ok := val2.(float64); ok {
			return v1 == v2
		}
	case string:
		if v2, ok := val2.(string); ok {
			return v1 == v2
		}
	case nil:
		return val2 == nil
	default:
		// For complex types, use reflect.DeepEqual as a fallback
		return reflect.DeepEqual(val1, val2)
	}
	return false
}

// compareValues compares values (and handles nested objects/arrays)
func compareValues(path string, val1, val2 interface{}, diff *DiffResult, wg *sync.WaitGroup, level int, levelSet map[int]struct{}) {
	if reflect.TypeOf(val1) != reflect.TypeOf(val2) {
		diff.AddModified(path, val1, val2)
		return
	}

	switch v1 := val1.(type) {
	case map[string]interface{}:
		compareObjects(path+".", v1, val2.(map[string]interface{}), diff, wg, level, levelSet)
	case []interface{}:
		compareArrays(path, v1, val2.([]interface{}), diff, wg, level, levelSet)
	default:
		if !areValuesEqual(val1, val2) {
			diff.AddModified(path, val1, val2)
		}
	}
}

// compareArrays compares two arrays element-wise
func compareArrays(path string, arr1, arr2 []interface{}, diff *DiffResult, wg *sync.WaitGroup, level int, levelSet map[int]struct{}) {
	len1, len2 := len(arr1), len(arr2)
	minLen := len1
	if len2 < len1 {
		minLen = len2
	}

	for i := 0; i < minLen; i++ {
		elemPath := fmt.Sprintf("%s[%d]", path, i)
		compareValues(elemPath, arr1[i], arr2[i], diff, wg, level, levelSet)
	}

	for i := minLen; i < len2; i++ {
		elemPath := fmt.Sprintf("%s[%d]", path, i)
		diff.AddAdded(elemPath, arr2[i])
	}

	for i := minLen; i < len1; i++ {
		elemPath := fmt.Sprintf("%s[%d]", path, i)
		diff.AddDeleted(elemPath, arr1[i])
	}
}

// parseJSON parses a JSON string into a map[string]interface{}
func ParseJSON(jsonStr string) (map[string]interface{}, error) {
	var result map[string]interface{}
	err := json.Unmarshal([]byte(jsonStr), &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
