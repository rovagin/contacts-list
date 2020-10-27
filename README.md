Технологии: React + Go + Pg + Docker.
 
Написать приложение «Телефонная книга».  
Поля: фио, телефон, email, текстовая заметка  

UI:
- Список контактов (сокр. Инфо о контакте)  
- Расширенная карточка инфо о контакте (отображение всех данных контакта)  
 
Функции:
- Создание нового контакта - back
- Поиск по Фио - front
- Поиск по телефону - front
- Изменить контакт - back
- Удалить контакт - back
- Проверка дубликатов (по номеру телефона) - back при создании
- Валидация вводимых данных (номер – цифры, фио – буквы, email) - front + back
 
Generate Some stub data: https://next.json-generator.com/EJQL5qiPK

TODO: добавить миграции

Для запуска выполнить: `docker-compose up -d`

После запуска необходимо выполнить: 
1. Зайти в контейнер mongo   
`docker-compose exec -it bash mongodb`  
2. Выполнить  
`mongo --exec 'db.users_contacts.insert({_id: 0})' mongodb://admin:password@localhost:27017/users_contacts?authSource=admin`
