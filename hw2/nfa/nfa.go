package nfa

//import "fmt"
// A state in the NFA is labeled by a single integer.
type state uint

// TransitionFunction tells us, given a current state and some symbol, which
// other states the NFA can move to.
//
// Deterministic automata have only one possible destination state,
// but we're working with non-deterministic automata.
type TransitionFunction func(st state, act rune) []state

func Reachable(
	// `transitions` tells us what our NFA looks like
	transitions TransitionFunction,
	// `start` and `final` tell us where to start, and where we want to end up
	start, final state,
	// `input` is a (possible empty) list of symbols to apply.
	input []rune,
) bool {
	// TODO Write the Reachable function,
	// return true if the nfa accepts the input,
	// return true if it doesn't reach to the final state with that input

	//Recursive Backtracking function to solve nfa
	if(len(input) == 0 && start == final){
		return true //found the solution! :)
	}

	if(len(input)<=0){
		return false //everytime it returns false, then backtrack
	}

	var x []state = transitions(start,input[0])

	if(len(x) == 0){
		return false
	}

	for i:=0;i<len(x);i++{
		start = x[i]
		if(Reachable(transitions,start,final,input[1:])){
			return true
		}
	}
	return false
}





func Reachable(transitions TransitionFunction, start, final state, input[] rune,) bool{
	if len input == 0{
		return start == final
	}
	for _, next := range transitions(start,input[0]){
		if Reachable(transitions,next,final,input[1:]){
			return true
		}
	}
	return false
}