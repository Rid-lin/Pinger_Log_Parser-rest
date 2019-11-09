CREATE TABLE "users" (
	"id"	INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
	"email"	varchar NOT NULL UNIQUE,
	"encrypted_password"	varchar NOT NULL
)