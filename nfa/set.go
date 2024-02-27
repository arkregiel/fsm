package nfa

import (
	"slices"

	"github.com/arkregiel/fsm/fsa"
)

// StateSet reprezentuje zbiór stanów
type StateSet struct {
	States []fsa.State
}

// AddState dodaje stan do zbioru
func (set *StateSet) AddState(q fsa.State) *StateSet {
	if !slices.Contains(set.States, q) {
		set.States = append(set.States, q)
	}
	return set
}

// ContainsFinalState sprawdza, czy w zbiorze występuje stan końcowy
func (set StateSet) ContainsFinalState() bool {
	for _, q := range set.States {
		if q.IsFinal {
			return true
		}
	}
	return false
}

// Extend rozszerza istniejący zbiór stanów
func (set *StateSet) Extend(stateSet StateSet) *StateSet {
	for _, q := range stateSet.States {
		set.AddState(q)
	}
	return set
}

// NewEmptyStateSet zwraca pusty zbiór stanów
func NewEmptyStateSet() *StateSet {
	return &StateSet{
		States: make([]fsa.State, 0),
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
