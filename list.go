package slru

import (
	"container/list"
	"sync"
)

// List is thread safe list base on list.List
type List struct {
	ll *list.List
	sync.Mutex
}

func NewList() *List {
	return &List{
		ll: list.New(),
	}
}
func (l *List) Back() *list.Element {
	l.Lock()
	defer l.Unlock()
	return l.ll.Back()
}

func (l *List) Front() *list.Element {
	l.Lock()
	defer l.Unlock()
	return l.ll.Front()
}

func (l *List) Init() *List {
	l.Lock()
	defer l.Unlock()
	l.ll.Init()
	return l
}
func (l *List) InsertAfter(v interface{}, mark *list.Element) *list.Element {
	l.Lock()
	defer l.Unlock()
	return l.ll.InsertAfter(v, mark)
}
func (l *List) InsertBefore(v interface{}, mark *list.Element) *list.Element {
	l.Lock()
	defer l.Unlock()
	return l.ll.InsertBefore(v, mark)
}
func (l *List) Len() int {
	l.Lock()
	defer l.Unlock()
	return l.ll.Len()
}
func (l *List) MoveAfter(e, mark *list.Element) {
	l.Lock()
	defer l.Unlock()
	l.ll.MoveAfter(e, mark)
}
func (l *List) MoveBefore(e, mark *list.Element) {
	l.Lock()
	defer l.Unlock()
	l.ll.MoveBefore(e, mark)
}
func (l *List) MoveToBack(e *list.Element) {
	l.Lock()
	defer l.Unlock()
	l.ll.MoveToBack(e)
}
func (l *List) MoveToFront(e *list.Element) {
	l.Lock()
	defer l.Unlock()
	l.ll.MoveToFront(e)
}
func (l *List) PushBack(v interface{}) *list.Element {
	l.Lock()
	defer l.Unlock()
	return l.ll.PushBack(v)
}
func (l *List) PushBackList(other *List) {
	l.Lock()
	defer l.Unlock()
	l.ll.PushBackList(other.ll)
}
func (l *List) PushFront(v interface{}) *list.Element {
	l.Lock()
	defer l.Unlock()
	return l.ll.PushFront(v)
}
func (l *List) PushFrontList(other *List) {
	l.Lock()
	defer l.Unlock()
	l.ll.PushFrontList(other.ll)
}
func (l *List) Remove(e *list.Element) interface{} {
	l.Lock()
	defer l.Unlock()
	return l.ll.Remove(e)
}
