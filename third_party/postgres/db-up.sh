docker-compose up -d
docker exec -it ggs-db psql -U postgres -a -f docker-entrypoint-initdb.d/init.sql