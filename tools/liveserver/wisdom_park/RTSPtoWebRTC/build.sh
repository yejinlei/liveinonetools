git clone --recursive https://github.com/deepch/RTSPtoWebRTC/
cd RTSPtoWebRTC
go mod vendor
go build -mod=vendor
cd ..
