PWD=$(pwd)
sudo docker-compose -f docker-compose-up.yml down
sudo docker-compose -f docker-compose-down.yml down
sudo docker-compose -f docker-compose-up.yml up -d
sudo docker-compose -f docker-compose-down.yml up -d

