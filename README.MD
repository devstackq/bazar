bazar-platform

Golang/Postgres

запуск БД  и приложения в докере
docker-compose up

swagger run
swag init -g ./server/server.go -d ./internal/ --parseInternal --output ./docs

Функциональность bazar-platform:

url : http://localhost:6969
-------------------------------

Админ панель(опционалка)in further:

/v1/any/:id :GET return 1 item
/v1/any :GET - return listItems

Вначале добавить, по убыванию:
Роли, Страну, Город, Юзера,  Бренд, Модель и .т.д

Создание ролей(админ, продавец, покупатель, тд)
<!-- header:token(admin) -->
/v1/role : POST
Request : {name : "saler"}; Response{Status, Message}
|| insert into bazar_roles(name) values('saler');


Страна - Грузия, Оаэ, Россия...
 header[access_token]
/v1/country : POST; {"name" : ""};  Response{Status, Message}

Город - Тбилииси, Дубай, Москоу
 header[access_token]
/v1/city/:country_id : POST; {"name" : ""}; Response{Status, Message}

Бренд  - Tesla, Toyota...
 header[access_token]
/v1/brand : POST; {"name" : ""}; Response{Status, Message}

Модель - Model X3, S;  Supra, Camry...
 header[access_token]
/v1/model/:brand_id : POST; {"name" : ""}; Response{Status, Message}

Категории - Легковые, Мотоциклы, Яхты...
 header[access_token]
/v1/category : POST; {"name" : ""}; Response{Status, Message}

Состояние - Новые, БУ, Битые
 header[access_token]
/v1/state : POST; {"name" : ""}; Response{Status, Message}

Топливо - Дизель, Бензин, Электричество
 header[access_token]
/v1/fuel : POST; {"name" : ""}; Response{Status, Message}

Трансмиссия - Акпп, Мкпп, Гибрид
 header[access_token]
/v1/transmission : POST; {"name" : ""}; Response{Status, Message}

Привод - Передний, Задний, Полный
 header[access_token]
/v1/drive_unit : POST; {"name" : ""}; Response{Status, Message}

Тип кузова -  Седан, купе, кабриолет..
 header[access_token]
/v1/body_type : POST; {"name" : ""}; Response{Status, Message}

Цвет - красный, белый...
 header[access_token]
/v1/color : POST; {"name" : ""}; Response{Status, Message}
------------------------------

Создание юзера(продавец)

/v1/auth/signup : POST
Request: 

{
	"email": "user@mail.kz",
	"username":"saler",
	"phone":87471370784,
	"first_name":"Glize",
	"last_name":"Union",
	"password":"123user",
    "company" : "dillerX",

    fetch(/v1/country) return list(countries)
    "country": {
        "id" : int
        "city" : {
            fetch(/v1/city/:countryID)  return list(cities)
            id: int
        }
    }
    fetch(/v1/role)   return list(roles)
    "role" {
        "id : int
    }
}

Response {statusCode, Message, Set Header[access_token, refresh_token]}
------------------------------

Вход в систему
/v1/auth/signin : POST
Request: 
{
	 string username
	 string password
}

Response {statusCode, Message, Header[access_token, refresh_token]}
------------------------------

Обновление access токена, при  статусе 401
/v1/auth/refresh -> POST
Request(Header[refresh_token]) 
Response {statusCode, Message, Header[access_token, refresh_token]}
------------------------------
(Если refresh токен недействителен, то redirect на  signin)
------------------------------

Выход из системы
/v1/auth/logout
Request(Header[access_token])
Response {statusCode, Message, Remove access/refresh token}
------------------------------

Профиль
Request : header[access_token]
/v1/profile : GET
/v1/machine/user?page_num=1 GET

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
Получить список полей с данными для :    transmision, brand, category, etc -> fetch(v1/transmission, v1/category, etc...)
1.1
{
    "title": "Sale the best car, because i need money",
    "description": "this car bla bla...",
    "year": 2009,
    "price": 29500.1,
    "odometer": 3000,
    "horse_power": 200,
    "volume": 3.5,
    "transmission": {
        "id": 1
    },
    "body_type": {
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
        "id": 1,
           "city": {
             "id": 1
        }
    },

    "fuel": {
        "id": 1
    },
    "drive_unit": {
        "id": 1
    }
}

Response {Status, lastCarID}

async, await
1.2
Zатем выгрузка фото на файловый сервер
/v1/machine/upload/lastCarID
Request {key : MultipartFile}
Resonse : {Status, Message}
------------------------------

Главная:

Вернет список машин, по дате добавления
пагинация по query-param; page_num; по умолчанию = 1 ( выдает по  9 машин)
/v1/machine?page_num=2 : GET
Response{
    []Machine{...}
}
------------------------------

Пагинация 
/v1/machine?page_num
limit = default 9
может с фильтром/сортировкой - работать
------------------------------

Карточка машины (Вернет машину по его ID)
/v1/machine/:id : GET
Response{
    Machine{...}
}
------------------------------

Фильтрация, по категории, состоянию,  бренду, модели

/v1/machine/filter query-param : ["category", "state", "brand", "model"] id
Response : {
    []Machine{...}
}

Фильтрация по цене - /v1/machine/filter query-param [priceTo/proceFrom] float
------------------------------

Поиск по title или description
/v1/machine/search/:POST query-param key_word
Response : {
    []Machine{...}
}
------------------------------

Сортировка, (может включать поля фильтрации)
/v1/machine/filter query-param : ["category", "state", "brand", "model", priceFrom, priceTo] +  
sort_created_at/sort_price/sort_year/sort_odometer - asc/desc
------------------------------