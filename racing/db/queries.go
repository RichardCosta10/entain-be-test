package db

const (
	racesList = "list"
	racesGetById = "get"
)

func getRaceQueries() map[string]string {
	return map[string]string{
		racesList: `
			SELECT 
				id, 
				meeting_id, 
				name, 
				number, 
				visible, 
				advertised_start_time,
				status 
			FROM races
		`,
		racesGetById: `
			SELECT
				id,
				meeting_id,
				name,
				number,
				visible,
				advertised_start_time,
				status
				FROM races WHERE id = ?
			`,
	}
}
