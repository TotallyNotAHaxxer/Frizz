package Engine

import Helpers "main/Modules/Gsrc/Helpers"

// Before loading we want to check all of the arrays for duplicate information
func Preloader(Vectocheck []string) {
	Vectocheck = Helpers.ValueRemover(Vectocheck)
}
