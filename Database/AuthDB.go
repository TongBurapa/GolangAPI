package Database

import "time"

func TokenDB(token string, expTime string) {
	db, err := Connect()
	defer db.Close()

	sqlStatement := `INSERT INTO token  
					(createdate,
					expiredate, 
					token)
					VALUES (
						$1, 
						$2, 
						$3)`

	_, err = db.Exec(sqlStatement, time.Now(), expTime, token)

	if err != nil {
		panic(err)
	}

}
