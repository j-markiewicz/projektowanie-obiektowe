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
