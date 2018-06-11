go-bindata -pkg file -o all.go ./all.txt
rm ./file/*
mv ./all.go file/
go build
rm ./file/*
