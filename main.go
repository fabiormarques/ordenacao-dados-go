package main

import (
	"cmp"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func recoverCustom() {
	r := recover()
	if r != nil {
		fmt.Println(r)
	}
}

type Pessoa struct {
	nome      string
	idade     int
	pontuacao int
}

func readContentInputFile(fileName string) string {
	file, err := os.ReadFile(fileName)
	check(err)

	return string(file)
}

func structuredData(data string) []Pessoa {
	pessoas := []Pessoa{}

	for idx, line := range strings.Split(data, "\n") {
		if idx != 0 && len(line) > 0 {
			dados := strings.Split(line, ",")

			idade, _ := strconv.Atoi(dados[1])
			pontuacao, _ := strconv.Atoi(dados[2])
			pessoas = append(pessoas, Pessoa{
				nome:      dados[0],
				idade:     idade,
				pontuacao: pontuacao,
			})
		}
	}

	return pessoas
}

func ordering(pessoas []Pessoa) []Pessoa {
	slices.SortFunc(pessoas,
		func(p1, p2 Pessoa) int {
			pess1 := p1.nome + strconv.Itoa(p1.idade)
			pess2 := p2.nome + strconv.Itoa(p2.idade)
			return cmp.Compare(pess1, pess2)
		})

	return pessoas
}

func writeOutput(pessoas []Pessoa, fileName string) {
	fileOutput, err := os.Create(fileName)
	check(err)
	defer fileOutput.Close()

	saida := []string{"Nome,Idade,Pontuação\n"}

	for _, linha := range pessoas {
		saida = append(saida, linha.nome+","+strconv.Itoa(linha.idade)+","+strconv.Itoa(linha.pontuacao)+"\n")
	}
	fmt.Fprintln(fileOutput, strings.Join(saida, ""))
}

func main() {
	defer recoverCustom()

	if len(os.Args) != 3 {
		panic("Informe os arquivos de entrada e saída de dados!\nExemplo: go run main.go <arquivo1> <arquivo2>")
	}

	data := readContentInputFile(os.Args[1])

	pessoas := structuredData(data)

	pessoas = ordering(pessoas)

	writeOutput(pessoas, os.Args[2])
}
