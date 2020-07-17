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
	state, err := db.DB.Prepare("INSERT INTO Links(Title,Address) VALUES(?,?)")
	if err != nil {
		log.Fatal(err)
	}
	res, err := state.Exec(link.Title, link.Address)
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
