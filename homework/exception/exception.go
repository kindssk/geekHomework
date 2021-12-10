package exception

import (
	"fmt"
	mysqlx "geekHomework/initialize"
	_ "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
)

type user struct {
	Id   int
	Name string
	Age  int
}

func FindUserById(id int) (*user, error) {
	u := new(user)
	row := mysqlx.MysqlDB.QueryRow("select * from users where id=?", id)
	err := row.Scan(&u.Id, &u.Name, &u.Age)
	if err != nil {
		return nil, errors.WithMessage(err, fmt.Sprintf("select * from users where id = %d", id))
	}
	return u, nil
}
