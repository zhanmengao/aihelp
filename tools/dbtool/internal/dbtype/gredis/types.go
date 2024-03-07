package gredis

type uniqVal struct {
	Key   string
	Val   interface{}
	Slice []interface{}
}

func uniqAdd(m []*uniqVal, k string, v, sv interface{}) []*uniqVal {
	exists := false
	for _, kk := range m {
		if kk.Key == k {
			exists = true
			kk.Val = v
			kk.Slice = append(kk.Slice, sv)
			break
		}
	}
	if !exists {
		m = append(m, &uniqVal{Key: k, Val: v, Slice: []interface{}{sv}})
	}
	return m
}
