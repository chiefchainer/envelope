# envelope
указываем токен в константах main.go где access token

просто запустить скрипт
```
go run main.go
```
запустить тесты (поменять кол-во тестов константа repeatCount в main_test.go)
```
go test -v
```

если не выдает резы проверьте что у вас нет: 29 Rate limit reached или же не указан токен или же токен мертвый или же айпи выдачи токена не совпадает с актуальным айпи 

```
MacBook-Pro-Burov:envelode burov$ go test -v
=== RUN   TestMainExecution
Количество "энвилоуп" в комментариях: 347
Время, затраченное на подсчёт энвилоупов: 4.094 с
    main_test.go:20: Тест 1 занял 4.094 секунд
Количество "энвилоуп" в комментариях: 347
Время, затраченное на подсчёт энвилоупов: 1.244 с
    main_test.go:20: Тест 2 занял 1.244 секунд
Количество "энвилоуп" в комментариях: 347
Время, затраченное на подсчёт энвилоупов: 1.090 с
    main_test.go:20: Тест 3 занял 1.090 секунд
Количество "энвилоуп" в комментариях: 347
Время, затраченное на подсчёт энвилоупов: 1.043 с
    main_test.go:20: Тест 4 занял 1.043 секунд
Количество "энвилоуп" в комментариях: 347
Время, затраченное на подсчёт энвилоупов: 1.135 с
    main_test.go:20: Тест 5 занял 1.136 секунд
Количество "энвилоуп" в комментариях: 347
Время, затраченное на подсчёт энвилоупов: 1.022 с
    main_test.go:20: Тест 6 занял 1.022 секунд
Количество "энвилоуп" в комментариях: 347
Время, затраченное на подсчёт энвилоупов: 1.117 с
    main_test.go:20: Тест 7 занял 1.117 секунд
Количество "энвилоуп" в комментариях: 347
Время, затраченное на подсчёт энвилоупов: 1.142 с
    main_test.go:20: Тест 8 занял 1.142 секунд
Количество "энвилоуп" в комментариях: 347
Время, затраченное на подсчёт энвилоупов: 1.169 с
    main_test.go:20: Тест 9 занял 1.169 секунд
Количество "энвилоуп" в комментариях: 347
Время, затраченное на подсчёт энвилоупов: 2.278 с
    main_test.go:20: Тест 10 занял 2.278 секунд
    main_test.go:26: Среднее время выполнения 10 тестов: 1.534 секунд
--- PASS: TestMainExecution (15.34s)
PASS
ok  	envelode	16.325s
```
