package api

import (
	"github.tesla.com/chrzhang/go-hex-arch/internal/ports"
)

type Adapter struct {
	db ports.DbPort
	arith ports.ArithmeticPort
}

func NewAdapter(db ports.DbPort, arith ports.ArithmeticPort) *Adapter {
	return &Adapter{db: db, arith: arith}
}

//API Adapter can access this method that is part of the Arithmetic port
//because there is access up in the struct, NewAdapter creates the object and
//we can use dependency injection
//API Adapter implements API Adapter port interface
func (apia Adapter) GetAddition(a, b int32) (int32, error) {
	answer, err := apia.arith.Addition(a, b)
	if err != nil {
		return 0, err
	}

	err = apia.db.AddToHistory(answer, "addition")
	if err != nil {
		return 0, err
	}

	return answer, nil
}
func (apia Adapter) GetSubtraction(a, b int32) (int32, error) {
	answer, err := apia.arith.Subtraction(a, b)
	if err != nil {
		return 0, err
	}

	err = apia.db.AddToHistory(answer, "subtraction")
	if err != nil {
		return 0, err
	}
	return answer, nil
}
func (apia Adapter) GetMultiplication(a, b int32) (int32, error) {
	answer, err := apia.arith.Multiplication(a, b)
	if err != nil {
		return 0, err
	}
	err = apia.db.AddToHistory(answer, "multiplication")
	if err != nil {
		return 0, err
	}

	return answer, nil
}
func (apia Adapter) GetDivision(a, b int32) (int32, error) {
	answer, err := apia.arith.Division(a, b)
	if err != nil {
		return 0, err
	}
	err = apia.db.AddToHistory(answer, "division")
	if err != nil {
		return 0, err
	}

	return answer, nil
}