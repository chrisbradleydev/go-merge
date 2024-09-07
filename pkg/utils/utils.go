package utils

func MergeMaps(a, b map[string]interface{}) map[string]interface{} {
	out := make(map[string]interface{}, len(a))
	for k, v := range a {
		out[k] = v
	}
	for k, v := range b {
		switch vt := v.(type) {
		case map[string]interface{}:
			if bv, ok := out[k]; ok {
				if bv, ok := bv.(map[string]interface{}); ok {
					out[k] = MergeMaps(bv, vt)
					continue
				}
			}
			out[k] = vt
		case []interface{}:
			if bv, ok := out[k]; ok {
				if bv, ok := bv.([]interface{}); ok {
					out[k] = MergeSlices(bv, vt)
					continue
				}
			}
			out[k] = vt
		default:
			out[k] = v
		}
	}
	return out
}

func MergeSlices(a, b []interface{}) []interface{} {
	merged := make([]interface{}, len(a))
	copy(merged, a)

	nameMap := make(map[string]int)
	for i, item := range merged {
		if m, ok := item.(map[string]interface{}); ok {
			if name, ok := m["name"].(string); ok {
				nameMap[name] = i
			}
		}
	}

	for _, item := range b {
		if m, ok := item.(map[string]interface{}); ok {
			if name, ok := m["name"].(string); ok {
				if i, exists := nameMap[name]; exists {
					merged[i] = MergeMaps(merged[i].(map[string]interface{}), m)
				} else {
					merged = append(merged, m)
					nameMap[name] = len(merged) - 1
				}
			} else {
				merged = append(merged, m)
			}
		} else {
			merged = append(merged, item)
		}
	}

	return merged
}
