package links

import (
	"log"

	"github.com/LEEDASILVA/graphQLServer/go/internal/users"

	db "github.com/LEEDASILVA/graphQLServer/go/internal/pkg/db/mysql"
)

type Link struct {
	ID      string
	Title   string
	Address string
	User    *users.User
}

func (l *Link) Save() int64 {
	stmt, err := db.DB.Prepare("INSERT INTO Links(Title,Address,UserID) VALUES(?,?,?)")
	if err != nil {
		log.Fatal(err)
	}
	res, err := stmt.Exec(l.Title, l.Address, l.User.ID)
	if err != nil {
		log.Fatal(err)
	}
	id, err := res.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Row *links* inserted!")
	return id
}

func GetAll() []Link {
	stmt, err := db.DB.Prepare("select L.id, L.title, L.address, L.UserID, U.Username from Links L inner join Users U on L.UserID = U.ID")
	if err != nil {
		log.Panic(err)
	}
	defer stmt.Close()
	rows, err := stmt.Query()
	if err != nil {
		log.Panic(err)
	}
	defer rows.Close()

	var ls []Link
	var username string
	var id string
	for rows.Next() {
		var link Link
		err := rows.Scan(&link.ID, &link.Title, &link.Address, &id, &username)
		if err != nil {
			log.Fatal(err)
		}
		ls = append(ls, link)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
	return ls
}
