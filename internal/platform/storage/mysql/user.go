package mysql

const sqlUserTable = "users"

type sqlUser struct {
	ID         string `db:"id"`
	Name       string `db:"name"`
	Surnames   string `db:"surnames"`
	Email      string `db:"email"`
	Password   string `db:"password"`
	Country    string `db:"country"`
	Phone      string `db:"phone"`
	PostalCode string `db:"postal_code"`
}

var sqlInsertUserColumns = []string{
	"id",
	"name",
	"surnames",
	"email",
	"password",
	"country",
	"phone",
	"postal_code",
}
