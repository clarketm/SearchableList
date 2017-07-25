/*

Copyright 2017 Travis Clarke. All rights reserved.
Use of this source code is governed by a Apache-2.0
license that can be found in the LICENSE file.

*/

package searchablelist_test

import (
	"fmt"

	"github.com/clarketm/searchablelist"
)

func Example() {
	sl := searchablelist.New()
	e1 := sl.PushFront(1)
	sl.InsertAfter(2, e1)
	e4 := sl.PushBack(1)
	sl.InsertBefore(2, e4)

	fmt.Println(sl.Contains(e1))
	fmt.Println(sl.ContainsElement(e4))
	fmt.Println(sl.ContainsValue(1))

	firstOneElement := sl.FindFirst(1)
	fmt.Println(e1 == firstOneElement)

	lastOneElement := sl.FindLast(1)
	fmt.Println(e4 == lastOneElement)

	allOneElements := sl.FindAll(1)
	fmt.Println(e1 == allOneElements[0])
	fmt.Println(e4 == allOneElements[1])

	// Output:
	// true
	// true
	// true
	// true
	// true
	// true
	// true
}
