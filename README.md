# todo_app_golang
Использованные библеотеки:
    "github.com/gorilla/mux" - для роутера
    "github.com/streadway/amqp" - message broker 
    "github.com/urfave/cli" - написал CLI команду для старта сервера
    "net/smtp" - для отправки email сообщения
    "github.com/valyala/fasthttp" , "net/http" - для сервера
    "github.com/go-pg/pg" - библиотека для работы с Postgres
    "github.com/go-ozzo/ozzo-validation","github.com/go-ozzo/ozzo-validation/is" - для проверки валидации email и пароля
    "golang.org/x/crypto/bcrypt" - для encrypt password

Запускается при команде 
  .\main start

Основные методы:
  main.go - создание сервера и rabbitMQ-consumer (тоесть получателя), отправка email
  endpoint.go - создание эндпойнтов, все хандлеры тут, эндпойнты и для Tasks, и для авторизации (7.2 пункт не выполнил, т,к фронт не писал)
  endpoint_test.go - юнит тесты для эндпойнтов (2 тестов)
  models.go - находится модели и интерфейсы для user и task
  postgre.go - подключение к дб (я использовал postgre) и методы
  
