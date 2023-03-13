package list_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/snet-commerce/list"
)

func TestList_Init(t *testing.T) {
	want := []int{4, 2, 1, 12, 101, 88}
	lst := list.New(want...)

	if err := compareListAndSlice(lst, want); err != nil {
		t.Fatal(err)
	}
}

func TestList_PushBack(t *testing.T) {
	want := []string{"str1", "str2", "str3", "str4"}
	lst := list.New(want...)

	if err := compareListAndSlice(lst, want); err != nil {
		t.Fatal(err)
	}

	pushBackValues := []string{"str5", "str6", "str7"}
	for _, pushVal := range pushBackValues {
		elem := lst.PushBack(pushVal)
		if val := elem.Value(); val != pushVal {
			t.Fatalf("new element must be equal %s, but got %s", pushVal, val)
		}
	}

	if backElem, lastPushElem := lst.Back(), pushBackValues[len(pushBackValues)-1]; backElem == nil || backElem.Value() != lastPushElem {
		t.Fatalf("new element is %s, but list back elemet is %s", lastPushElem, backElem.Value())
	}

	want = append(want, pushBackValues...)
	if err := compareListAndSlice(lst, want); err != nil {
		t.Fatal(err)
	}
}

func TestList_PushFront(t *testing.T) {
	want := []string{"str1", "str2", "str3", "str4"}
	lst := list.New(want...)

	if err := compareListAndSlice(lst, want); err != nil {
		t.Fatal(err)
	}

	pushFrontValues := []string{"str5", "str6", "str7"} // will be inserted in reverse order, since every push is performed in head
	for _, pushVal := range pushFrontValues {
		elem := lst.PushFront(pushVal)
		if val := elem.Value(); val != pushVal {
			t.Fatalf("new element must be equal %s, but got %s", pushVal, val)
		}
	}

	if frontElem, firstPushElem := lst.Front(), pushFrontValues[len(pushFrontValues)-1]; frontElem == nil || frontElem.Value() != firstPushElem {
		t.Fatalf("new element is %s, but list front elemet is %s", firstPushElem, frontElem.Value())
	}

	for _, val := range pushFrontValues {
		want = append([]string{val}, want...)
	}
	if err := compareListAndSlice(lst, want); err != nil {
		t.Fatal(err)
	}
}

func TestList_Iteration(t *testing.T) {
	want := []int{4, 2, 1, 12, 101, 88}
	lst := list.New(want...)

	i := 0
	for e := lst.Front(); e != nil; e = e.Next() {
		if lstVal, slValue := e.Value(), want[i]; lstVal != slValue {
			t.Fatalf("values on forward iteration (index %d) are not equal, list has element %d, slice %d", i, lstVal, slValue)
		}
		i++
	}

	i = len(want) - 1
	for e := lst.Back(); e != nil; e = e.Prev() {
		if lstVal, slValue := e.Value(), want[i]; lstVal != slValue {
			t.Fatalf("values on backward iteration (index %d) are not equal, list has element %d, slice %d", i, lstVal, slValue)
		}
		i--
	}
}

func TestList_InsertBefore(t *testing.T) {
	lst := list.New[int]()
	lst.PushBack(22)
	lst.PushBack(33)
	lst.PushFront(11)
	lst.PushFront(0)

	frontElem := lst.Front()
	lst.InsertBefore(-11, frontElem)
	lst.InsertBefore(-5, frontElem)

	frontElem = lst.Front()
	lst.InsertBefore(-22, frontElem)

	backElem := lst.Back()
	newElem := lst.InsertBefore(27, backElem)
	lst.InsertBefore(24, newElem) // insert right before newly created element

	want := []int{-22, -11, -5, 0, 11, 22, 24, 27, 33}
	if err := compareListAndSlice(lst, want); err != nil {
		t.Fatal(err)
	}
}

