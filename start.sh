echo "Starting up the database"
docker compose up db -d
until docker exec wagers-db sh -c 'export MYSQL_PWD=mysql; mysql -u root -e ";"'
do
    echo "Waiting for wagers-mysql database connection...\n"
    sleep 4
done
echo "Database is up and running"

echo "Running migrations"
docker compose --profile tools run --rm migrate up
echo "Migrations complete"

echo "Starting up the application"
docker compose up api --build -d
echo "Application is up and running"
