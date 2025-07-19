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
ACCESS_TOKEN_SECRET=MY_SECRET
ACCESS_TOKEN_EXPIRE_DURATION=48h
REFRESH_TOKEN_SECRET=MY_SECRET
REFRESH_TOKEN_EXPIRE_DURATION=120h
```

## 2.Example database diagram

![img_1.png](img_1.png)

## 3.Example request & response

- API Concert List (Success)
![img.png](img.png)

- API Concert Booking (Success)
![img_7.png](img_7.png)

- API Concert Booking (Failed Sale Not Started)
![img_2.png](img_2.png)

- API Concert Booking (Failed Overbooking)
![img_3.png](img_3.png)

- API Concert Purchase History (Success)
![img_5.png](img_5.png)

- API User Purchase History (Success)
![img_4.png](img_4.png)

- API Reach Rate Limiter
![img_6.png](img_6.png)
