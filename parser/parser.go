package parser

import (
	"errors"
	"strconv"
)

type Result struct {
	Type   string
	Value  string
	Length int
}

func GetEstructura(s string) (*Result, error) {
	if len(s) < 5 {
		return nil, errors.New("entrada invalida, se deben ingresar 5 caracteres como minimo")
	}
	intLength, err := strconv.Atoi(s[2:4])
	if err != nil {
		return nil, errors.New("Se espera que el valor de length sea numerico!")
	}

	resultado := &Result{s[0:2], s[4:], intLength}
	return resultado, nil
}
