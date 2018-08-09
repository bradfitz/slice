// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build go1.8

package slice

import (
	"fmt"
	"reflect"
	"sort"
)

// Sort sorts the provided slice using the function less.
// If slice is not a slice, Sort panics.
//
// Deprecated: use Go 1.8+'s sort.Slice in the standard library
// instead. This package predates adding it to the standard library
// and is no longer maintained for new architectures.
func Sort(slice interface{}, less func(i, j int) bool) {
	sort.Slice(slice, less)
}

// SortInterface returns a sort.Interface to sort the provided slice
// using the function less.
func SortInterface(slice interface{}, less func(i, j int) bool) sort.Interface {
	sv := reflect.ValueOf(slice)
	if sv.Kind() != reflect.Slice {
		panic(fmt.Sprintf("slice.Sort called with non-slice value of type %T", slice))
	}
	return &funcs{
		length: sv.Len(),
		less:   less,
		swap:   reflect.Swapper(slice),
	}
}

type funcs struct {
	length int
	less   func(i, j int) bool
	swap   func(i, j int)
}

func (f *funcs) Len() int           { return f.length }
func (f *funcs) Less(i, j int) bool { return f.less(i, j) }
func (f *funcs) Swap(i, j int)      { f.swap(i, j) }
