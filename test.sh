source ./infra/.env
docker-compose -f ./infra/docker-compose.yaml up --build -d
echo "Waiting for multichain to boot..."
sleep 300
go test -v ./...
docker-compose -f ./infra/docker-compose.yaml down
echo "Done!"