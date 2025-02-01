Этот проект реализует сервис для загрузки превью (thumbnails) видеороликов YouTube. 

Установка
1. Скопировать код приложения - загрузить его напрямую или клонировать данный репозиторий: 
git clone git@github.com:IvanMishnev/youtube-thumbnails-service.git
cd youtube-thumbnails-service

3. Установить зависимости: 
go mod tidy

Запуск сервера
1. Перейти в папку server/cmd:
cd server/cmd

2. Запустить сервер командой: go run main.go

Запуск и использование клиента

Приложение предусматривает работу с 2 параметрами командной строки:
--URLs: укажите ссылки на видеоролики YouTube, перечислив их через запятую без пробела
--async: опциональный флаг. Укажите его для асинхронной обработки ссылок

1. Перейти в папку client-cli
2. Запустить команду go run main.go *args*, где *args* - аргументы командной строки.
Примеры использования: 
go run main.go --async --URLs https://www.youtube.com/watch?v=CWeXOm49kE0,https://youtu.be/ujChUYkPvec,https://www.youtube.com/watch?v=WqEweV0eScg
go run main.go --URLs https://youtu.be/ujChUYkPvec


Файлы будут загружены в автоматически созданную папку client-cli/thumbnails


База с кэшем изображений cache.db автоматически создается в папке server. Для очистки кэша достаточно удалить данный файл.