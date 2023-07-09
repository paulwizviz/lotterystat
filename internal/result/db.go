package result

import (
	"reflect"
	"strings"
)

type drawType interface {
	EuroDraw | Set4LifeDraw
}

type structTag struct {
	FieldName string
	Tag       string
}

func sqliteTags[T drawType](typ *T) []structTag {
	ev := reflect.Indirect(reflect.ValueOf(typ))
	tags := []structTag{}
	for i := 0; i < ev.Type().NumField(); i++ {
		tag := structTag{}
		tag.FieldName = ev.Type().Field(i).Name
		t := ev.Type().Field(i).Tag
		tElems := strings.Split(string(t), " ")
		for _, tElem := range tElems {
			if strings.Contains(tElem, "sqlite") {
				sElems := strings.Split(tElem, ":")
				tag.Tag = sElems[1][1 : len(sElems[1])-1]
			}
		}
		tags = append(tags, tag)
	}
	return tags
}
