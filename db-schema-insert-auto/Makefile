CONTAINER_NAME=temp-postgres
DB_NAME=testbulk

# public
## run app
run: db schema
	go run .

## varify result
select:
	docker run -it --rm --network bridge postgres psql -h $(shell docker inspect -f '{{range.NetworkSettings.Networks}}{{.IPAddress}}{{end}}' temp-postgres) -U postgres -W $(DB_NAME) -c "select * from names order by nconst limit 5;"

## psql cli
psql:
	docker run -it --rm --network bridge postgres psql -h $(shell docker inspect -f '{{range.NetworkSettings.Networks}}{{.IPAddress}}{{end}}' temp-postgres) -U postgres


#-------------------
# private 
notice:
	echo -e '遇到docker執行權限問題的話可能要設置 docker group\n\n sudo groupadd docker\n sudo usermod -aG docker $$USER \n newgrp docker \n'
db:
	docker run --name $(CONTAINER_NAME) -p 5433:5432 -e PGPASSWORD=12345 -e POSTGRES_PASSWORD=12345 -e POSTGRES_DB=$(DB_NAME) -d postgres

schema:
	docker run -it --rm --network bridge postgres \
	psql -h $(shell docker inspect -f '{{range.NetworkSettings.Networks}}{{.IPAddress}}{{end}}' temp-postgres) -U postgres -W $(DB_NAME) -c "CREATE TABLE names (nconst varchar(255), primary_name varchar(255), birth_year varchar(4), death_year varchar(4) DEFAULT '', primary_professions varchar[], known_for_titles varchar[]);"

rmdb:
	docker stop $(CONTAINER_NAME) 
	docker container rm $(CONTAINER_NAME) 

.PHONY: db psql schema rmdb run select 
.SILENT: db psql schema rmdb run select