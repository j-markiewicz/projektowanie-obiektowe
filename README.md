# Projektowanie Obiektowe

## [Zadanie 1 - Paradygmaty](./1/)

Sortowanie bąbelkowe

Proszę napisać [program w Pascalu](./1/1.pas), który zawiera dwie procedury, jedna generuje listę 50 losowych liczb od 0 do 100. Druga procedura sortuje liczbę za pomocą sortowania bąbelkowego.

- [x] 3.0 Procedura do generowania 50 losowych liczb od 0 do 100
- [x] 3.5 Procedura do sortowania liczb
- [x] 4.0 Dodanie parametrów do procedury losującej określającymi zakres losowania: od, do, ile
- [x] 4.5 5 testów jednostkowych testujące procedury
- [x] 5.0 [Skrypt w bashu](./1/run.sh) do uruchamiania aplikacji w Pascalu via docker

Termin: 25.03

## [Zadanie 2 - Wzorce architektury](./2/)

Symfony (PHP)

Należy stworzyć aplikację webową na bazie frameworka Symfony na obrazie kprzystalski/projobj-php:latest. Baza danych dowolna, sugeruję SQLite.

- [x] 3.0 Należy stworzyć jeden [model](./2/products/src/Entity/Product.php) z [kontrolerem](./2/products/src/Controller/ProductsController.php) z produktami, zgodnie z CRUD (JSON)
- [x] 3.5 Należy stworzyć [skrypty do testów endpointów via curl](./2/tests/) (JSON)
- [ ] 4.0 Należy stworzyć dwa dodatkowe kontrolery wraz z modelami (JSON)
- [ ] 4.5 Należy stworzyć widoki do wszystkich kontrolerów
- [ ] 5.0 Stworzenie panelu administracyjnego

Termin: 2.04

## [Zadanie 3 - Wzorce kreacyjne](./3/)

Spring Boot (Kotlin)

Proszę stworzyć prosty serwis do autoryzacji, który zasymuluje autoryzację użytkownika za pomocą przesłanej nazwy użytkownika oraz hasła. Serwis powinien zostać wstrzyknięty do kontrolera (4.5).
Aplikacja ma oczywiście zawierać jeden kontroler i powinna zostać napisana w języku Kotlin. Oparta powinna zostać na frameworku Spring Boot. Serwis do autoryzacji powinien być singletonem.

- [x] 3.0 Należy [stworzyć jeden kontroler](./3/src/main/kotlin/example/projektowanieobiektowe/po3/Po3Application.kt) wraz z danymi wyświetlanymi z listy na endpoint’cie w formacie JSON - Kotlin + Spring Boot
- [x] 3.5 Należy stworzyć klasę do autoryzacji (mock) jako Singleton w formie eager
- [x] 4.0 Należy obsłużyć dane autoryzacji przekazywane przez użytkownika
- [ ] 4.5 Należy wstrzyknąć singleton do głównej klasy via @Autowired lub kontruktor (constructor injection)
- [ ] 5.0 Obok wersji Eager do wyboru powinna być wersja Singletona w wersji lazy

![zapytania do serwera przy pomocy curl, pokazujące metody GET na liście i elemencie, POST na liście, oraz PATCH, PUT, i DELETE na elemencie](./3/curl.png)

|   |   |   |
| - | - | - |
| ![formularz z nazwą użytkownika i hasłem](./3/login-1.png) | ![wypełniony formularz](./3/login-2.png) | ![request w devtoolsach pokazujących wysłany nagłówek Authorization](./3/login-3.png) |
