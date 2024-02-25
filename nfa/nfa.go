package nfa

import (
	"fmt"
	"slices"

	"github.com/arkregiel/fsm/fsa"
)

// NFA to struktura reprezentująca niedeterministyczny automat skończony
type NFA struct {
	fsa.FiniteAutomaton                              // {Q, Σ, F}
	Transitions         map[fsa.Transition]*StateSet // Funkcja przejść - δ: Q × Σ -> P(Q)
	InitialStates       *StateSet                    // Zbiór stanów początkowych S
}

// Accepts sprawdza czy automat akceptuje dane wyrażenie
func (nfa NFA) Accepts(word string) (bool, error) {
	currentStatesSet := nfa.InitialStates
	for _, symbol := range word {
		if !slices.Contains(nfa.Alphabet, symbol) {
			return false, fmt.Errorf("syntax error: %c not in alphabet", symbol)
		}

		currentStatesSet = nfa.makeTransitions(currentStatesSet, symbol)
	}
	return currentStatesSet.ContainsFinalState(), nil
}

// makeTransitions zwraca następne stany, do których ze stanów obecnych
// przejdzie automat pod wpływem podanego symbolu
func (nfa NFA) makeTransitions(currentStates *StateSet, symbol rune) *StateSet {
	newStatesSet := NewEmptyStateSet()
	for _, q := range currentStates.States {
		if nextStates, ok := nfa.Transitions[fsa.Transition{State: q, Symbol: symbol}]; ok {
			newStatesSet.Extend(*nextStates)
		}
	}
	return newStatesSet
}
