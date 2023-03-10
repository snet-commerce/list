package list

type Element[E any] struct {
	list  *List[E]
	next  *Element[E]
	prev  *Element[E]
	value E
}

func (e *Element[E]) Next() *Element[E] {
	return e.next
}

func (e *Element[E]) Prev() *Element[E] {
	return e.prev
}

func (e *Element[E]) Value() E {
	return e.value
}

type List[E any] struct {
	head *Element[E]
	tail *Element[E]
	len  int
}

func New[E any]() *List[E] {
	return &List[E]{}
}

func (l *List[E]) Len() int {
	return l.len
}

func (l *List[E]) Clear() {
	l.head = nil
	l.tail = nil
	l.len = 0
}

func (l *List[E]) Front() *Element[E] {
	return l.head
}

func (l *List[E]) Back() *Element[E] {
	return l.tail
}

func (l *List[E]) PushFront(val E) *Element[E] {
	elem := l.newElement(val)

	if l.head == nil {
		l.head = elem
		l.tail = elem
	} else {
		elem.next = l.head
		l.head.prev = elem
		l.head = elem
	}

	l.incrementLen()

	return elem
}

func (l *List[E]) PushBack(val E) *Element[E] {
	elem := l.newElement(val)

	if l.head == nil {
		l.head = elem
		l.tail = elem
	} else {
		l.tail.next = elem
		elem.prev = l.tail
		l.tail = elem
	}

	l.incrementLen()

	return elem
}

func (l *List[E]) InsertBefore(val E, bfr *Element[E]) *Element[E] {
	if !l.elemBelongToList(bfr) {
		return nil
	}

	elem := l.newElement(val)
	elem.next = bfr
	elem.prev = bfr.prev

	bfr.prev = elem
	l.incrementLen()

	return elem
}

func (l *List[E]) InsertAfter(val E, aft *Element[E]) *Element[E] {
	if !l.elemBelongToList(aft) {
		return nil
	}

	elem := l.newElement(val)
	elem.next = aft.next
	elem.prev = aft

	aft.next = elem
	l.incrementLen()

	return elem
}

func (l *List[E]) RemoveFront() *Element[E] {
	if l.head == nil {
		return nil
	}

	rm := l.head
	l.head = l.head.next
	if l.head != nil {
		l.head.prev = nil
	}
	rm.next = nil
	l.decrementLen()

	return rm
}

func (l *List[E]) RemoveBack() *Element[E] {
	if l.tail == nil {
		return nil
	}

	rm := l.tail
	l.tail = l.tail.prev
	if l.tail != nil {
		l.tail.next = nil
	}
	rm.prev = nil
	l.decrementLen()

	return rm
}

func (l *List[E]) Remove(elem *Element[E]) *Element[E] {
	if !l.elemBelongToList(elem) {
		return nil
	}

	// TODO
	l.decrementLen()
}

func (l *List[E]) newElement(val E) *Element[E] {
	return &Element[E]{
		value: val,
		list:  l,
	}
}

func (l *List[E]) incrementLen() {
	l.len++
}

func (l *List[E]) decrementLen() {
	l.len--
}

func (l *List[E]) elemBelongToList(elem *Element[E]) bool {
	return elem.list == l
}
