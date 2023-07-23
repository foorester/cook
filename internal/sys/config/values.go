package config

var Key = newCfgKeyReg()

func newCfgKeyReg() *cfgKeyReg {
	return &cfgKeyReg{
		// API Server

		APIServerHost:     "http.api.server.host",
		APIServerPort:     "http.api.server.port",
		APIServerTimeout:  "http.api.server.shutdown.timeout.secs",
		APIErrorExposeInt: "api.errors.expose.internal",

		// Postgres

		SQLiteUser:     "db.sqlite.user",
		SQLitePass:     "db.sqlite.pass",
		SQLiteDB:       "db.sqlite.database",
		SQLiteHost:     "db.sqlite.host",
		SQLiteSchema:   "db.sqlite.schema",
		SQLiteFilePath: "db.sqlite.filepath",

		// Postgres

		PgUser:   "db.pg.user",
		PgPass:   "db.pg.pass",
		PgDB:     "db.pg.database",
		PgHost:   "db.pg.host",
		PgPort:   "db.pg.port",
		PgSchema: "db.pg.schema",
		PgSSL:    "db.pg.sslmode",

		// Mongo

		MongoUser: "db.mongo.user",
		MongoPass: "db.mongo.pass",
		MongoDB:   "db.mongo.database",
		MongoHost: "db.mongo.host",
		MongoPort: "db.mongo.port",
	}
}

type cfgKeyReg struct {
	APIServerHost     string
	APIServerPort     string
	APIServerTimeout  string
	APIErrorExposeInt string

	// SQLite

	SQLiteUser     string
	SQLitePass     string
	SQLiteDB       string
	SQLiteHost     string
	SQLitePort     string
	SQLiteSchema   string
	SQLiteSSL      string
	SQLiteFilePath string

	// Postgres

	PgUser   string
	PgPass   string
	PgDB     string
	PgHost   string
	PgPort   string
	PgSchema string
	PgSSL    string

	// Mongo

	MongoUser string
	MongoPass string
	MongoDB   string
	MongoHost string
	MongoPort string
}
