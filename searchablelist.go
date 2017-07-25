/*

Copyright 2017 Travis Clarke. All rights reserved.
Use of this source code is governed by a Apache-2.0
license that can be found in the LICENSE file.

SEARCHABLE LIST
	Package `searchablelist` implements the `list` package doubly linked list
	and extends it with search methods

    type SearchableList
        New () *SearchableList
        ContainsElement (t *Element) bool
        Contains (t *Element) bool
        ContainsValue (v interface{}) bool
        FindFirst (v interface{}) *Element
        FindLast (v interface{}) *Element
        FindAll (v interface{}) []*Element

    type Element
        func (e *Element) Next() *Element
        func (e *Element) Prev() *Element

    type List
        func (l *List) Back() *Element
        func (l *List) Front() *Element
        func (l *List) Init() *List
        func (l *List) InsertAfter(v interface{}, mark *Element) *Element
        func (l *List) InsertBefore(v interface{}, mark *Element) *Element
        func (l *List) Len() int
        func (l *List) MoveAfter(e, mark *Element)
        func (l *List) MoveBefore(e, mark *Element)
        func (l *List) MoveToBack(e *Element)
        func (l *List) MoveToFront(e *Element)
        func (l *List) PushBack(v interface{}) *Element
        func (l *List) PushBackList(other *List)
        func (l *List) PushFront(v interface{}) *Element
        func (l *List) PushFrontList(other *List)
        func (l *List) Remove(e *Element) interface{}

*/

package searchablelist

import "container/list"

// SearchableList list.List
type SearchableList struct {
	*list.List
}

// New () *list.List
func New() *SearchableList {
	return &SearchableList{new(list.List).Init()}
}

// (l *SearchableList)

// ContainsElement (t *list.Element) bool
func (l *SearchableList) ContainsElement(t *list.Element) bool {
	if l.Len() > 0 {
		for e := l.Front(); e != nil; e = e.Next() {
			if e == t {
				return true
			}
		}
	}
	return false
}

// Contains (t *list.Element) bool
// alias -> ContainsElement
func (l *SearchableList) Contains(t *list.Element) bool {
	return l.ContainsElement(t)
}

// ContainsValue (v interface{}) bool
func (l *SearchableList) ContainsValue(v interface{}) bool {
	if l.Len() > 0 {
		for e := l.Front(); e != nil; e = e.Next() {
			if e.Value == v {
				return true
			}
		}
	}
	return false
}

// FindFirst (v interface{}) *list.Element
func (l *SearchableList) FindFirst(v interface{}) *list.Element {
	if l.Len() > 0 {
		for e := l.Front(); e != nil; e = e.Next() {
			if e.Value == v {
				return e
			}
		}
	}
	return nil
}

// FindLast (v interface{}) *list.Element
func (l *SearchableList) FindLast(v interface{}) *list.Element {
	if l.Len() > 0 {
		for e := l.Back(); e != nil; e = e.Prev() {
			if e.Value == v {
				return e
			}
		}
	}
	return nil
}

// FindAll (v interface{}) []*list.Element
func (l *SearchableList) FindAll(v interface{}) []*list.Element {
	if l.Len() > 0 {
		elList := []*list.Element{}
		for e := l.Front(); e != nil; e = e.Next() {
			if e.Value == v {
				elList = append(elList, e)
			}
		}
		return elList
	}
	return nil
}
