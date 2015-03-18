package syncstack

import (
  "sync"
)

// Stack is the main type. Holds a pointer to the top element,
// the size of the stack and the mutex.
type Stack struct {
  top   *element
  size  int
  mutex *sync.Mutex
}

type element struct {
  value interface{}
  next  *element
}

// Len returns the stack's size.
func (s *Stack) Len() int {
  s.mutex.Lock()
  defer s.mutex.Unlock()
  return s.size
}

// Push puts a new value on top of the stack.
func (s *Stack) Push(value interface{}) {
  s.mutex.Lock()
  defer s.mutex.Unlock()
  s.top = &element{value, s.top}
  s.size++
}

// Pop removes the top element from the stack and returns its value.
// If the stack is empty, return nil.
func (s *Stack) Pop() (value interface{}) {
  s.mutex.Lock()
  defer s.mutex.Unlock()
  if s.size > 0 {
    value, s.top = s.top.value, s.top.next
    s.size--
    return
  }
  return nil
}

// Peek returns the top value without removing it.
func (s *Stack) Peek() (value interface{}) {
  return s.top.value
}

// NewStack returns a pointer to a new initialized stack
func NewStack() *Stack {
  return &stack{
    top:   nil,
    size:  0,
    mutex: new(sync.Mutex),
  }
}
