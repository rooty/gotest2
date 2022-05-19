curl --request POST \
  --url http://localhost:5000/v1/hello \
  --header 'Content-Type: application/json' \
  --data '{
	"firstname":"aaa",
	"lastname":"bbb"
}'
