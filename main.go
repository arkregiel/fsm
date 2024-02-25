package main

import (
	"fmt"

	"github.com/arkregiel/fsm/dfa"
	"github.com/arkregiel/fsm/fsa"
)

func main() {
	// deterministyczny automat skończony, który akceptuje wyrażenia,
	// w których symbol 'a' nie może występować później niż symbol 'b'
	// Σ = {"a", "b"}
	// δ:
	//  q0: a -> q0, b -> q1 | START FINAL
	//  q1: a -> q2, b -> q1 | FINAL
	//  q2: a, b -> q2

	alphabet := []rune{'a', 'b'}
	states := []fsa.State{
		{Number: 0, IsFinal: true},
		{Number: 1, IsFinal: true},
		{Number: 2, IsFinal: false},
	}

	// tabela przejść
	transitions := map[fsa.Transition]fsa.State{
		{State: states[0], Symbol: alphabet[0]}: states[0],
		{State: states[0], Symbol: alphabet[1]}: states[1],
		{State: states[1], Symbol: alphabet[0]}: states[2],
		{State: states[1], Symbol: alphabet[1]}: states[1],
		{State: states[2], Symbol: alphabet[0]}: states[2],
		{State: states[2], Symbol: alphabet[1]}: states[2],
	}

	d0 := dfa.DFA{
		FiniteAutomaton: fsa.FiniteAutomaton{
			States:   states,
			Alphabet: alphabet,
		},
		Transitions:  transitions,
		InitialState: states[0],
	}

	fmt.Println(d0.Accepts(""))      // true
	fmt.Println(d0.Accepts("aa"))    // true
	fmt.Println(d0.Accepts("aabbb")) // true
	fmt.Println(d0.Accepts("ba"))    // false
	fmt.Println(d0.Accepts("aabba")) // false
}
