NETWORK=bridge
HOST="172.17.0.7"
DB_USER=postgres
DB_PASSWORD=12345
DB=grpost


docker run -it --rm --network $NETWORK postgres psql -h $HOST -U $DB_USER -W$"$DB_PASSWORD" $DB -e "$1"