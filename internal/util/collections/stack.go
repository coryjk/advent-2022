package collections

import (
	"container/list"
	"strings"
	"sync"
)

type Stack struct {
	dll   *list.List
	mutex sync.Mutex
}

func NewStack() *Stack {
	return &Stack{dll: list.New()}
}

func (s *Stack) Push(x interface{}) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	s.dll.PushBack(x)
}

func (s *Stack) Pop() interface{} {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	if s.dll.Len() == 0 {
		return nil
	}
	tail := s.dll.Back()
	val := tail.Value
	s.dll.Remove(tail)
	return val
}

func (s *Stack) Peek() interface{} {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	if s.dll.Len() == 0 {
		return nil
	}
	return s.dll.Back().Value
}

func (s *Stack) Len() int {
	return s.dll.Len()
}

func (s *Stack) String(toString func(x interface{}) string) string {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	builder := strings.Builder{}
	for i := s.dll.Front(); i != nil; i = i.Next() {
		val := toString(i.Value)
		builder.WriteRune('[')
		builder.WriteString(val)
		builder.WriteRune(']')
	}
	return builder.String()
}
