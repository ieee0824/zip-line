docker build -t zipl .

echo "drawin build"
mkdir -p macOS
docker run -e GOOS=darwin -v $PWD/macOS:/tmp/build zipl go build -o /tmp/build/zipl cmd/zipl/main.go
chown -R $SUDO_USER:$SUDO_USER macOS

echo "linux build"
mkdir -p linux
docker run -e GOOS=linux -v $PWD/linux:/tmp/build zipl go build -o /tmp/build/zipl cmd/zipl/main.go
chown -R $SUDO_USER:$SUDO_USER linux