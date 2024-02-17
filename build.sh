version="1.0.0"
os_types=("darwin" "linux" "windows")
for os_type in ${os_types[@]};do
    echo "打${os_type}包"
    mkdir -p target/jwt-authorization-${version}
    cp auth.conf target/jwt-authorization-${version}
    CGO_ENABLED=0 GOOS=${os_type} GOARCH=amd64 go build -o target/jwt-authorization -ldflags "-s -w" .
    if [ "${os_type}" = "windows" ];then
        mv target/jwt-authorization target/jwt-authorization-${version}/jwt-authorization.exe
        cd target && zip -r jwt-authorization-${os_type}-${version}.zip jwt-authorization-${version}/* && cd ..
    else
        mv target/jwt-authorization target/jwt-authorization-${version}/jwt-authorization
        cd target && tar -zcvf jwt-authorization-${os_type}-${version}.tar.gz jwt-authorization-${version}/* && cd ..
    fi
    rm -rf target/jwt-authorization-${version}
    echo "打包完成"
done