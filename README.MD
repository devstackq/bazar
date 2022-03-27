bazar-platform

Функциональность:

Transport - http(Gin lib)
Database - postgres

testUrl: http://localhost:6969


-------------------------------
Вначале надо создать роли

header:token(admin)
/v1/role : POST
Request : {name : "saler"}
-------------------------------
insert into bazar_roles(name) values('saler');

  insert into bazar_user(email, password, phone, username, first_name, last_name, created_at, role_id, country_id) values('devstackq@mail.kz', 'password', 87471370784, 'devstackq', 'Bekzh', 'Satarkul', '2022-04-12 04:05:06', 1,1);
-------------------------------

Создание юзера(продавец)

/v1/signup : POST
Request: 
{
	string  email
	string username
	int phone
	string first_name
	string last_name
	string password
    string company
    int country_id -> fetch(/v1/country)
	int    city_id -> fetch(/v1/city/:countryID)
}
Response {statusCode, Message, Header[access_token, refresh_token]}
------------------------------

Вход в систему
/v1/signin : POST
Request: 
{
	string `json:"username"`
	string `json:"password"`
}

Response {statusCode, Message, Header[access_token, refresh_token]}
------------------------------

Обновление access/refresh токена, при случаях 401 статуса
/v1/refresh -> middlewares(return incorrect access token)
Request(Header[refresh_token])

Response {statusCode, Message, Header[access_token, refresh_token]}
(Если refresh токен устарел, то redirect - signin)
------------------------------

Выход из системы
/v1/logout
Request(Header[access_token])
Response {statusCode, Message}
------------------------------

Профиль

Request : header[access_token]

/v1/profile : GET
Response:  {
    UserInfo
    ListCreatedCarsByUserID
}
------------------------------

Создание машины 
Прежде чем добавить машины: -> добавить  в БД, админские crud(cateogry,model, brand, etc)

/v1/machine : POST
Request:
 header[access_token]
 По спискам  transmision, brand, category, etc -> fetch(v1/trans, v1/category, etc...)

{
    "title": "Sale the best car, because i need money",
    "description": "this car bla bla...",
    "year": 2009,
    "price": 29500.1,
    "odometer": 3000,
    "horse_power": 200,
    "volume": 3.5,
    "trans": {
        "id": 1
    },
    "body": {
        "id": 1
    },
    "color": {
        "id": 1
    },
    "category": {
        "id": 1
    },
    "brand": {
        "id": 1,
        "model": {
            "id": 1
        }
    },
    "state": {
        "id": 1
    },
    "country": {
        "id": 1
    },
    "city": {
        "id": 1
    },
    "fuel": {
        "id": 1
    },
    "drive_unit": {
        "id": 1
    }
}

создание машины, затем выгрузка фото на файловый сервер
async, await
------------------------------

Главная:

Вернет список машин, по дате добавления
пагинация по query-param; page_num; по умолчанию = 1 (выдаст 9 машин)
/v1/machine?page_num=2 : GET
Response{
    []Machine{...}
}
------------------------------

Карточка машины 

Вернет машину по его ID
/v1/machine/:id : GET
Response{
    Machine{...}
}
------------------------------

Фильтрация

/v1/filter query-param : ["category", "state", "brand", "model"] id
Response : {
    []Machine{...}
}

Фильтрация по цене - /v1/filter query-param [priceTo/proceFrom] float
------------------------------

Поиск по title или description
/v1/search/:POST query-param key_word
Response : {
    []Machine{...}
}
------------------------------

Сортировка, (может включать поля фильтрации)
/v1/filter query-param : ["category", "state", "brand", "model", priceFrom, priceTo] +  sort_created_at/sort_price/sort_year/sort_odometer - asc/desc
------------------------------

Админка(опционалка) CRUD по:

/v1/any/:id :GET;
/v1/any :GET - list

Категориям - Легковые, Мотоциклы, Яхты...
 header[access_token]
/v1/category : POST; {"name" : ""}

Состояние - Новые, БУ, Битые
 header[access_token]
/v1/state : POST; {"name" : ""}

Бренд  - Tesla, Toyota...
 header[access_token]
/v1/brand : POST; {"name" : ""}

Модель - Model X3, S;  Supra, Camry...
 header[access_token]
/v1/model/:brand_id : POST; {"name" : ""}

Топливо - Дизель, Бензин, Электричество
 header[access_token]
/v1/fuel : POST; {"name" : ""}

Трансмиссия - Акпп, Мкпп, Гибрид
 header[access_token]
/v1/trans : POST; {"name" : ""}

Привод - Передний, Задний, Полный
 header[access_token]
/v1/drive_unit : POST; {"name" : ""}

Тип кузова -  Седан, купе, кабриолет..
 header[access_token]
/v1/body_type : POST; {"name" : ""}

Цвет - красный, белый...
 header[access_token]
/v1/color : POST; {"name" : ""}

Страна - Каз-ан, Оаэ, Руссия...
 header[access_token]
/v1/country : POST; {"name" : ""}

Город - Азтана, Дубай, Москоу
 header[access_token]
/v1/city/:country_id : POST; {"name" : ""}
------------------------------

entity.POST("", handler.Create)
entity.GET("", handler.GetList)
entity.GET("/:id", handler.GetByID)