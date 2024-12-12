package route

import (
	"reflect"
	"strings"
	"sync"
)

type __fastCopyable interface {
	copy(any) error
}

type __mapping = map[string]int

var __modelCache sync.Map

func copyContent[O any, T any](item *O) *T {
	t := new(T)
	fc, ok := any(item).(__fastCopyable)
	if ok {
		if fc.copy(t) == nil {
			return t
		}
	}
	__slowCopy(item, t)
	return t
}

func __slowCopy[O any, T any](from *O, to *T) {
	frv := reflect.ValueOf(from)
	for frv.Kind() == reflect.Ptr {
		frv = frv.Elem()
	}
	trv := reflect.ValueOf(to)
	for trv.Kind() == reflect.Ptr {
		trv = trv.Elem()
	}

	frt := frv.Type()
	trt := trv.Type()
	var frtMapping, trtMapping __mapping
	cache, loaded := __modelCache.Load(frt.String())
	if loaded {
		frtMapping = cache.(__mapping)
	} else {
		frtMapping = __mapping{}
		for i := 0; i < frt.NumField(); i++ {
			name := strings.Split(frt.Field(i).Tag.Get("json"), ",")[0]
			if name == "" {
				continue
			}
			frtMapping[name] = i
		}
		__modelCache.Store(frt.String(), frtMapping)
	}
	cache, loaded = __modelCache.Load(trt.String())
	if loaded {
		trtMapping = cache.(__mapping)
	} else {
		trtMapping = __mapping{}
		for i := 0; i < trt.NumField(); i++ {
			name := strings.Split(trt.Field(i).Tag.Get("json"), ",")[0]
			if name == "" {
				continue
			}
			trtMapping[name] = i
		}
		__modelCache.Store(trt.String(), trtMapping)
	}

	for oName, oIndex := range frtMapping {
		tIndex, ok := trtMapping[oName]
		if !ok {
			continue
		}
		trv.Field(tIndex).Set(frv.Field(oIndex))
	}
}
