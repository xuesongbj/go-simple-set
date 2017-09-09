/*
Open Source Initiative OSI - The MIT License (MIT):Licensing
The MIT License (MIT)
Copyright (c) 2013 Ralph Caraveo (deckarep@gmail.com)
Permission is hereby granted, free of charge, to any person obtaining a copy of
this software and associated documentation files (the "Software"), to deal in
the Software without restriction, including without limitation the rights to
use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies
of the Software, and to permit persons to whom the Software is furnished to do
so, subject to the following conditions:
The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.
THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/
package simpleSet

// Package golang-simple-set implements a simple and generic set collection.
import (
	"sync"
)

type SafeSet struct {
	m UnsafeSet
	sync.RWMutex
}

func newSafeSet() SafeSet {
	return SafeSet{m: newUnsafeSet()}
}

func (s *SafeSet) Add(item string) {
	s.Lock()
	defer s.Unlock()

	s.m[item] = struct{}{}
}

func (s *SafeSet) Remove(item string) {
	s.Lock()
	defer s.Unlock()
	delete(s.m, item)
}

func (s *SafeSet) Has(item string) bool {
	s.RLock()
	defer s.RUnlock()

	_, ok := s.m[item]
	return ok
}

func (s *SafeSet) Len() int {
	return len(s.List())
}

func (s *SafeSet) Clear() {
	s.Lock()
	defer s.Unlock()
	s.m = map[string]struct{}{}
}

func (s *SafeSet) Isempty() bool {
	if s.Len() == 0 {
		return true
	}
	return false
}

func (s *SafeSet) List() []string {
	s.RLock()
	defer s.RUnlock()

	list := make([]string, 0, 1000)
	for i := range s.m {
		list = append(list, i)
	}
	return list
}
