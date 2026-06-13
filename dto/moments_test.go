package dto

import (
	"reflect"
	"strings"
	"testing"
)

func TestMomentsTimelineJSONTagsMatchFieldNames(t *testing.T) {
	types := []any{
		TimelineObject{},
		AppInfo{},
		WeappInfo{},
		ContentObject{},
		MediaList{},
		Media{},
		VideoSize{},
		URL{},
		Thumb{},
		Size{},
		VideoColdDLRule{},
		ActionInfo{},
		AppMsg{},
		Location{},
		StreamVideo{},
	}

	for _, value := range types {
		typ := reflect.TypeOf(value)
		for i := 0; i < typ.NumField(); i++ {
			field := typ.Field(i)
			if xmlTag := field.Tag.Get("xml"); xmlTag != "" {
				t.Fatalf("%s.%s still has xml tag %q", typ.Name(), field.Name, xmlTag)
			}

			jsonTag := field.Tag.Get("json")
			if jsonTag == "" {
				t.Fatalf("%s.%s is missing json tag", typ.Name(), field.Name)
			}

			jsonName, _, _ := strings.Cut(jsonTag, ",")
			if jsonName != field.Name {
				t.Fatalf("%s.%s json tag name = %q, want %q", typ.Name(), field.Name, jsonName, field.Name)
			}
		}
	}
}
