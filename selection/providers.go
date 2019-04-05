package selection

type YesNoAnswerProvider interface {
	GetAnswer() bool
}

type NumericalSelectionProvider interface {
	GetSelection(maxSelect int) int
}

type AnswerProvider interface {
	YesNoAnswerProvider
	NumericalSelectionProvider
}
