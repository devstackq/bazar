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
	string `json:"email"`
	string `json:"username"`
	int `json:"phone"`
	string `json:"first_name"`
	string `json:"last_name"`
	string `json:"password"`
    int `json:"country_id"`
	int    `json:"city_id"`
	int `json:"role_id"`
}
Response {statusCode, Message, Header[access_token, refresh_token]}

Вход в систему
/v1/signin : POST
Request: 
{
	string `json:"username"`
	string `json:"password"`
}

Response {statusCode, Message, Header[access_token, refresh_token]}
------------------------------
Обновление access/refresh токена
/v1/refresh -> middlewares(return incorrect access token)
Request(Header[refresh_token])
Response {statusCode, Message, Header[access_token, refresh_token]}
------------------------------


Профиль

Request : header[access_token]

/v1/profile : GET
Response:  {
    UserInfo
    ListCreatedCarsByUserID    
}

Создание машины 
Прежде чем добавить машины: -> добавить  в БД, админские crud(cateogry,model, brand, etc)

/v1/machine : POST
Request:
 header[access_token]
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
    },
    "creator": {
        "id": "1"
    }
}
------------------------------

Главная:

Вернет машину по его ID
/v1/machine/:id : GET
Response{
    Machine{...}
}

Вернет список машин, по дате добавления
/v1/machine : GET

Фильтрация

/v1/filter query-param : ["category", "state", "brand", "model"] id
Response : {
    []Machine{...}
}

Фидьтрация по цене - /v1/filter query-param [priceTo/proceFrom] float

------------------------------

Поиск по title или description
/v1/search/:keyword :POST
Response : {
    []Machine{...}
}

------------------------------
Сортировка
/v1/filter query-param : ["category", "state", "brand", "model"] +  sort_created_at/sort_price/sort_year/sort_odometer - asc/desc

------------------------------

for Guests:
-------------------------------
Вернет список созданных машин - пользователем созданных
/v1/machine/user/:id :GET
Response {
    []Machine{}...
}
-------------------------------
Вернет информацию о продавце
/v1/profile/bio/:id
Response{
    User{...}
}
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