# MyGolangWeb
This repository is my golang web project, using material design style frontend and golang backend. 


### 1.Install Source Code

```shell
mkdir webapp webapp/src webapp/pkg webapp/bin
cd webapp
source gvp
git clone https://github.com/zhangmingkai4315/MyGolangWeb.git src
cd src
gpm install
```

### 2.Install Database


We will use mysql for data persist layer, the most simple way to install mysql is use docker, just follow the below commands, you will get a mysql instance for develop:

```
$ docker pull mysql
$ docker create --name mysql-data --volume /var/lib/mysql mysql
$ docker run --detach --name mysql-process --env MYSQL_ROOT_PASSWORD=123456 --volumes-from mysql-data mysql
$ mysql -uroot -p
Enter password: 

$mysql -uroot -p < setup.sql 
```
if you are using mac osx, remember your also should do port forward for your docker machine to the real docker containe.

```
$ docker-machine ip
192.168.99.100
$ docker create --name mysql-data --volume /var/lib/mysql mysql
$docker run --detach --name mysql-process -p 3306:3306 --env MYSQL_ROOT_PASSWORD=123456 --volumes-from mysql-data mysql

$ mysql -h 192.168.99.100 -uroot -p
Enter password: 

$mysql -uroot -h 192.168.99.100 -p < setup.sql 

```

