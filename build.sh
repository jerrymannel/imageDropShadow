rm -rf bin || true

echo "env GOOS=linux GOARCH=386 go build -ldflags="-s -w" -o bin/imageDropShadow-linux-386 ."
env GOOS=linux GOARCH=386 go build -ldflags="-s -w" -o bin/imageDropShadow-linux-386 .

echo "env GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o bin/imageDropShadow-linux-amd64 ."
env GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o bin/imageDropShadow-linux-amd64 .

echo "env GOOS=darwin GOARCH=amd64 go build -ldflags="-s -w" -o bin/imageDropShadow-mac ."
env GOOS=darwin GOARCH=amd64 go build -ldflags="-s -w" -o bin/imageDropShadow-mac .

echo "env GOOS=windows GOARCH=386 go build -ldflags="-s -w" -o bin/imageDropShadow-windows-386.exe ."
env GOOS=windows GOARCH=386 go build -ldflags="-s -w" -o bin/imageDropShadow-windows-386.exe .

echo "env GOOS=windows GOARCH=amd64 go build -ldflags="-s -w" -o bin/imageDropShadow-windows-amd64.exe ."
env GOOS=windows GOARCH=amd64 go build -ldflags="-s -w" -o bin/imageDropShadow-windows-amd64.exe .


cd bin

md5sum imageDropShadow-linux-386 > md5sum.txt
md5sum imageDropShadow-linux-amd64 >> md5sum.txt
md5sum imageDropShadow-mac >> md5sum.txt
md5sum imageDropShadow-windows-386.exe >> md5sum.txt
md5sum imageDropShadow-windows-amd64.exe >> md5sum.txt