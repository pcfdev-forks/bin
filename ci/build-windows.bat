set PATH=C:\Go\bin;%PATH%

set GOPATH=%CD%\gopath;%CD%\concourse;%CD%\gopath\src\github.com\vito\houdini\deps
set PATH=%CD%\gopath\bin;%PATH%

go get github.com/jteeuwen/go-bindata

go build -o go-bindata.exe github.com/jteeuwen/go-bindata/go-bindata

md windows
.\go-bindata.exe -pkg bindata -o gopath\src\github.com\concourse\bin\bindata\bindata.go windows/...

go build -o .\binary-windows\concourse_windows_amd64.exe github.com/concourse/bin/cmd/concourse
