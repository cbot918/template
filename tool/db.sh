NETWORK=qcode-server_default
HOST=db
DB_USER=root
DB_PASSWORD=12345
DB=qcode

docker run -it --rm --network $NETWORK mysql mysql -h db -u $DB_USER -p$"$DB_PASSWORD" $DB -e "$1"