package pkg_test

import (
	"reflect"
	"testing"

	"github.com/conorkenn/data_structures_go/pkg/tree.go"
)

func TestInsert(t *testing.T) {
	myTree := tree.NewTreeNode(4)
	myTree.Insert(3)
	myTree.Insert(5)

	if myTree.GetValue() != 4 {
		t.Errorf("got %d want %d", myTree.GetValue(), 4)
	}
}

func TestInorderTraversal(t *testing.T) {
	expected := []int{3, 4, 5}
	myTree := tree.NewTreeNode(4)
	myTree.Insert(3)
	myTree.Insert(5)

	actual := myTree.InorderTraversal()
	equal := reflect.DeepEqual(expected, actual)
	if !equal {
		t.Errorf("expected %v and actual %v are different", expected, actual)
	}
}

func TestPostorderTraversal(t *testing.T) {
	expected := []int{3, 5, 4}
	myTree := tree.NewTreeNode(4)
	myTree.Insert(3)
	myTree.Insert(5)

	actual := myTree.PostorderTraversal()
	equal := reflect.DeepEqual(expected, actual)
	if !equal {
		t.Errorf("expected %v and actual %v are different", expected, actual)
	}
}

func TestPreorderTraversal(t *testing.T) {
	expected := []int{4, 3, 5}
	myTree := tree.NewTreeNode(4)
	myTree.Insert(3)
	myTree.Insert(5)

	actual := myTree.PreorderTraversal()
	equal := reflect.DeepEqual(expected, actual)
	if !equal {
		t.Errorf("expected %v and actual %v are different", expected, actual)
	}
}

func TestSearch(t *testing.T) {
	myTree := tree.NewTreeNode(4)
	myTree.Insert(3)
	myTree.Insert(5)

	if myTree.Search(3) == false {
		t.Errorf("Search failed")
	}

}
