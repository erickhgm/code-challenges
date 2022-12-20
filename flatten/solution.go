package flatten

import "reflect"

func Flatten(nn []any) []any {
	var flatt []any
	var flatten func(n []any)

	flatten = func(n []any) {
		for _, v := range n {
			if reflect.TypeOf(v).Kind() != reflect.Slice &&
				reflect.TypeOf(v).Kind() != reflect.Array {
				flatt = append(flatt, v)
			} else {
				temp, _ := v.([]any)
				flatten(temp)
			}

		}
	}
	flatten(nn)
	return flatt
}
