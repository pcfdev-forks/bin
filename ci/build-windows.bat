set PATH=C:\Go\bin;%PATH%

set GOPATH=%CD%\gopath;%CD%\concourse;%CD%\gopath\src\github.com\vito\houdini\deps
set PATH=%CD%\gopath\bin;%PATH%

go build -o .\binary\concourse_windows_amd64.exe github.com/concourse/bin/cmd/concourse
