## Билд
```
docker-compose build
```

## Запуск
```
docker-compose up -d
```

## Остановка
```
docker-compose stop
```

## Чистка БД
```
rm -rf ttdata
mkdir ttdata
```

## Конфиг 
### Конфиг разработки
> src/app/config/dev_config.yml

### Конфиг деплоя
> src/app/config/prod_config.yml

### Задать конфиг
> src/app/main.go:11 (env)