package controller

import (
	myvars "github.com/GabrielHernanQuinteros/demoArticulos/vars"
	mytools "github.com/GabrielHernanQuinteros/prueba/video"
)

func CrearRegistroSQL(registro myvars.EstrucReg) error {

	bd, err := mytools.ConectarDB(myvars.ConnectionString)

	if err != nil {
		return err
	}

	_, err = bd.Exec("INSERT INTO video_games (name, genre, year) VALUES (?, ?, ?)", registro.Name, registro.Genre, registro.Year) //Modificar

	return err

}

func BorrarRegistroSQL(id int64) error {

	bd, err := mytools.ConectarDB(myvars.ConnectionString)

	if err != nil {
		return err
	}

	_, err = bd.Exec("DELETE FROM video_games WHERE id = ?", id) //Modificar

	return err
}

func ModificarRegistroSQL(registro myvars.EstrucReg) error {

	bd, err := mytools.ConectarDB(myvars.ConnectionString)

	if err != nil {
		return err
	}

	_, err = bd.Exec("UPDATE video_games SET name = ?, genre = ?, year = ? WHERE id = ?", registro.Name, registro.Genre, registro.Year, registro.Id) //Modificar

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
	rows, err := bd.Query("SELECT id, name, genre, year FROM video_games") //Modificar

	if err != nil {
		return arrRegistros, err
	}

	// Iterate rows...
	for rows.Next() {
		// In each step, scan one row
		var registro myvars.EstrucReg

		err = rows.Scan(&registro.Id, &registro.Name, &registro.Genre, &registro.Year) //Modificar

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

	row := bd.QueryRow("SELECT id, name, genre, year FROM video_games WHERE id = ?", id) //Modificar

	err = row.Scan(&registro.Id, &registro.Name, &registro.Genre, &registro.Year) //Modificar

	if err != nil {
		return registro, err
	}

	// Success!
	return registro, nil

}
