package auth

type Config struct {
	TimeToExpireDays int    `default:"1"`
	SignKey          string `default:"super_secret_key"`
}
