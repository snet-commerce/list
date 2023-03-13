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

func compareListAndSlice[E any](lst *list.List[E], sl []E) error {
	if lstLen, slLen := lst.Len(), len(sl); lstLen != slLen {
		return fmt.Errorf("expected slice has %d elements, but list has %d", slLen, lstLen)
	}

	if lstSlice := lst.Slice(); !reflect.DeepEqual(lstSlice, sl) {
		return fmt.Errorf("expected slice and slice from list are not equal, expected %v != actual %v", sl, lstSlice)
	}

	return nil
}
