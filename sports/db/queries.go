package db

const (
	sportsList = "list"
)

func getEventsQueries() map[string]string {
	return map[string]string{
		sportsList: `
			SELECT 
				id, 
				name, 
				stadium, 
				opponent, 
				advertised_start_time,
				status 
			FROM sports
		`,
	}
}
