package main

import (
	"fmt"
)

func main() {
	const name = "Ian"
	const age = 37

	fmt.Printf("%s is %d years old and the type of variables are %T and %T", name, age, name, age)
	fmt.Println(`
		Compiling this landmark anthology of poetry in English
		about dogs and musical instruments is like swimming through bricks.
		To date, I have only, “On the Death of Mrs. McTuesday’s Pug,
		Killed by a Falling Piano,” a somewhat obvious choice.
		True, an Aeolian harp whispers alluringly
		in the background of the anonymous sonnet, “The Huntsman’s Hound,”
		but beyond that — silence.
		
		I should resist this degrading donkey-work in favor of my own writing,
		wherein contentment surely lies.
		But A. Smith stares smugly from the reverse of the twenty pound note,
		and when my bank manager guffaws, small particles of saliva stream like a meteor shower
		through the infinity of dark space
		between his world and mine.
	`)
}
