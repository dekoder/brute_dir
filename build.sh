go-bindata -pkg file -o all.go ./all.txt
mkdir file
mv ./all.go file/
go build
rm -rf ./file
