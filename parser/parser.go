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

	// comprueba si el largo de la cadena es igual a los 4 primeros caracteres + el largo indicado por length
	if len(s) != intLength+4 {
		return nil, errors.New("El valor length indicado no coincide con el largo del value")
	}

	switch s[0:2] {
	case "TX":
		//si se especifica tipo TX y dentro del value se encuentra un numero se retorna un error
		// el cero equivale al ascii 48 y el nueve al ascii 57
		for _, c := range s[4:] {
			if c >= 48 && c <= 57 {
				return nil, errors.New("Se especifico tipo TX pero el value es numerico")
			}
		}

	case "NN":
		//si se especifica tipo NN y dentro del value se encuentra un caracter diferente a un numero se retorna un error
		// el cero equivale al ascii 48 y el nueve al ascii 57
		for _, c := range s[4:] {
			if c < 48 || c > 57 {
				return nil, errors.New("Se especifico tipo NN pero el value no es numerico")
			}
		}
	default:
		return nil, errors.New("El tipo debe ser TX o NN")
	}

	resultado := &Result{s[0:2], s[4:], intLength}
	return resultado, nil
}
