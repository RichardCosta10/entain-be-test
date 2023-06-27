package db

import (
	"time"

	"syreclabs.com/go/faker"
)

func (r *racesRepo) seed() error {
	r.db.Exec(`DROP TABLE IF EXISTS races`)
	statement, err := r.db.Prepare(`CREATE TABLE IF NOT EXISTS races (id INTEGER PRIMARY KEY, meeting_id INTEGER, name TEXT, number INTEGER, visible INTEGER, advertised_start_time DATETIME, status TEXT)`)
	if err == nil {
		_, err = statement.Exec()
	}

	for i := 1; i <= 100; i++ {
		startTime := faker.Time().Between(time.Now().AddDate(0, 0, -1), time.Now().AddDate(0, 0, 2))
		statement, err = r.db.Prepare(`INSERT OR IGNORE INTO races(id, meeting_id, name, number, visible, advertised_start_time, status) VALUES (?,?,?,?,?,?,?)`)
		if err == nil {
			_, err = statement.Exec(
				i,
				faker.Number().Between(1, 10),
				faker.Team().Name(),
				faker.Number().Between(1, 12),
				faker.Number().Between(0, 1),
				startTime.Format(time.RFC3339),
				func(startTime time.Time) string {
					if startTime.Before(time.Now()) {
						return "closed"
					} else {
						return "open"
					}
				}(startTime),
			)
		}
	}

	return err
}
