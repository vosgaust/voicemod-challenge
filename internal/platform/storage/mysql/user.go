package mysql

const sqlUserTable = "users"

type sqlUser struct {
	ID         string `db:"id" structs:"id"`
	Name       string `db:"name" structs:"name"`
	Surnames   string `db:"surnames" structs:"surnames"`
	Email      string `db:"email" structs:"email"`
	Password   string `db:"password" structs:"password"`
	Country    string `db:"country" structs:"country"`
	Phone      string `db:"phone" structs:"phone"`
	PostalCode string `db:"postal_code" structs:"postal_code"`
}

var sqlUserColumns = []string{
	"id",
	"name",
	"surnames",
	"email",
	"password",
	"country",
	"phone",
	"postal_code",
}

var sqlUpdatableColumns = []string{
	"name",
	"surnames",
	"email",
	"password",
	"country",
	"phone",
	"postal_code",
}
