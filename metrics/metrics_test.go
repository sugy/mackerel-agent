package metrics

import (
	"reflect"
	"testing"
)

func TestMerge(t *testing.T) {
	var v = Values(map[string]float64{
		"aa": 10,
	})
	var vv = Values(map[string]float64{
		"bb": 20,
		"cc": 30,
	})

	v.Merge(vv)

	if !reflect.DeepEqual(v, Values(map[string]float64{
		"aa": 10,
		"bb": 20,
		"cc": 30,
	})) {
		t.Errorf("somthing went wrong")
	}
}

func TestMergeValuesCustomIdentifiers(t *testing.T) {
	var v0 = Values(map[string]float64{
		"aa": 10,
	})
	var v1 = Values(map[string]float64{
		"bb": 20,
		"cc": 30,
	})
	var v2 = Values(map[string]float64{
		"dd": 40,
		"ee": 50,
	})
	var v3 = Values(map[string]float64{
		"ff": 60,
		"gg": 70,
	})

	v := MergeValuesCustomIdentifiers([]ValuesCustomIdentifier{
		ValuesCustomIdentifier{Values: v0},
	}, ValuesCustomIdentifier{Values: v1})

	if !reflect.DeepEqual(v, []ValuesCustomIdentifier{
		ValuesCustomIdentifier{
			Values: Values(map[string]float64{
				"aa": 10,
				"bb": 20,
				"cc": 30,
			}),
			CustomIdentifier: nil,
		}}) {
		t.Errorf("somthing went wrong")
	}

	customIdentifiers := "foo-bar"
	v = MergeValuesCustomIdentifiers(v, ValuesCustomIdentifier{Values: v2, CustomIdentifier: &customIdentifiers})

	if !reflect.DeepEqual(v, []ValuesCustomIdentifier{
		ValuesCustomIdentifier{
			Values: Values(map[string]float64{
				"aa": 10,
				"bb": 20,
				"cc": 30,
			}),
			CustomIdentifier: nil,
		},
		ValuesCustomIdentifier{
			Values: Values(map[string]float64{
				"dd": 40,
				"ee": 50,
			}),
			CustomIdentifier: &customIdentifiers,
		},
	}) {
		t.Errorf("somthing went wrong")
	}

	sameCustomIdentifiers := "foo-bar"
	v = MergeValuesCustomIdentifiers(v, ValuesCustomIdentifier{Values: v3, CustomIdentifier: &sameCustomIdentifiers})

	if !reflect.DeepEqual(v, []ValuesCustomIdentifier{
		ValuesCustomIdentifier{
			Values: Values(map[string]float64{
				"aa": 10,
				"bb": 20,
				"cc": 30,
			}),
			CustomIdentifier: nil,
		},
		ValuesCustomIdentifier{
			Values: Values(map[string]float64{
				"dd": 40,
				"ee": 50,
				"ff": 60,
				"gg": 70,
			}),
			CustomIdentifier: &customIdentifiers,
		},
	}) {
		t.Errorf("somthing went wrong")
	}
}
