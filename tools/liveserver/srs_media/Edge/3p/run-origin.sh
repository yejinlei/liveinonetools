PWD=$(pwd)
sudo docker-compose -f docker-compose-origin.yml down
sudo docker-compose -f docker-compose-origin.yml up -d
