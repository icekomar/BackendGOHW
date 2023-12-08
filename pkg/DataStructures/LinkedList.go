package DataStructures

import (
	"errors"
	"fmt"
)

type node struct {
	val  int
	next *node
}
type LinkedList struct {
	len  int
	Head *node
}

func Print(list *LinkedList) {
	temp := list.Head
	for temp != nil {
		fmt.Print(temp.val)
		fmt.Print("->")
		temp = temp.next
	}
	fmt.Println()
}

// NewList Создания нового односвязного списка с количеством узлов q
func NewList(q int) LinkedList {
	if q > 0 {
		list := LinkedList{len: q, Head: newNode(0)}
		temp := list.Head
		for i := 1; i < q; i++ {
			temp.next = newNode(0)
			temp = temp.next
		}
		return list
	}
	return LinkedList{len: 0, Head: nil}
}

// создание нового узла
func newNode(k int) *node {
	newnode := node{val: k, next: nil}
	return &newnode
}

// Add добавление в конец списка
func (list *LinkedList) Add(k int) {
	temp := list.Head
	// если список пустой
	if temp == nil {
		newnode := newNode(k)
		list.Head = newnode
		return
	} // иначе
	for temp.next != nil {
		temp = temp.next
	}
	temp.next = newNode(k)
	list.len++

}

// Pop удаление с конца списка
func (list *LinkedList) Pop() error {
	// если список пустой
	if list.len < 1 {
		return errors.New("nothing to delete")
	}
	if list.len > 1 {
		temp := list.Head
		var temp2 *node
		for temp.next != nil {
			temp2 = temp
			temp = temp.next
		}
		temp2.next = nil
		list.len--
		return nil
	} // если в списке только один элемент
	list.Head = nil
	list.len = 0
	return nil
}

// At Показ значения на позиции pos (позиции начинаются с 1)
func (list *LinkedList) At(pos int) (int, error) {
	if pos > list.len || pos < 1 {
		return 0, errors.New("position out of list")
	}
	i := 1
	temp := list.Head
	for i != pos {
		temp = temp.next
		i++
	}
	return temp.val, nil

}

// Size Размер списка
func (list *LinkedList) Size() int {
	return list.len
}

// DeleteFrom Удаление из спискм узла на позиции pos (позиции начинаются с 1)
func (list *LinkedList) DeleteFrom(pos int) error {
	switch true {
	case pos > list.len || pos < 1:
		return errors.New("position out of list")
	case pos == 1: // если надо удалить первый узел
		list.Head = list.Head.next
	case pos == list.len: // если последний
		temp := list.Head
		for temp.next.next != nil {
			temp = temp.next
		}
		temp.next = nil
	default:
		i := 1
		temp := list.Head
		for temp.next != nil && i+1 != pos {
			temp = temp.next
			i++
		}
		temp.next = temp.next.next

	}
	list.len--
	return nil
}

// UpdateAt Изменение значения в узле на позиции pos (позиции начинаются с 1)
func (list *LinkedList) UpdateAt(pos, val int) error {
	if pos > list.len || pos < 1 {
		return errors.New("position out of list")
	}
	i := 1
	temp := list.Head
	for i != pos {
		temp = temp.next
		i++
	}
	temp.val = val
	return nil

}

func FromSliceToList(s []int) LinkedList {
	list := LinkedList{len: len(s), Head: newNode(s[0])}
	temp := list.Head
	for i := 1; i < len(s); i++ {
		temp.next = newNode(s[i])
		temp = temp.next
	}
	return list
}

func (list *LinkedList) InsertAt(pos, val int) error {

	switch {
	case pos > list.len+1 || pos < 1:
		return errors.New("position out of list")
	case pos == 1:
		newnode := newNode(val)
		newnode.next = list.Head
		list.Head = newnode
	case pos == (list.len + 1):
		newnode := newNode(val)
		temp := list.Head
		for temp.next != nil {
			temp = temp.next
		}
		temp.next = newnode
	default:
		newnode := newNode(val)
		i := 1
		temp := list.Head
		for temp.next != nil && i+1 != pos {
			temp = temp.next
			i++
		}
		newnode.next = temp.next
		temp.next = newnode
	}
	list.len++
	return nil
}
