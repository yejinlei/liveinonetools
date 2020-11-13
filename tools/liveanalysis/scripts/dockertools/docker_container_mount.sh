CONTAINER=$1
PID=$(docker inspect -f {{.State.Pid}} $CONTAINER)
MOUNTPOINT=$(docker inspect -f {{.GraphDriver.Data.Mountpoint}} $CONTAINER)
mkdir ./mnt
mount -o bind $MOUNTPOINT ./mnt
