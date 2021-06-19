init_mysql:
	@echo "initialize MySQL database"
	@mysql -uroot -h 127.0.0.1 -P 6606 gofixtures_test < initialize.mysql.sql

test: init_mysql
	@echo "implement me!"