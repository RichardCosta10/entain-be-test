package db

import (
	"time"

	"syreclabs.com/go/faker"
)

func (e *sportsRepo) seed() error {
	statement, err := e.db.Prepare(`CREATE TABLE IF NOT EXISTS sports (id INTEGER PRIMARY KEY, name TEXT, stadium TEXT, opponent TEXT, advertised_start_time DATETIME, status TEXT, created_at DATETIME DEFAULT CURRENT_TIMESTAMP, updated_at DATETIME DEFAULT CURRENT_TIMESTAMP)`)
	if err == nil {
		_, err = statement.Exec()
	}

	for i := 1; i <= 100; i++ {
		startTime := faker.Time().Between(time.Now().AddDate(0, 0, -1), time.Now().AddDate(0, 0, 2))
		statement, err = e.db.Prepare(`INSERT OR IGNORE INTO sports(id, name, stadium, opponent, advertised_start_time, status) VALUES (?,?,?,?,?,?)`)
		if err == nil {
			_, err = statement.Exec(
				i,
				faker.Team().Name(),
				faker.Address().City() + " Stadium",
				faker.Team().Name(),
				startTime.Format(time.RFC3339),
				func(startTime time.Time) string {
					if startTime.Before(time.Now()) {
						return "played"
					} else {
						return "upcoming"
					}
				}(startTime),
			)
		}
	}

	return err
}
