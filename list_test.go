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

	pushBackVal := "str5"
	elem := lst.PushBack(pushBackVal)
	if val := elem.Value(); val != pushBackVal {
		t.Fatalf("new element must be equal %s, but got %s", pushBackVal, val)
	}

	if backElem := lst.Back(); backElem == nil || backElem.Value() != pushBackVal {
		t.Fatalf("new element is %s, but list back elemet is %s", pushBackVal, backElem.Value())
	}

	want = append(want, pushBackVal)
	if err := compareListAndSlice(lst, want); err != nil {
		t.Fatal(err)
	}
}

func TestList_PushFront(t *testing.T) {

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
