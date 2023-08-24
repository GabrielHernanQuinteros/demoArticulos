package controller

import (
	myvars "github.com/GabrielHernanQuinteros/demoArticulos/vars"
	mytools "github.com/GabrielHernanQuinteros/demoCommon"
)

func CrearRegistroSQL(registro myvars.EstrucReg) error {

	bd, err := mytools.ConectarDB(myvars.ConnectionString)

	if err != nil {
		return err
	}

	_, err = bd.Exec("INSERT INTO articulos (codigo, nombre, precio, stock) VALUES (?, ?, ?, ?)", registro.Codigo, registro.Nombre, registro.Precio, registro.Stock) //Modificar

	return err

}

func BorrarRegistroSQL(id int64) error {

	bd, err := mytools.ConectarDB(myvars.ConnectionString)

	if err != nil {
		return err
	}

	_, err = bd.Exec("DELETE FROM articulos WHERE id = ?", id) //Modificar

	return err
}

func ModificarRegistroSQL(registro myvars.EstrucReg) error {

	bd, err := mytools.ConectarDB(myvars.ConnectionString)

	if err != nil {
		return err
	}

	_, err = bd.Exec("UPDATE articulos SET codigo = ?, nombre = ?, precio = ?, stock = ? WHERE id = ?", registro.Codigo, registro.Nombre, registro.Precio, registro.Stock, registro.Id) //Modificar

	return err
}

func TraerRegistrosSQL() ([]myvars.EstrucReg, error) {

	//Declare an array because if there's error, we return it empty
	arrRegistros := []myvars.EstrucReg{}

	bd, err := mytools.ConectarDB(myvars.ConnectionString)

	if err != nil {
		return arrRegistros, err
	}

	// Get rows so we can iterate them
	rows, err := bd.Query("SELECT * FROM articulos") //Modificar

	if err != nil {
		return arrRegistros, err
	}

	// Iterate rows...
	for rows.Next() {
		// In each step, scan one row
		var registro myvars.EstrucReg

		err = rows.Scan(&registro.Id, &registro.Codigo, &registro.Nombre, &registro.Stock, &registro.Precio) //Modificar

		if err != nil {
			return arrRegistros, err
		}

		// and append it to the array
		arrRegistros = append(arrRegistros, registro)
	}

	return arrRegistros, nil

}

func TraerRegistroPorIdSQL(id int64) (myvars.EstrucReg, error) {

	var registro myvars.EstrucReg

	bd, err := mytools.ConectarDB(myvars.ConnectionString)

	if err != nil {
		return registro, err
	}

	row := bd.QueryRow("SELECT * FROM articulos WHERE id = ?", id) //Modificar

	err = row.Scan(&registro.Id, &registro.Codigo, &registro.Nombre, &registro.Stock, &registro.Precio) //Modificar

	if err != nil {
		return registro, err
	}

	// Success!
	return registro, nil

}

func TraerRegistroPorNombreSQL(parNombre string) (myvars.EstrucReg, error) {

	var registro myvars.EstrucReg

	bd, err := mytools.ConectarDB(myvars.ConnectionString)

	if err != nil {
		return registro, err
	}

	row := bd.QueryRow("SELECT * FROM articulos WHERE nombre = ?", parNombre) //Modificar

	err = row.Scan(&registro.Id, &registro.Codigo, &registro.Nombre, &registro.Stock, &registro.Precio) //Modificar

	if err != nil {
		return registro, err
	}

	// Success!
	return registro, nil

}
