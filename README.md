# labs-four

Projeto de conclusão de pós-graduação

## Indice
1. [Build-Dev](#build-dev)
2. [DockerCompose](#dockercompose)
3. [Testes](#testes)

## Build-Dev

O app é acessível na porta 8080.

- na raiz do projeto rodar o comando
```bash
go run ./cmd/auction
```

## DockerCompose

O docker compose pode ser configurado para escolher duas possibilidades de execução:

  - **Rate Limit com Redis**
  - **Rate Limit em Memoria**

Para isso, a variável ***RATE_LIMIT_TYPE*** no docker compose deve ser alterada para ***redis*** ou ***memory***.

A cada troca de valor desta chave, rodar o comando

```bash
docker compose up --build
```

# Testes

Para os testes, após a execução do docker compose, podemos usar:

- para testar por ip:
```bash
for i in {1..7}; do curl -i http://localhost:8080/hello; echo; done
```

- para testar por token:
```bash
for i in {1..20}; do curl -i -H "API_KEY: meu-token" http://localhost:8080/hello; echo; done
```

Tanto usando ***memory*** quanto o ***redis*** o comportamento é o mesmo.
