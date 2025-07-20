**Readme.MD**

## 1.Example .env
```
APP_VERSION=v1
DB_HOST=127.0.0.1
DB_NAME=postgres
DB_USER=postgres
DB_PASSWORD=PASSWORD
DB_PORT=5432
DB_SSL=disable
DB_TIMEZONE=Asia/Jakarta
DB_AUTO_MIGRATE=false
REDIS_ADDR=127.0.0.1:6379
REDIS_PASSWORD=
HTTP_HOST=127.0.0.1
HTTP_PORT=10001
```

## 2.Example database diagram

![img.png](img.png)

## 3.Example request & response

- API Create New Article (Success)
![img_2.png](img_2.png)

- API Create New Article (Failed)
![img_1.png](img_1.png)

- API Get List Article (Success - No Filter)
![img_3.png](img_3.png)

- API Get List Article (Success - Filter Keyword)
![img_4.png](img_4.png)

- API Get List Article (Success - Filter Author Name)
![img_5.png](img_5.png)

## 4. Unit test
![img_6.png](img_6.png)