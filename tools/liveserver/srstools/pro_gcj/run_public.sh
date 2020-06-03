sudo docker network create system-network;
docker-compose -f public_plug.yml down;
sleep 5;
docker-compose -f public_plug.yml up;
