package interfaces

type Model interface {
	Marshal() 			string
	Unique()  			any
	Unmarshal(string)   
}