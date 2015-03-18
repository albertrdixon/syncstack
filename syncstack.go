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

// Return the stack's length
func (s *stack) Len() int {
  s.mutex.Lock()
  defer s.mutex.Unlock()
  return s.size
}

// Push a new element onto the stack
func (s *stack) Push(value interface{}) {
  s.mutex.Lock()
  defer s.mutex.Unlock()
  s.top = &element{value, s.top}
  s.size++
}

// Remove the top element from the stack and return it's value
// If the stack is empty, return nil
func (s *stack) Pop() (value interface{}) {
  s.mutex.Lock()
  defer s.mutex.Unlock()
  if s.size > 0 {
    value, s.top = s.top.value, s.top.next
    s.size--
    return
  }
  return nil
}

// NewStack returns a pointer to a new initialized stack
func NewStack() *stack {
  return &stack{
    top:   nil,
    size:  0,
    mutex: new(sync.Mutex),
  }
}
