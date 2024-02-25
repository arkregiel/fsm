package main

import (
	"fmt"

	"github.com/arkregiel/fsm/dfa"
	"github.com/arkregiel/fsm/fsa"
	"github.com/arkregiel/fsm/nfa"
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

	fmt.Println("DFA")
	fmt.Println(d0.Accepts(""))      // true
	fmt.Println(d0.Accepts("aa"))    // true
	fmt.Println(d0.Accepts("aabbb")) // true
	fmt.Println(d0.Accepts("ba"))    // false
	fmt.Println(d0.Accepts("aabba")) // false
	fmt.Println()

	// Niederministyczny automat skończony, który akceptuje wyrażenia,
	// w których na przedostatnim miejscu występuje a
	// Σ = {"a", "b"}
	// δ:
	//  q0: a -> {q0, q1}, b -> q0 | START
	//  q1: a, b -> q2
	//  q2: {} | FINAL
	nfaStates := []fsa.State{
		{Number: 0, IsFinal: false},
		{Number: 1, IsFinal: false},
		{Number: 2, IsFinal: true},
	}

	nfaTransitions := map[fsa.Transition]*nfa.StateSet{
		{State: nfaStates[0], Symbol: alphabet[0]}: nfa.NewStateSet([]fsa.State{nfaStates[0], nfaStates[1]}),
		{State: nfaStates[0], Symbol: alphabet[1]}: nfa.NewStateSet([]fsa.State{nfaStates[0]}),
		{State: nfaStates[1], Symbol: alphabet[0]}: nfa.NewStateSet([]fsa.State{nfaStates[2]}),
		{State: nfaStates[1], Symbol: alphabet[1]}: nfa.NewStateSet([]fsa.State{nfaStates[2]}),
	}

	n0 := nfa.NFA{
		FiniteAutomaton: fsa.FiniteAutomaton{
			States:   nfaStates,
			Alphabet: alphabet,
		},
		Transitions:   nfaTransitions,
		InitialStates: nfa.NewStateSet([]fsa.State{nfaStates[0]}),
	}

	fmt.Println("NFA")
	fmt.Println(n0.Accepts(""))       // false
	fmt.Println(n0.Accepts("a"))      // false
	fmt.Println(n0.Accepts("aa"))     // true
	fmt.Println(n0.Accepts("aabb"))   // false
	fmt.Println(n0.Accepts("aab"))    // true
	fmt.Println(n0.Accepts("bab"))    // true
	fmt.Println(n0.Accepts("aabbaa")) // true
}
