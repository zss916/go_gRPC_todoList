.PHONY: mysql_start

mysql_start:
	@echo "开启本地数据库"
	@mysql.server start

open_mysql:
	@echo "进入数据库"
	@mysql -u root -p

show_databases:
	@echo "显示所有已有数据库"
	@show databases;

create_database:
	@echo "创建数据库,例子(create database todo_db charset=utf8mb4;)"
	@create database todoList;

use_database:
	@echo "选中数据库"
	@use todoList;

check_database:
	@echo "查看当前使用的数据库"
	@select database();

show_tables:
	@echo "显示当前数据库所有的表"
	@show tables;

desc_table:
	@echo "显示当前数据表user的结构"
	@DESC user;

show_table_data:
	@echo "查看user表里面所有的数据/ 查看表中id 就用SELECT id FROM user"
	@SELECT * FROM user


redis:
	@echo "brew开启redis"
	@brew services start redis


swag:
	@echo "swag 初始化, 打开http://localhost:3000/swagger/index.html"
	@swag init -g main.go


docker_compose_run:
	@echo "docker 启动项目"
	@docker-compose up -d


#RunApi
#国内源: https://goproxy.io
#grpc 集成ETCD => https://blog.51cto.com/u_15314183/5457223
