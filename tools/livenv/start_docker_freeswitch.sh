sudo docker run --net=host -d -p 5060:5060 -p 5080:5080 -v /usr/src/freeswitch.git/conf:/usr/src/freeswitch/conf  bettervoice/freeswitch-container:1.6.6

