CONTAINER=$1
PID=$(docker inspect -f {{.State.Pid}} $CONTAINER)
nsenter --target $PID --uts --ipc --net --pid /bin/sh
