if [ -n "$1" ] ; then
   echo "kill process contains $1"
      ps -aux |grep $1|grep -v grep|awk '{print "USER: "$1"\nCMD : "$11" "$12"\nKILL: "$2"\n"}'
      ps -aux |grep $1|grep -v grep|awk '{print $2}'|xargs kill -9
else
   echo "Usage: /killall ..."
fi
exit
