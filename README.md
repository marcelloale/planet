# planet


## Utilizando Air com Docker Compose

Air é uma ferramenta que permite o live reloading de aplicações Go. Para utilizá-lo com Docker Compose, siga os passos abaixo:

> **Nota:** A partir da versão 1.27.0, o `docker-compose` vem integrado ao Docker e pode ser chamado como `docker compose` (sem o hífen). Por exemplo, para verificar a versão, você pode usar:
> ```sh
> docker compose version
> ```


1. Primeiro, construa as imagens com o comando:
`docker compose build`
Constrói as imagens definidas no `docker-compose.yml`.
    ```sh
    docker compose build
    ```

2. Em seguida, inicie os contêineres:
`docker-compose up`
Inicia os contêineres definidos no `docker-compose.yml`.
    ```sh
    docker compose up
    ```

3. Se quiser executar um comando em um contêiner em execução, você pode usar:
`docker-compose exec`
Executa um comando em um contêiner em execução.
    ```sh
    docker compose exec app bash
    ```

4. Para parar e remover os contêineres, redes e volumes, utilize:
`docker-compose down`
Para e remove os contêineres, redes e volumes definidos no `docker-compose.yml`.
    ```sh
    docker compose down
    ```

