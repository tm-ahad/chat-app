package interfaces

type Set[T Model] interface {
	Unmarshal(string) []Model
}
