go env -w GO111MODULE=on
go env -w GOPROXY=https://goproxy.cn,direct
go get -v github.com/nsf/gocode
go get -v github.com/rogpeppe/godef
go get -v golang.org/x/tools/cmd/godoc
go get -v golang.org/x/tools/cmd/goimports
go get -v github.com/golang/lint/golint
go get -v github.com/derekparker/delve/cmd/dlv
go get -v github.com/lukehoban/go-find-references
go get -v github.com/lukehoban/go-outline
go get -v sourcegraph.com/sqs/goreturns
go get -v github.com/tpng/gopkgs
go get -v github.com/newhook/go-symbols
go get -v golang.org/x/tools/cmd/gorename
go get -u -v github.com/cweill/gotests/...
go get -u -v github.com/fatih/gomodifytags
go get -u -v github.com/josharian/impl
go get -u -v github.com/davidrjenni/reftools/cmd/fillstruct
go get -u -v github.com/haya14busa/goplay/cmd/goplay
go get -u -v github.com/godoctor/godoctor
