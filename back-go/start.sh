killall back-go
. /etc/profile
go build
sleep 0.5
mv back.log back.log.bak
nohup ./back-go > back.log 2>&1 &
