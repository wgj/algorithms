package algorithms

import "testing"

func TestStackStruct(t *testing.T) {
	s := new(stack)
	if s == nil {
		t.Errorf("new(stack) == %T, want %T", *s, stack{})
	}
	if len(s.items) != 0 {
		t.Errorf("len(s.items) == %d, want %d", len(s.items), 0)
	}
	if s.index != 0 {
		t.Errorf("s.index == %d, want %d", s.index, 0)
	}
}

func TestPush(t *testing.T) {
	s := new(stack)
	for i := 0; i < 3; i++ {
		s.push(i)
		got := s.items[i]
		if got != i {
			t.Errorf("push(%d) == %d, want %d", i, got, i)
		}
	}
}

func TestIndex(t *testing.T) {
	s := new(stack)
	for i := 0; i < 3; i++ { /* 3 is the first iteration when s.items is larger that s.index */
		s.push(i)
	}
	if s.index != 3 {
		t.Errorf("s.index == %d, want %d", s.index, 3)
	}
}

func TestPop(t *testing.T) {
	s := new(stack)
	want := 2
	indexWant := 2
	for i := 0; i < 3; i++ {
		s.push(i)
	}
	got, err := s.pop()
	indexGot := s.index
	if err != nil {
		t.Errorf("s.pop(): %s", err)
	}
	if want != got {
		t.Errorf("s.pop() == %d, want %d", got, want)
	}
	if indexWant != indexGot {
		t.Errorf("s.index == %d, want %d", indexGot, indexWant)
	}
}

func TestPopEmpty(t *testing.T) {
	s := new(stack)
	i, err := s.pop()
	if err == nil {
		t.Errorf("s.pop() == %d, %s, want error", i, err)
	}
}
