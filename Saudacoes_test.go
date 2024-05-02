package moduloSaudacoes

import (
	"regexp"
	"testing"
)

// TestOlaNome chama moduloSaudacoes.Ola com um nome, verificando um valor de retorno válido
func TestOlaNome(t *testing.T) {
	nome := "Luis"
	want := regexp.MustCompile(`\b` + nome + `\b`)
	msg, err := Ola("Luis")
	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Ola("Luis") retornou %q.  Deveria ter retornado string com o nome %v. Erro: %v`, msg, nome, err)
	}
}

// TestOlaEmpty chama  moduloSaudacoes.Ola com uma string vazia, para conferir o tratamento de erros para não permitir nome em branco
func TestOlaEmpty(t *testing.T) {
	msg, err := Ola("")
	if msg != "" || err == nil {
		t.Fatalf(`Ola("") retornou %q sem o nome preenchido. Erro: %v`, msg, err)
	}
}
