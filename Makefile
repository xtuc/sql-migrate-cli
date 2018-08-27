build:
	cd cmd/migrate && go build -ldflags="-s -w"
	cd cmd/undo && go build -ldflags="-s -w"

install:
	sudo cp cmd/migrate/migrate /usr/bin/xtuc-mysql-migrate
	sudo cp cmd/undo/undo /usr/bin/xtuc-mysql-undo
