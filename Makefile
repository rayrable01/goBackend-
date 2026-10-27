service-run:
	export conn_string=postgres://postgres:0000@localhost:5432/postgres & \
	go run main.go