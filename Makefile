appVersion=1.0-build

configuration = ./config/config-local.yml

build:
	go build -o out/server ./main.go

run:
	out/server -configFile=${configuration}