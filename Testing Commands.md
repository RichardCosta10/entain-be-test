### Question 1

Added **visible_only** key of type bool to filter:

curl -X "POST" "http://localhost:8000/v1/list-races" \
     -H 'Content-Type: application/json' \
     -d $'{
  "filter": {"visible_only": true}
}'

### Question 2

Added **order_by** with keys `field` (string) and `desc` (bool). If desc is omitted it will sort asc

curl -X "POST" "http://localhost:8000/v1/list-races" \
     -H 'Content-Type: application/json' \
     -d $'{
  "filter": {"meeting_ids": [4,5,6], "visible_only": true},
  "order_by": {"field": "advertised_start_time", "desc": true}      
}'

### Question 3

Can see status in the output 

### Question 4

Get single race by id with below request: 

curl "http://localhost:8000/v1/races/{id}"

### Question 5

List sporting events, can filter by id

curl -X "POST" "http://localhost:8000/v1/list-sporting-events" \
     -H 'Content-Type: application/json' \
     -d $'{
  "filter": {}
}'