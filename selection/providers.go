package selection

// YesNoAnswerProvider provides true for yes and false for no
type YesNoAnswerProvider interface {
	GetAnswer() bool
}

// NumericalSelectionProvider provides a selected number within a given range
type NumericalSelectionProvider interface {
	GetSelection(maxSelect int) int
}

// AnswerProvider can provide a yes/no answer or a selected number
type AnswerProvider interface {
	YesNoAnswerProvider
	NumericalSelectionProvider
}
