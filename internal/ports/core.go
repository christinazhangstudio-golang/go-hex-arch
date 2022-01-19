package ports

//adapters for core layer

// interface core adapter needs to implement
// in order for adapter to implement this, it must implement the methods
type ArithmeticPort interface {
	Addition(a int32, b int32) (int32, error)
	Subtraction(a int32, b int32) (int32, error)
	Multiplication(a int32, b int32) (int32, error)
	Division(a int32, b int32) (int32, error)
}