package mysql

type Config struct {
	User     string `default:"user"`
	Password string `default:"user"`
	Host     string `default:"localhost"`
	Port     string `default:"3306"`
	Database string `default:"voicemod"`
}
