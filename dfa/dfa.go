package dfa

import (
	"fmt"
	"slices"

	"github.com/arkregiel/fsm/fsa"
)

// DFA to struktura reprezentująca deterministyczny automat skończony
type DFA struct {
	fsa.FiniteAutomaton                              // {Q, Σ, F}
	Transitions         map[fsa.Transition]fsa.State // Funkcja przejść - δ: Q × Σ -> Q
	InitialState        fsa.State                    // Stan początkowy - q0
}

// Accepts sprawdza czy automat akceptuje dane wyrażenie
func (dfa DFA) Accepts(word string) (bool, error) {
	// zaczynamy od stanu początkowego
	q := dfa.InitialState
	for _, symbol := range word {
		if !slices.Contains(dfa.Alphabet, symbol) {
			return false, fmt.Errorf("syntax error: %c not in alphabet", symbol)
		}

		if newQ, ok := dfa.Transitions[fsa.Transition{State: q, Symbol: symbol}]; ok {
			q = newQ
		} else {
			return false, fmt.Errorf("transition error: no transition for state %v with symbol '%c'", q, symbol)
		}

	}

	return q.IsFinal, nil
}
