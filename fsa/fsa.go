package fsa

// State to struktura reprezentująca stan automatu
type State struct {
	Number  uint32 // numer stanu
	IsFinal bool   // czy stan jest stanem końcowym
}

// Transition to struktura reprezentująca przejście automatu
type Transition struct {
	State  State // stan
	Symbol rune  // symbol, dla którego ma nastąpić przejście
}

// FiniteAutomaton reprezentuje automat skończony.
// Struktura nie zawiera tabeli przejść i stanu bądź zbioru stanów
// początkowych. Te mogą różnić się w zależności od rodzaju automatu (DFA lub NFA)
type FiniteAutomaton struct {
	States      []State // Zbiór stanów - Q
	Alphabet    []rune  // Zbiór symboli wejściowych - Σ
	FinalStates []State // Zbiór stanów końcowych - F
}

type IFiniteStateMachine interface {
	// Metoda Accepts ma sprawdzać czy automat akceptuje dane wyrażenie
	// i zwracać błąd jeśli wystąpi problem z wyrażeniem bądź automatem
	Accepts(word string) (bool, error)
}
