package main

import "strconv"

func StringToInt64(parCadena string) (int64, error) {

	auxNumero, err := strconv.ParseInt(parCadena, 0, 64)

	if err != nil {
		return 0, err
	}

	return auxNumero, err
}
