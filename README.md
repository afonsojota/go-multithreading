# Desafio Go: Multithreading

## Como Executar

### Pré-requisitos

- [Go](https://golang.org/doc/install) (versão 1.16 ou superior)

### Como Executar o Programa

1. **Clonando o Repositório:**
    - Abra um terminal e navegue até o diretório onde você deseja clonar o repositório.
    - Execute o seguinte comando para clonar o repositório:
      ```bash
      git clone https://github.com/afonsojota/go-multithreading.git
2. **Passos para Execução:**
    - Abra um terminal e navegue até o diretório onde você salvou os arquivos do projeto e acesse a pasta Multithreading.


3. **Executando o Programa:**
    - No terminal, execute o seguinte comando para compilar e executar o programa:
      ```bash
      go run main.go
      ```

4. **Resultados Esperados:**
    - O programa fará as requisições simultâneas às duas APIs de consulta de CEP.
    - Exibirá no terminal o endereço retornado e apontará a API que respondeu mais rapido e a API que respondeu mais devagar.
    - Caso uma ou ambas as requisições ultrapassem o limite de tempo (1 segundo), será exibida uma mensagem de timeout.
