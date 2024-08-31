## REST сервер
* POST api/add
* GET api/list

### Go 1.22, Postgresql, logrus, chi, jwt, docker  

При сохранении заметок орфографические ошибки валидируются при помощи сервиса Яндекс.Спеллер  
Авторизация и аутентификация при помощи jwt. Логин и пароль передаются в теле запроса в ручку POST "/login", в ответ приходит токен.  
Доступ к ручкам из группы "/api" только при наличии Bearer token.    
Пользователи имеют доступ только к своим заметкам - id пользователя берется из токена. 