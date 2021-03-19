CREATE TABLE IF NOT EXISTS users (
    aux_id        INT NOT NULL AUTO_INCREMENT,
	id            CHAR(36) NOT NULL,
	name          VARCHAR(100) NOT NULL,
	surnames      VARCHAR(100) NOT NULL,
	email      VARCHAR(100) NOT NULL,
	`password`      VARCHAR(100) NOT NULL,
	country      VARCHAR(100),
	phone      VARCHAR(100),
	postal_code      VARCHAR(100),

	PRIMARY KEY (aux_id),
	UNIQUE KEY idx_unique_id (id),
	UNIQUE KEY idx_unique_aux_id (aux_id),
	UNIQUE KEY idx_unique_email (email)

) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;
