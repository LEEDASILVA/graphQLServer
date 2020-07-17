package links

import (
	"log"
	"os/user"

	db "github.com/LEEDASILVA/grapQLServer/go/internal/pkg/db/mysql"
)

type Link struct {
	ID      string
	Title   string
	Address string
	User    *user.User
}

func (l *Link) Save() int64 {
	stmt, err := db.DB.Prepare("INSERT INTO Links(Title,Address) VALUES(?,?)")
	if err != nil {
		log.Fatal(err)
	}
	res, err := stmt.Exec(l.Title, l.Address)
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
	stmt, err := db.DB.Prepare("SELECT id, title, address FROM Links")
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

	for rows.Next() {
		var link Link
		err := rows.Scan(&link.ID, &link.Title, &link.Address)
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
