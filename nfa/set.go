package nfa

import (
	"github.com/arkregiel/fsm/fsa"
)

// StateSet reprezentuje zbiór stanów
type StateSet struct {
	States map[fsa.State]bool
}

// AddState dodaje stan do zbioru
func (set *StateSet) AddState(q fsa.State) *StateSet {
  set.States[q] = true
	return set
}

// ContainsFinalState sprawdza, czy w zbiorze występuje stan końcowy
func (set StateSet) ContainsFinalState() bool {
	for q := range set.States {
		if q.IsFinal {
			return true
		}
	}
	return false
}

// Extend rozszerza istniejący zbiór stanów
func (set *StateSet) Extend(stateSet StateSet) *StateSet {
	for q := range stateSet.States {
		set.AddState(q)
	}
	return set
}

// NewEmptyStateSet zwraca pusty zbiór stanów
func NewEmptyStateSet() *StateSet {
	return &StateSet{
		States: make(map[fsa.State]bool),
	}
}

// NewStateSet zwraca zbiór stanów na podstawie podanej tablicy stanów
func NewStateSet(states []fsa.State) *StateSet {
	newSet := NewEmptyStateSet()
	for _, q := range states {
		newSet.AddState(q)
	}
	return newSet
}
