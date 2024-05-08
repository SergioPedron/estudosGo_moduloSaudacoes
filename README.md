estudosGo_moduloSaudacoes

Estudos Golang - Exibe uma Saudação

Primeiro contato com Golang

Criação de um módulo simples para uso em outras aplicações para gerar uma saudação.

Funções exportadas:

- Ola(nome string) (string, error)
  * Recebe um nome e retorna uma mensagem aleatória de saudação

- OlaPessoas(nomes []string) (map[string]string, error)
  * Recebe um map de nomes e retorna um map de mensagens aleatórias de saudações para estes nomes
    
- OlaEm(nome string) (string, error)
  * Recebe um nome e retorna uma mensagem de saudação com a hora e data atual
    

Obs.: 
- Ambas as funções podem retornar algum erro.
- Contempla testes
