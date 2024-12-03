
# Sistema de Gerenciamento de Filas com RabbitMQ e Docker


### Visão Geral
Este projeto implementa um sistema de filas robusto utilizando Golang, RabbitMQ e Docker, projetado para processamento eficiente de mensagens com arquitetura de produtor-consumidor.




## Principais Características

### Tecnologias Utilizadas

- Golang para desenvolvimento do backend
- RabbitMQ para gerenciamento de filas de mensagens
- Docker para containerização e implantação simplificada


####

## Arquitetura

- Produtor: Responsável por enviar mensagens para a fila
- Consumidor: Processa mensagens recebidas e gera gráficos com os dados



####

## Funcionalidades

- Envio de mensagens para filas do RabbitMQ
- Consumo e processamento de mensagens
- Geração de visualizações gráficas dos dados processados

## Benefícios
- Escalabilidade através de containerização
- Processamento assíncrono de mensagens
- Fácil monitoramento e visualização de dados

## Pré-requisitos
 - Docker
 - Golang


# Configuração e Execução

Clonar o repositório
```shell
git clone [esse-repositorio]
```

Subir os serviços com Docker (RabbitMQ)
```shell
docker-compose up -d
```

Baixe as Dependência do projeto
```shell
go mod tidy
```

# Rodar aplicação
Para executar esses comandos, abra terminais diferentes e execute cada um deles em paralelo, pois todos precisam estar rodando simultaneamente. 

Inicie a fila (na raiz do projeto)

```shell
go run consumer/main.go
```

Inicie a API que vai enviar os itens para fila(na raiz do projeto)

```shell
go run producer/main.go
```

Agora acesse o RabbitMQ no seu navegador atraves dessa url

```shell
http://localhost:15672/
```
Credenciais do RabbitMQ

usuario: Admin 

Senha : Admin


## Agora envie uma mensagem para a fila por essa url.

Caso seja pelo Curl ou postman deve ser Metodo GET.

*localhost:8080/send?msg=Hello*

Ou cole essa url em seu *navegador*.

A aplicação agora está integrada a uma ferramenta de monitoramento avançada, proporcionando insights detalhados sobre desempenho, escalabilidade e saúde do sistema. Com essa solução, você pode acompanhar métricas em tempo real, identificar gargalos de performance e tomar decisões estratégicas para otimizar e escalar sua infraestrutura de forma inteligente e eficiente.








## Autores

- [@Guilherme Oliveira](https://www.linkedin.com/in/guilherme-oliveira-121b16239/)

