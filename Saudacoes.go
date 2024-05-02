package moduloSaudacoes

import (
	"errors"
	"fmt"
	"math/rand"
)

// Recebe um nome e retorna uma mensagem de saudação.  Se nome vier em branco, retorna um erro no segundo parâmetro
func Ola(nome string) (string, error) {
	// Função com inicial em maíuscula pode ser exportada, ou seja utilizada em outro pacote //
	// Tratamento de erros no golang é no retorno da função. Neste caso retorna dois parâmetros:  Nome e um erro se for o caso
	if nome == "" {
		return "", errors.New("Infome um nome")
	}
	// O operador := declara e inicializa uma variável conforme o tipo recebido, neste caso uma string
	mensagem := fmt.Sprintf(retornaMensagemAleatoria(), nome)
	return mensagem, nil
}

//--------------------------------------------------------------------------------------------------------------------------------------------------------------//

func OlaPessoas(nomes []string) (map[string]string, error) {
	// Maps associam uma chave a um valor. Possui a seguinte estrutura: map[TipoChave]TipoValor
	// O make é utilizado para inicializar para que depois seja possível modificar.
	mensagens := make(map[string]string)
	// Agora deve-se percorrer o map de nomes recebido por parâmetro, passar cada nome para a função Ola para que retorne a mensagem aleatória armazenamdo no map de mensagens
	// Como não vamos utilizar o índice do for, substituímos pelo identificador em brando do go, o _,
	for _, nome := range nomes {
		mensagem, erro := Ola(nome)
		if erro != nil {
			return nil, erro
		}
		mensagens[nome] = mensagem
	}
	// após ler os nomes, retornamos um map com o par de chave/valor igual a nome/mensagem
	return mensagens, nil
}

//--------------------------------------------------------------------------------------------------------------------------------------------------------------//

// Retorna uma mensagem aleatória conforme mensagens adicionadas em um slice, que é semelhente a um array porém com tamanho dinâmico
// Função inicia com letra minúscula pois não é exportada para fora do pacote
func retornaMensagemAleatoria() string {
	// Um slice com mensagens
	modelosMensagens := []string{
		"Olá, %v. Bem Vindo",
		"Saudações %v",
		"Oi %v, que bom ver você por aqui",
	}
	return modelosMensagens[rand.Intn(len(modelosMensagens))]
}

//--------------------------------------------------------------------------------------------------------------------------------------------------------------//
