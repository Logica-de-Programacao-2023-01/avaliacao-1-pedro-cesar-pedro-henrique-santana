package q5

import (
	"strings"
)

//Pedro começou a frequentar aulas de programação. Na primeira aula, sua tarefa foi escrever um programa simples. O
//programa deveria fazer o seguinte: na sequência de caracteres fornecida, composta por letras latinas maiúsculas e
//minúsculas, ele:
//
//* deleta todas as vogais;
//* insere um caractere "." antes de cada consoante;
//* substitui todas as consoantes maiúsculas pelas correspondentes em minúsculas.
//
//As vogais são as letras "A", "O", "E", "U", "I", e o restante são consoantes. A entrada do programa é exatamente uma
//sequência de caracteres e a saída deve ser uma única sequência de caracteres, resultante após o processamento do
//programa na sequência de caracteres inicial.
//
//Ajude Pedro a lidar com esta tarefa fácil.

func ProcessString(s string) string {
	s = strings.ToLower(s)
	//consoante := []string{"b", "c", "d", "f", "g", "h", "j", "k", "l", "m", "n", "p", "q", "r", "s", "t", "v", "x", "w", "y", "z"}

	s = strings.ReplaceAll(s, "a", "")
	s = strings.ReplaceAll(s, "e", "")
	s = strings.ReplaceAll(s, "i", "")
	s = strings.ReplaceAll(s, "o", "")
	s = strings.ReplaceAll(s, "u", "")

	s = strings.ReplaceAll(s, "b", ".b")
	s = strings.ReplaceAll(s, "c", ".c")
	s = strings.ReplaceAll(s, "d", ".d")
	s = strings.ReplaceAll(s, "f", ".f")
	s = strings.ReplaceAll(s, "g", ".g")
	s = strings.ReplaceAll(s, "h", ".h")
	s = strings.ReplaceAll(s, "j", ".j")
	s = strings.ReplaceAll(s, "k", ".k")
	s = strings.ReplaceAll(s, "l", ".l")
	s = strings.ReplaceAll(s, "m", ".m")
	s = strings.ReplaceAll(s, "n", ".n")
	s = strings.ReplaceAll(s, "p", ".p")
	s = strings.ReplaceAll(s, "q", ".q")
	s = strings.ReplaceAll(s, "r", ".r")
	s = strings.ReplaceAll(s, "s", ".s")
	s = strings.ReplaceAll(s, "t", ".t")
	s = strings.ReplaceAll(s, "v", ".v")
	s = strings.ReplaceAll(s, "w", ".w")
	s = strings.ReplaceAll(s, "x", ".x")
	s = strings.ReplaceAll(s, "y", ".y")
	s = strings.ReplaceAll(s, "z", ".z")

	return s
}
