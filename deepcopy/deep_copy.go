package main

import (
	"reflect"
)

func DeepCopy[T any](src, dst *T) {
	srv := reflect.ValueOf(src)
	srv = srv.Elem()
	drv := reflect.ValueOf(dst)
	drv = drv.Elem()
	addrMap := map[uint64]reflect.Value{}
	handleStruct(srv, drv, addrMap)
}

func genKey(src reflect.Value) uint64 {
	switch src.Kind() {
	case reflect.Pointer, reflect.Chan, reflect.Map, reflect.UnsafePointer, reflect.Func, reflect.Slice:
		return uint64(src.Pointer())<<6 + uint64(src.Kind())<<1
	case reflect.Struct, reflect.Interface:
		return uint64(src.UnsafeAddr())<<6 + uint64(src.Kind())<<1 + 1
	default:
		return 0
	}
}

func handle(src, dst reflect.Value, addrMap map[uint64]reflect.Value) {
	switch src.Kind() {
	case reflect.Struct:
		handleStruct(src, dst, addrMap)
	case reflect.Interface:
		handleInterface(src, dst, addrMap)
	case reflect.Pointer:
		handlePointer(src, dst, addrMap)
	case reflect.Array:
		handleArray(src, dst, addrMap)
	case reflect.Slice:
		handleSlice(src, dst, addrMap)
	case reflect.Chan:
		dst.Set(reflect.MakeChan(src.Type(), src.Cap()))
	case reflect.Map:
		dst.Set(reflect.MakeMap(src.Type()))
		handleMap(src, dst, addrMap)
	default:
		dst.Set(src)
	}
}

func handlePointer(src, dst reflect.Value, addrMap map[uint64]reflect.Value) {
	src = src.Elem()
	addr := genKey(src)
	ndst, ok := addrMap[addr]
	if !ok {
		ndst = reflect.New(src.Type()).Elem()
		if addr > 0 {
			addrMap[addr] = ndst
		}
		handle(src, ndst, addrMap)
	}
	dst.Set(ndst.Addr())
}

func handleStruct(src, dst reflect.Value, addrMap map[uint64]reflect.Value) {
	srcAddr := genKey(src)
	if srcAddr > 0 {
		addrMap[srcAddr] = dst
	}
	for i := 0; i < src.NumField(); i++ {
		srcf := src.Field(i)
		addr := genKey(srcf)
		dstf, ok := addrMap[addr]
		if !ok {
			ndstf := dst.Field(i)
			if addr > 0 {
				addrMap[addr] = ndstf
			}
			handle(srcf, ndstf, addrMap)
		} else {
			dst.Field(i).Set(dstf)
		}
	}
}

func handleInterface(src, dst reflect.Value, addrMap map[uint64]reflect.Value) {
	src = src.Elem()
	handle(src, dst, addrMap)
}

func handleArray(src, dst reflect.Value, addrMap map[uint64]reflect.Value) {
	for i := 0; i < src.Len(); i++ {
		srcf := src.Index(i)
		addr := genKey(srcf)
		dstf, ok := addrMap[addr]
		if !ok {
			ndstf := dst.Index(i)
			if addr > 0 {
				addrMap[addr] = ndstf
			}
			handle(srcf, ndstf, addrMap)
		} else {
			dst.Index(i).Set(dstf)
		}
	}
}

func handleSlice(src, dst reflect.Value, addrMap map[uint64]reflect.Value) {
	srcAddr := genKey(src)
	if srcAddr > 0 {
		addrMap[srcAddr] = dst
	}
	dst.Grow(src.Len() - dst.Len())
	dst.SetLen(src.Len())
	for i := 0; i < src.Len(); i++ {
		srcf := src.Index(i)
		addr := genKey(srcf)
		dstf, ok := addrMap[addr]
		if !ok {
			ndstf := dst.Index(i)
			if addr > 0 {
				addrMap[addr] = ndstf
			}
			handle(srcf, ndstf, addrMap)
		} else {
			dst.Index(i).Set(dstf)
		}
	}
}

func handleMap(src, dst reflect.Value, addrMap map[uint64]reflect.Value) {
	srcAddr := genKey(src)
	if srcAddr > 0 {
		addrMap[srcAddr] = dst
	}
	iter := src.MapRange()
	for iter.Next() {
		k := iter.Key()
		v := iter.Value()

		kAddr := genKey(k)
		kdst, ok := addrMap[kAddr]
		if !ok {
			kdst = reflect.New(k.Type()).Elem()
			if kAddr > 0 {
				addrMap[kAddr] = kdst
			}
			handle(k, kdst, addrMap)
		}

		vAddr := genKey(v)
		vdst, ok := addrMap[vAddr]
		if !ok {
			vdst = reflect.New(v.Type()).Elem()
			if vAddr > 0 {
				addrMap[vAddr] = vdst
			}
			handle(v, vdst, addrMap)
		}
		dst.SetMapIndex(kdst, vdst)
	}
}
