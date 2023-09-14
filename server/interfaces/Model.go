package interfaces

type Model interface {
	Marshal() 			string
	Unique()  			string
	Unmarshal(string)   
}