func TestList_InsertAfter(t *testing.T) {
	lst := list.New[int]()
	lst.PushBack(20)
	lst.PushBack(30)
	lst.PushBack(40)
	lst.PushFront(10)
	lst.PushFront(0)

	frontElem := lst.Front()
	newElem := lst.InsertAfter(5, frontElem)
	newElem = lst.InsertAfter(7, newElem)
	lst.InsertAfter(9, newElem)

	backElem := lst.Back()
	backElem = lst.InsertAfter(50, backElem)
	lst.InsertAfter(60, backElem)

	want := []int{0, 5, 7, 9, 10, 20, 30, 40, 50, 60}
	if err := compareListAndSlice(lst, want); err != nil {
		t.Fatal(err)
	}
}

func TestList_Remove(t *testing.T) {
	sl := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	lst := list.New(sl...)

	if rmVal, slVal := lst.RemoveFront(), sl[0]; rmVal != slVal {
		t.Fatalf("removed value %d is not equal to front value %d", rmVal, slVal)
	}

	if rmVal, slVal := lst.RemoveBack(), sl[len(sl)-1]; rmVal != slVal {
		t.Fatalf("removed value %d is not equal to back value %d", rmVal, slVal)
	}

	rmElem := lst.Front()
	for i := 0; i < 3; i++ {
		rmElem = rmElem.Next()
	}

	// 4 = 1 + 3 -> first one is removed via RemoveFront and other 3 are iterated before
	if rmVal, slVal := lst.Remove(rmElem), sl[4]; rmVal != slVal {
		t.Fatalf("removed value %d is not equal to value %d at index %d", rmVal, slVal, 4)
	}
}

func TestList_Move(t *testing.T) {
	sl := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	lst := list.New(sl...)

	backElem := lst.Back()
	lst.MoveToFront(backElem) // 9, 1, 2, 3, 4, 5, 6, 7, 8

	backElem = lst.Back()
	lst.MoveToFront(backElem) // 8, 9, 1, 2, 3, 4, 5, 6, 7

	elem5 := lst.Back().Prev().Prev() // element with 5
	lst.MoveToBack(elem5)             // 8, 9, 1, 2, 3, 4, 6, 7, 5

	elem6 := lst.Back().Prev().Prev()  // element with 6
	lst.MoveBefore(elem6, lst.Front()) // 6, 8, 9, 1, 2, 3, 4, 7, 5

	elem4 := lst.Back().Prev().Prev() // element with 4
	lst.MoveAfter(elem4, lst.Back())  // 6, 8, 9, 1, 2, 3, 7, 5, 4

	elem1 := lst.Front().Next().Next().Next() // element with 1
	elem3 := lst.Back().Prev().Prev().Prev()  // element with 3
	lst.MoveAfter(elem1, elem3)               // 6, 8, 9, 2, 3, 1, 7, 5, 4

	elem2 := lst.Front().Next().Next().Next() // element with 2
	elem7 := lst.Back().Prev().Prev()         // element with 7
	lst.MoveBefore(elem2, elem7)              // 6, 8, 9, 3, 1, 2, 7, 5, 4

	want := []int{6, 8, 9, 3, 1, 2, 7, 5, 4}
	if err := compareListAndSlice(lst, want); err != nil {
		t.Fatal(err)
	}
}

func TestList_Find(t *testing.T) {
	type number struct {
		n int
	}

	lst := list.New(
		number{n: 44},
		number{n: 88},
		number{n: 11},
		number{n: -7},
		number{n: 8},
		number{n: -12},
	)

	elem := lst.Find(func(numb number) bool {
		return numb.n < 0
	})
	if actual, expected := elem.Value().n, -7; actual != expected { // first negative in list
		t.Fatalf("first struct with negative value in list must have value %d, but got %d", expected, actual)
	}
}

func compareListAndSlice[E any](lst *list.List[E], sl []E) error {
	if lstLen, slLen := lst.Len(), len(sl); lstLen != slLen {
		return fmt.Errorf("expected slice has %d elements, but list has %d", slLen, lstLen)
	}

	if lstSlice := lst.Slice(); !reflect.DeepEqual(lstSlice, sl) {
		return fmt.Errorf("expected slice and slice from list are not equal, expected %v != actual %v", sl, lstSlice)
	}

	return nil
}
