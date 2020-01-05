auth_user:
	go build ./bin/auth_server main.go

proto:
	protoc --micro_out=proto/ --go_out=proto/ proto_files/user/*.proto