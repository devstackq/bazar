bazar-platform

Функциональность:

Transport - http(Gin lib)
Database - postgres

testUrl: http://localhost:6969


-------------------------------
Вначале надо создать роли
header:token(admin) - optional
/v1/role : POST
Request : {name : "saler"}
-------------------------------
insert into bazar_roles(name) values('saler');

  insert into bazar_user(email, password, phone, username, first_name, last_name, created_at, role_id, country_id) values('devstackq@mail.kz', 'password', 87471370784, 'devstackq', 'Bekzh', 'Satarkul', '2022-04-12 04:05:06', 1,1);
-------------------------------

Создание юзера

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
Response {statusCode, Message}

Вход в систему
/v1/signin : POST
Request: 
{
	string `json:"email"`
	string `json:"username"`
	string `json:"password"`
}
Response {statusCode, Message}
------------------------------

Профиль

header:token -  middleware(token->) getUserID
/v1/profile : GET
Response:  {
    UserInfo
    ListCreatedCarsByUserID    
}

Вернет список созданных машин - пользователем созданных

header:token -  middleware(token->) getUserID
/v1/machine/user :GET

Response {
    []Machine{}...
}

Вернет информацию о продавце
header:token -  middleware(token->) getUserID
/v1/profile/bio
Response{
    User{...}
}
------------------------------

header:token -  middleware(token->) getUserID
Создание машины 
Прежде чем добавить машины: -> добавить  в БД, админские crud(cateogry,model, brand, etc)

/v1/machine : POST
Request:
{
    "title": "Sale the best car, because i need money",
    "description": "this car bla bla...",
    "year": 2000,
    "price": 20000,
    "odometer": 30002,
    "horse_power": 100,
    "volume": 3.5,
    "trans": {
        "id": 2
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
        "id": 2
    },
    "country": {
        "id": 2
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
        "id": 1
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

/v1/filter query-param : ["category", "state", "brand", "model"]
Response : {
    []Machine{...}
}

Поиск по title или description
/v1/search/:keyword :POST
Response : {
    []Machine{...}
}
------------------------------

Админка(опционалка) CRUD по:

Категориям - Лешковые, Мотоциклы, Яхты...
/v1/category : POST; {"name" : ""}

Состояние - Новые, БУ, Битые
/v1/state : POST; {"name" : ""}

Бренд  - Tesla, Toyota...
/v1/brand : POST; {"name" : ""}

Модель - Model X3, S;  Supra, Camry...
/v1/model/:brand_id : POST; {"name" : ""}

Топливо - Дизель, Бензин, Электричество
/v1/fuel : POST; {"name" : ""}

Трансмиссия - Акпп, Мкпп, Гибрид
/v1/trans : POST; {"name" : ""}

Привод - Передний, Задний, Полный
/v1/drive_unit : POST; {"name" : ""}

Тип кузова -  Седан, купе, кабриолет..
/v1/body_type : POST; {"name" : ""}

Цвет - красный, белый...
/v1/color : POST; {"name" : ""}

Страна - Каз-ан, Оаэ, Руссия...
/v1/country : POST; {"name" : ""}

Город - Азтана, Дубай, Москоу
/v1/city/:country_id : POST; {"name" : ""}
------------------------------


entity.POST("", handler.Create)
entity.GET("", handler.GetList)
entity.GET("/:id", handler.GetByID)


