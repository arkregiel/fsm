# Finite State Machines

Prosta implementacja programowa automatów skończonych

## Źródła inspiracji

- [Automata & Python - Computerphile](https://www.youtube.com/watch?v=32bC33nJR3A)
- [Non-Deterministic Automata - Computerphile](https://www.youtube.com/watch?v=NhWDVqR4tZc)

## Deterministic finite automaton

[*Deterministyczny automat skończony*](https://en.wikipedia.org/wiki/Deterministic_finite_automaton) reprezentuje piątka: `DFA = {Q, Σ, δ, q0, F}`, gdzie:

- `Q` - skończony zbiór stanów automatu
- `Σ` - skończony zbiór symboli wejściowych (alfabet)
- `δ` - funkcja przejść (`δ: Q × Σ -> Q`)
- `q0` - stan początkowy (`q0 ∈ Q`)
- `F` - zbiór stanów końcowych (`F ⊆ Q`)

## Nondeterministic finite automaton

[*Niedeterministyczny automat skończony*](https://en.wikipedia.org/wiki/Nondeterministic_finite_automaton) reprezentuje piątka: `NFA = {Q, Σ, δ, S, F}`, gdzie:

- `Q` - skończony zbiór stanów automatu
- `Σ` - skończony zbiór symboli wejściowych (alfabet)
- `δ` - funkcja przejść (`δ: Q × Σ -> P(Q)`, `P` - [*powerset*](https://en.wikipedia.org/wiki/Power_set))
- `S` - zbiór stanów początkowych (`S ⊆ Q`)
- `F` - zbiór stanów końcowych (`F ⊆ Q`)
