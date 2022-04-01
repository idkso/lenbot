package database

func init() {
	DBEntries = append(DBEntries, &DBEntry{
		Name: "coins",
		Fields: []string{
			"id INTEGER NOT NULL PRIMARY KEY",
			"pocket INTEGER NOT NULL",
			"bank INTEGER NOT NULL",
			"banksize INTEGER NOT NULL",
			"lastdaily INTEGER NOT NULL",
		},
	})
}
