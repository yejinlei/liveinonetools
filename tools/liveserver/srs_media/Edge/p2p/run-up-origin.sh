PWD=$(pwd)
sudo docker-compose -f docker-compose-up-origin.yml down
sudo docker-compose -f docker-compose-up-origin.yml up -d
