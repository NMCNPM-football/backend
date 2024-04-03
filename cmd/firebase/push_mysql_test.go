package firebase

import (
	"database/sql"
	"fmt"
	"github.com/NMCNPM-football/backend/cmd/firebase/mysql"
	"github.com/NMCNPM-football/backend/internal/models"
	"log"
	"testing"
)

func TestPushData(t *testing.T) {
	log.Println("Testing Push Data")
	list := pushMySQLData()

	if len(list) > 0 {
		log.Println(fmt.Sprint("#v", list[0]))
		log.Println("Data is pushed successfully")
	} else {
		log.Println("Data is not pushed successfully")
	}
}

func pushMySQLData() []models.User {
	newMySQL := mysql.NewGoMithMysql("SE104", "localhost", "root", "root", "mysql")
	newMySQL.SetConnectionString("root:root@tcp(localhost:3306)/SE104?charset=utf8&parseTime=True&charset=utf8mb4&collation=utf8mb4_unicode_ci", 10)

	db := newMySQL.GetDBConnection()
	results, err := db.Query("SELECT * FROM users")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	defer results.Close()

	var list []models.User

	newUserData := models.User{}
	var createdAt, updatedAt sql.NullTime
	var phone sql.NullString
	for results.Next() {
		err := results.Scan(
			&newUserData.ID,
			//&newUserData.CreatedAt,
			//&newUserData.UpdatedAt,
			&createdAt,
			&updatedAt,
			&newUserData.DeletedAt,
			&newUserData.Name,
			&newUserData.Email,
			//&newUserData.Phone,
			&phone,
			&newUserData.Club,
			&newUserData.Address,
			&newUserData.Password,
			&newUserData.ClubID,
			&newUserData.Position,
			&newUserData.IsVerifiedEmail,
			&newUserData.SeaSon,
			&newUserData.UpdatedBy,
			&newUserData.DeletedBy)
		if err != nil {
			log.Fatalf("Error in fetching data %s \n", err.Error())
		}
	}
	list = append(list, newUserData)
	return list
}
