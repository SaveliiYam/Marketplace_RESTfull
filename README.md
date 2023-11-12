# Маркетплейс на языке Golang

## Описание 

* REST API сервис.
* Заточен под создание шаблонных магазинов.
* Обладает возможностью создания категорий, брендов и товаров
* Имеется полноценная аутентификация и авторизация. Передача через JWT токены доступа.
* Реализована полноценная корзина хранения. Доступ к корзине имеют только авторизированные пользователи.
* Используется база данных PostgreSQL.

## Обращение к API
### Создание пользователя
* POST localhost:8000/auth/sign-up
```json
{
    "username": "username",
    "name":"name",
    "password":"password",
    "status":true //если создается администратор
}
```
* Вывод
```json
{
    "id": 1
}
```
### Авторизация
* Токен действителен 12 часов
* POST localhost:8000/auth/sign-in
```json
{
    "username": "username",
    "password":"password"
}
```
* Вывод
```json
{
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2OTk3NTI1NDQsImlhdCI6MTY5OTcwOTM0NCwidXNlcl9pZCI6MX0.EPqYdsM4t2Fs3ssPDtUYaisK_n9exMuhD6vQPwWIHrM"
}
```

### Создание каталогов
#### Предусмотрена защита от повторяющихся наименований
#### При создании каталога на вывод идет id данного каталога
* POST localhost:8000/api/v1/admin/categories/
```json
{
    "title":"Летняя обувь",
    "description":"обувь для лета"
}
```
* POST localhost:8000/api/v1/admin/brand/
```json
{
    "title":"Nike",
    "description":"Американский бренд"
}
```
* POST localhost:8000/api/v1/admin/product/
```json
{
    "title":"Air Jordan",
    "description":"Летом может быть жарко",
    "price": "10.2",
    "brand": "Nike",
    "category": "Летняя обувь"
}
```

Товары в корзину может добавлять только пользователь:
* POST localhost:8000/api/v1/user/basket/


```json
{
    "product_id":1
}
```

* Пример ответа для всех предыдущих запросов
```json
{
    "id": 1
}
```

### Получение каталогов
#### Можно получить каталоги либо по id, либо целиком
* GET localhost:8000/api/v1/user/categories/
* GET localhost:8000/api/v1/admin/categories/
```json
{
    "data": [
        {
            "id": 1,
            "title": "Летняя обувь"
        },
        {
            "id": 2,
            "title": "Зимняя обувь"
        }
    ]
}
```
* GET localhost:8000/api/v1/user/brands/
* GET localhost:8000/api/v1/admin/brands/
```json
{
    "data": [
        {
            "id": 1,
            "title": "Nike",
            "description": "Американский бренд"
        },
        {
            "id": 2,
            "title": "Adidas",
            "description": "Тоже американский бренд"
        }
    ]
}
```

* GET localhost:8000/api/v1/user/products/
* GET localhost:8000/api/v1/admin/products/
```json
{
    "data": [
        {
            "id": 1,
            "title": "Air Jordan",
            "description": "Летом может быть жарко",
            "price": "10.20",
            "brand": "1",
            "category": "1"
        },
        {
            "id": 2,
            "title": "Air Jordan Zima",
            "description": "Зимой будет холодно",
            "price": "10.22",
            "brand": "1",
            "category": "2"
        }       
    ]
}
```