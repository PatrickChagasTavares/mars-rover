# Rover em Marte

Este projeto implementa um programa em Go para controlar rovers em Marte. Os rovers são implantados em um planalto retangular e devem ser movidos de acordo com uma série de instruções, mantendo sua posição e orientação.


## Exemplo de uso:
<img width="496" alt="mars-rover" src="https://github.com/PatrickChagasTavares/mars-rover/assets/49497853/40868e6e-1261-43a9-99b3-27bc6df3fae0">


## 🚀 Começando

Siga estas instruções para configurar o projeto na sua máquina local para desenvolvimento e teste.

### 📋 Pré-requisitos

Ferramentas necessárias:

- [Golang](https://golang.org/doc/install)


## 📦 Desenvolvimento

Comandos importantes para rodar o projeto e validar:

- `make run`: Compila e executa o código principal.
- `make test`: Executa os testes do projeto e mostra a cobertura.
- `make test-cover`: O mesmo do `make test`, porém abre o brawser para mais detalhes.
- `make help`

## 🗂 Estrutura do Projeto

### Descrição dos Pacotes e Arquivos Principais:

- `./cmd/main.go`: O código que inicia a aplicação.
- `./internal/domain`: Lógica de negócios relacionada ao controle dos rovers.
- `./internal/domain/model`: Definições de structs relacionadas aos rovers.
- `./internal/domain/usecase`: Casos de uso para controle dos rovers.
- `./pkg/scanner`: Pacote para ler e interpretar instruções do usuário.
- `./makefile`: Arquivo de make para automatizar tarefas comuns.

## 🛠️ Construído com

- [Golang](https://golang.org) - Linguagem de Programação
- [Asdf](https://asdf-vm.com) - usado para versionamento do Golang
