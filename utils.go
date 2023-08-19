package main

import "strconv"

func fnStringToInt64(parCadena string) (int64, error) {

	auxNumero, err := strconv.ParseInt(parCadena, 0, 64)

	if err != nil {
		return 0, err
	}

	return auxNumero, err
}
