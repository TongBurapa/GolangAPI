package Database

func Insert(users UsersArr) {

	db, err := Connect()
	defer db.Close()

	sqlStatement := `INSERT INTO userinfo  
					(ID, 
					name,
					lastname, 
					role)
					VALUES (
						$1, 
						$2, 
						$3, 
						$4)`

	for _, data := range users {
		_, err = db.Exec(sqlStatement, data.ID, data.Name, data.LastName, data.Role)
	}

	if err != nil {
		panic(err)
	}
}

func Delete(id UserDelete) {
	db, err := Connect()
	defer db.Close()

	sqlStatement := `DELETE FROM userinfo WHERE ID = $1`

	for _, data := range id.ID {
		_, err = db.Exec(sqlStatement, data)
	}

	if err != nil {
		panic(err)
	}
}

func Update(usersUpdate UsersArr) {
	db, err := Connect()
	defer db.Close()

	sqlStatement := `UPDATE userinfo  
					SET lastname = $2,
					name = $3,
					role = $4
					WHERE ID = $1`

	for _, data := range usersUpdate {
		_, err = db.Exec(sqlStatement, data.ID, data.LastName, data.Name, data.Role)
	}

	if err != nil {
		panic(err)
	}
}

func GetAllUsers() (result []Users, err error) {
	db, err := Connect()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	sqlStatement := `SELECT ID, lastname, name, role FROM userinfo`

	rows, err := db.Query(sqlStatement)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var usersGet Users
		err := rows.Scan(&usersGet.ID, &usersGet.Name, &usersGet.LastName, &usersGet.Role)
		if err != nil {
			return nil, err
		}
		result = append(result, usersGet)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return result, nil
}
