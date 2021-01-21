git clone --recursive https://github.com/deepch/RTSPtoWebRTC/
cd RTSPtoWebRTC
go mod vendor
#CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -mod=vendor
go build -mod=vendor
cd ..
