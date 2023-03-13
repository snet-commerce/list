package list

// Element is a single element of List.
// Each Element has a reference to a list it belongs to, so it is safe (still avoid doing that) to pass
// Element to other list, but no action will be taken
type Element[E any] struct {
	list  *List[E]
	next  *Element[E]
	prev  *Element[E]
	value E
}

// Next returns next Element or nil if next element is not present
func (e *Element[E]) Next() *Element[E] {
	return e.next
}

// Prev returns previous Element or nil if previous element is not present
func (e *Element[E]) Prev() *Element[E] {
	return e.prev
}

// Value returns value stored in current Element
func (e *Element[E]) Value() E {
	return e.value
}

// List returns List to which current Element belongs
func (e *Element[E]) List() *List[E] {
	return e.list
}

// List is a doubly linked list data structure
type List[E any] struct {
	head *Element[E]
	tail *Element[E]
	len  int
}

// New initialises new List with initial number of elements passed
func New[E any](values ...E) *List[E] {
	list := &List[E]{}
	for i := range values {
		list.PushBack(values[i])
	}
	return list
}

// Len returns current number of elements in List
func (l *List[E]) Len() int {
	return l.len
}

// Clear cleanups List
func (l *List[E]) Clear() {
	l.head = nil
	l.tail = nil
	l.len = 0
}

// Front returns first element of the list
func (l *List[E]) Front() *Element[E] {
	return l.head
}

// Back returns last element of the list
func (l *List[E]) Back() *Element[E] {
	return l.tail
}

// PushFront insert element in the head of List
func (l *List[E]) PushFront(val E) *Element[E] {
	elem := &Element[E]{value: val, list: l}

	if l.head == nil {
		l.head = elem
		l.tail = elem
	} else {
		elem.next = l.head
		l.head.prev = elem
		l.head = elem
	}

	l.len++

	return elem
}

// PushBack insert element in the tail of List
func (l *List[E]) PushBack(val E) *Element[E] {
	elem := &Element[E]{value: val, list: l}

	if l.head == nil {
		l.head = elem
		l.tail = elem
	} else {
		l.tail.next = elem
		elem.prev = l.tail
		l.tail = elem
	}

	l.len++

	return elem
}

func (l *List[E]) InsertBefore(val E, bfr *Element[E]) *Element[E] {
	if bfr == nil || bfr.list != l {
		return nil
	}

	elem := &Element[E]{value: val, list: l}
	if bfr.prev == nil {
		elem.next = bfr
		bfr.prev = elem
		l.head = elem
	} else {
		elem.prev = bfr.prev
		elem.next = bfr
		bfr.prev.next = elem
		bfr.prev = elem
	}
	l.len++

	return elem
}

func (l *List[E]) InsertAfter(val E, aft *Element[E]) *Element[E] {
	if aft == nil || aft.list != l {
		return nil
	}

	elem := &Element[E]{value: val, list: l}
	if aft.next == nil {
		elem.prev = aft
		aft.next = elem
		l.tail = elem
	} else {
		elem.prev = aft
		elem.next = aft.next
		aft.next.prev = elem
		aft.next = elem
	}
	l.len++

	return elem
}

func (l *List[E]) RemoveFront() (val E) {
	if l.head == nil {
		return val
	}

	val = l.head.Value()
	l.head = l.head.next
	if l.head != nil {
		l.head.prev = nil
	} else {
		l.tail = nil
	}
	l.len--

	return val
}

func (l *List[E]) RemoveBack() (val E) {
	if l.tail == nil {
		return val
	}

	val = l.tail.Value()
	l.tail = l.tail.prev
	if l.tail != nil {
		l.tail.next = nil
	} else {
		l.head = nil
	}
	l.len--

	return val
}

func (l *List[E]) Remove(elem *Element[E]) (val E) {
	if elem == nil || elem.list != l {
		return val
	}

	if elem.prev != nil {
		elem.prev.next = elem.next
	} else {
		l.head = elem.next
	}

	if elem.next != nil {
		elem.next.prev = elem.prev
	} else {
		l.tail = elem.prev
	}

	l.len--

	return elem.Value()
}

func (l *List[E]) MoveToFront(elem *Element[E]) {
	l.MoveBefore(elem, l.head)
}

func (l *List[E]) MoveToBack(elem *Element[E]) {
	l.MoveAfter(elem, l.tail)
}

func (l *List[E]) MoveBefore(elem, bfr *Element[E]) {
	if elem == nil || bfr == nil || elem.list != l || bfr.list != l || elem == bfr {
		return
	}

	elem.prev.next = elem.next
	if elem.next != nil {
		elem.next.prev = elem.prev
	} else {
		l.tail = elem.prev
	}

	elem.prev = bfr.prev
	elem.next = bfr
	if bfr.prev != nil {
		bfr.prev.next = elem
	} else {
		l.head = elem
	}
	bfr.prev = elem
}

func (l *List[E]) MoveAfter(elem, aft *Element[E]) {
	if elem == nil || aft == nil || elem.list != l || aft.list != l || elem == aft {
		return
	}

	elem.prev.next = elem.next
	if elem.next != nil {
		elem.next.prev = elem.prev
	} else {
		l.tail = elem.prev
	}

	elem.prev = aft
	elem.next = aft.next
	if aft.next != nil {
		aft.next.prev = elem
	} else {
		l.tail = elem
	}
	aft.next = elem
}

func (l *List[E]) PushListBack(other *List[E]) {
	for e := other.Front(); e != nil; e = e.Next() {
		l.PushBack(e.Value())
	}
}

func (l *List[E]) Find(matcher func(E) bool) *Element[E] {
	for e := l.Front(); e != nil; e = e.Next() {
		if matcher(e.Value()) {
			return e
		}
	}
	return nil
}

func (l *List[E]) Append(values ...E) {
	for i := range values {
		l.PushBack(values[i])
	}
}

func (l *List[E]) Slice() []E {
	sl := make([]E, 0, l.len)
	for e := l.Front(); e != nil; e = e.Next() {
		sl = append(sl, e.Value())
	}
	return sl
}
