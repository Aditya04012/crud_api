package repository

import (
	"database/sql"
	"golang-server/golang_api/models"
)

func GetUsers(db *sql.DB) ([]models.User,error){

	//fetching users from db &&  sending error if exist
	rows,err:=db.Query(
		"SELECT id,name,email from users",
	)
 

	if err!=nil{
		return nil,err
	}
	//closing reading of db

	defer rows.Close()


	//creating list of Users
	users:=[]models.User{}


	//taking input in u and appending in list
	for rows.Next(){
		var u models.User
		err=rows.Scan(&u.Id,&u.Name,&u.Email)
		if err!=nil{
			return nil,err
		}
		users=append(users, u)
	}
 
	//returning
	return users,rows.Err()
}

/*
db.Query() returns *sql.Rows, and row.Scan() is not valid for scanning a single row directly.
Since you are fetching a user by ID, use db.QueryRow().


| Operation          | Method      | Purpose          |
|--------------------|-------------|------------------|
| Fetch one row      | `QueryRow()`| Get one user by ID |
| Fetch multiple rows| `Query()`   | Get all users    |
| Insert a row       | `Exec()`    | Create a user    |
| Update a row       | `Exec()`    | Edit a user      |
| Delete a row       | `Exec()`    | Delete a user    |

*/


//geting user by id 
func GetUserById(db *sql.DB,id int) (models.User,error){
	var u models.User
	row:=db.QueryRow(
		"SELECT id,name,email from users WHERE id=$1",
		id,
	)
	err:=row.Scan(&u.Id,&u.Name,&u.Email)

	if err!=nil{
		return models.User{},err
	}
  return u,nil
}

func DeleteUserById(db *sql.DB,id int)(string,error){
	_,err:=db.Exec(
		"DELETE FROM users WHERE id=$1",id,
	)
	if err!=nil{
		return "User doest exist",err
	}

	return "Deleted succesfully",nil
}
func Create(db *sql.DB,u *models.User) error{

    err := db.QueryRow("INSERT INTO users (name, email) VALUES ($1, $2) RETURNING id",
        u.Name,
        u.Email,
    ).Scan(&u.Id)

    return err
}