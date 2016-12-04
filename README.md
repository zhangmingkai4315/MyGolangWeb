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

We are using postgre database for data store, so before next step ,you should install it first. After that you need setting you database username and password, and using psql tools to import the database file "setup.sql"

```
psql -h YourDatabaseIp -p YourDabasePort --username=postgres -c "create dabase golangchina"
psgl -h YourDatabaseIp -p YourDabasePort --username=postgres -f setup.sql -d golangchina
```


