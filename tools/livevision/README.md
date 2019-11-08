## 下载gocv（go get 或 go mod download）

## 下载并编译对应opencv
make deps
make download
make builds

## 编译gocv
go install gocv.io/x/gocv

## 设置编译环境
export CGO_CPPFLAGS="-I/usr/local/include"
export CGO_LDFLAGS="-L/usr/local/lib -lopencv_core -lopencv_face -lopencv_videoio -lopencv_imgproc -lopencv_highgui -lopencv_imgcodecs -lopencv_objdetect -lopencv_features2d -lopencv_video -lopencv_dnn -lopencv_xfeatures2d"
