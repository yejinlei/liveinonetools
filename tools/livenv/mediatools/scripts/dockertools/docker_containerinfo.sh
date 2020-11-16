CONTAINER=$1
PID=$(docker inspect -f {{.State.Pid}} $CONTAINER)
MOUNTPOINT=$(docker inspect -f {{.GraphDriver.Data.Mountpoint}} $CONTAINER)
echo 'pid:'$PID 'mountpoint:'$MOUNTPOINT
