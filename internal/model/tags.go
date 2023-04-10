package model

import (
	"reflect"
	"strings"
)

type DrawTypeConstraint interface {
	EuroDraw
}

func SqliteTags[T DrawTypeConstraint](typ *T) map[string]string {
	ev := reflect.Indirect(reflect.ValueOf(typ))
	tags := make(map[string]string)
	for i := 0; i < ev.Type().NumField(); i++ {
		fn := ev.Type().Field(i).Name
		tag := ev.Type().Field(i).Tag
		tItems := strings.Split(string(tag), " ")
		for _, ti := range tItems {
			if strings.Contains(ti, "sqlite") {
				si := strings.Split(ti, ":")
				tags[fn] = si[1][1 : len(si[1])-1]
			}
		}
	}
	return tags
}
