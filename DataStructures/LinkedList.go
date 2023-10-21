package DataStructures

import (
	"errors"
	"fmt"
)

type Node struct {
	Val  int
	Next *Node
}
type LinkedList struct {
	Len  int
	Head *Node
}

func Print(list *LinkedList) {
	temp := list.Head
	for temp != nil {
		fmt.Print(temp.Val)
		fmt.Print("->")
		temp = temp.Next
	}
	fmt.Println()
}

// Создания нового односвязного списка с количеством узлов q
func NewList(q int) LinkedList {
	if q > 0 {
		list := LinkedList{Len: q, Head: newNode(0)}
		temp := list.Head
		for i := 1; i < q; i++ {
			temp.Next = newNode(0)
			temp = temp.Next
		}
		return list
	}
	return LinkedList{Len: 0, Head: nil}
}

// создание нового узла
func newNode(k int) *Node {
	node := Node{Val: k, Next: nil}
	return &node
}

// добавление в конец списка
func (list *LinkedList) Add(k int) {
	temp := list.Head
	// если список пустой
	if temp == nil {
		node := newNode(k)
		list.Head = node
	} else { // иначе
		for temp.Next != nil {
			temp = temp.Next
		}
		node := newNode(k)
		temp.Next = node
	}
	list.Len++

}

// удаление с конца списка
func (list *LinkedList) Pop() error {
	// если список пустой
	if list.Len < 1 {
		return errors.New("Nothing to delete")
	}

	if list.Len > 1 {
		temp := list.Head
		var temp2 *Node
		for temp.Next != nil {
			temp2 = temp
			temp = temp.Next
		}
		temp2.Next = nil
		list.Len--

	} else { // если в списке только один элемент
		list.Head = nil
		list.Len = 0
	}
	return nil
}

// Показ значения на позиции pos (позиции начинаются с 1)
func (list *LinkedList) At(pos int) (int, error) {
	if pos > list.Len || pos < 1 {
		return 0, errors.New("Position out of list")
	} else {
		i := 1
		temp := list.Head
		for i != pos {
			temp = temp.Next
			i++
		}
		return temp.Val, nil
	}
}

// Размер списка
func (list *LinkedList) Size() int {
	return list.Len
}

// Удаление из спискм узла на позиции pos (позиции начинаются с 1)
func (list *LinkedList) DeleteFrom(pos int) error {
	if pos > list.Len || pos < 1 {
		return errors.New("Position out of list")
	} else if pos == 1 { // если надо удалить первый узел
		list.Head = list.Head.Next
	} else if pos == list.Len { // если последний
		temp := list.Head
		for temp.Next.Next != nil {
			temp = temp.Next
		}
		temp.Next = nil
	} else {
		i := 1
		temp := list.Head
		for temp.Next != nil && i+1 != pos {
			temp = temp.Next
			i++
		}
		temp.Next = temp.Next.Next

	}
	list.Len--
	return nil
}

// Изменение значения в узле на позиции pos (позиции начинаются с 1)
func (list *LinkedList) UpdateAt(pos, val int) error {
	if pos > list.Len || pos < 1 {
		return errors.New("Position out of list")
	} else {
		i := 1
		temp := list.Head
		for i != pos {
			temp = temp.Next
			i++
		}
		temp.Val = val
		return nil
	}
}

func FromSliceToList(s []int) LinkedList {
	list := LinkedList{Len: len(s), Head: newNode(s[0])}
	temp := list.Head
	for i := 1; i < len(s); i++ {
		temp.Next = newNode(s[i])
		temp = temp.Next
	}
	return list
}

func (list *LinkedList) InsertAt(pos, val int) error {
	if pos > list.Len+1 || pos < 1 {
		return errors.New("Position out of list")
	}
	node := newNode(val)
	if pos == 1 {
		node.Next = list.Head
		list.Head = node
	} else if pos == (list.Len + 1) {
		temp := list.Head
		for temp.Next != nil {
			temp = temp.Next
		}
		temp.Next = node
	} else {
		i := 1
		temp := list.Head
		for temp.Next != nil && i+1 != pos {
			temp = temp.Next
			i++
		}
		node.Next = temp.Next
		temp.Next = node
	}
	list.Len++
	return nil
}
