moduloSaudacoes
Primeiro contato com a linguagem GO

Criação de um módulo simples para uso em outras aplicações para gerar uma saudação.

Módulo: moduloSaudacoes Funções exportadas:

Ola(nome string) (string, error) Recebe um nome e retorna uma mensagem aleatória de saudações

OlaPessoas(nomes []string) (map[string]string, error) Recebe um map de nomes e retorna um map de mensagens aleatórias de saudações para estes nomes

Obs.: Ambas as funções podem retornar algum erro.
