// pacote main
package main

// importando o fmt
import (
	"errors"
	"fmt"
	"time"
)

// essa funcao nao faz nada -> só mostra os comentários em cima da função
func funcao(){

}



// nome parâmetros tipo de retorno
func nomeDaFuncao(parametro int) (int, error){
	
	if parametro <= 0{
		return 0, errors.New("retornou zero ou abaixo")
	}
	return parametro, nil

}

// função para demonstrar ponteiros
func subtrai1(parametro *int){
	*parametro -= 1
}

// struct
type Pessoa struct{
	Nome string
	Idade int
}

// interface
type Falante interface{
	falar()
	falarAlto()
}
// pessoa está implementando a interface Falante
func (p Pessoa) falar(){
	fmt.Println("Olá")
}

// qualquer pessoa possui esta função! p.meApresentar()
func (p Pessoa) meApresentar(){
	fmt.Println("Olá, eu sou ", p.Nome, " tenho ", p.Idade, " anos")
}

func (p *Pessoa) fazerAniversario(){
	p.Idade += 1
}

func tempoPraExecutar (tempo time.Time, numero int, canal chan time.Duration){
	canal <- time.Since(tempo)
}


// declarando a função main
func main() {

	// joga isso na pilha (só roda no final do código)
	defer fmt.Println("rodei no final")


	/* Muito usado para SQL -> deu erro -> go back com defer
	defer func(){
		return
	} */

	// declarando variáveis
	stringVar := "string"
	_ = stringVar

	fmt.Println("Hello world")



	// if e else -> igual C
	/*
	if 2 > 1 {
		fmt.Println("2 é maior que 1")

	} else if 3 == 3{

	}
	*/

	// switch -> igual C
	switch stringVar {
	case "string":
		fmt.Println("string escrita")
	}

	// se colocar o mouse em cima disso, da para ver comentário
	funcao()

	// for
	for i := 0; i <= 10; i++{
		fmt.Println("Contei ", i)
	}
	i := 0

	// for (2)
	for{
		i++
		fmt.Println("Contei (2)", i)
		if i >= 10{
			break
		}
	}
	i = 0
	// for (3)
	for i <= 10{
		fmt.Println("Contei (3)", i)
		i++
	}
	// Não existe o WHILE -> tem que simular com o FOR

	for i := range 10{
		fmt.Println("contei (4) ", i)
	}

	inteiroDaFuncao, err := nomeDaFuncao(0)
	// se jogarmos um retorno para uma variável "-", o Go não apita erro pq ta jogando pro lixo
	
	fmt.Println(inteiroDaFuncao, err)

	if err != nil{
		fmt.Println("Deu erro: ", err)
	}else{
		fmt.Println(inteiroDaFuncao)
	}

	// o array define o tamanho
	var array [3]int
	// o slice não define o tamanho
	var slice []int

	// tamanho fixo -> 3
	fmt.Println(len(array))
	// tamanho dinâmico -> 0, 1, ...
	fmt.Println(len(slice))

	slice = append(slice, 5)
	slice = append(slice, 10)

	fmt.Println(array)
	fmt.Println(slice)

	fmt.Println("\n",len(array))
	fmt.Println(len(slice))

	// capacidade -> a capacidade do slide nunca vai ser estourada
	//				 pq ele sempre aumenta sua capacidade
	fmt.Println("\n",cap(array))
	fmt.Println(cap(slice))


	// tipo da chave (indíce)
	// tipo do valor
	var mapa = map[int]string{
		1: "numero 1",
		2: "numero 2",
	}
	fmt.Println(mapa)

	for indice, valor := range mapa{
		fmt.Println(indice, valor)
	}

	slice = append(slice, 20)
	slice = append(slice, 30)
	slice = append(slice, 40)

	fmt.Println("slice",slice)

	// inclui o primeiro até o penúltimo
	fmt.Println(slice[:3])

	fmt.Println(mapa)

	// serve para deletar alguma chave
	delete(mapa, 2)

	fmt.Println(mapa)

	// é possível verificar se existe aquele índice
	valor, existe := mapa[1]
	if !existe{
		fmt.Println("Chave inexistente")
	} else {
		fmt.Println(valor)
	}

	fmt.Println("\n\n\n\n\n")

	// ponteiros
		var ponteiro = 10
		subtrai1(&ponteiro)
		fmt.Println(ponteiro)

	fmt.Println("\n\n\n\n\n")


	// structs e métodos
		var eu Pessoa
		eu.Idade = 20
		eu.Nome = "Cauan"

		var Nao_eu Pessoa
		Nao_eu.Idade = 36
		Nao_eu.Nome = "Joao"

		fmt.Println(eu)

		switch Nao_eu{
		case Nao_eu:
			fmt.Println("Não sou eu!")
		}

		// atribuindo map a um struct
		var mapaPessoa = map[string]Pessoa{
			"111.111.111-11": eu,
		}
		fmt.Println(mapaPessoa)

		eu.meApresentar()

		eu.fazerAniversario()

		eu.meApresentar()

	fmt.Println("\n\n\n\n\n")


	// interfaces
		eu.falar()

	fmt.Println("\n\n\n\n\n")


	// goroutines

		go fmt.Println("print 1")
		go fmt.Println("print 2")
		go fmt.Println("print 3")
		go fmt.Println("print 4")
		go fmt.Println("print 5")
		go fmt.Println("print 6")
		go fmt.Println("print 7")
		time.Sleep(1 * time.Second)

	fmt.Println("\n\n\n\n\n")

	// channel
		
		var canal = make (chan time.Duration)
		tempoAgora := time.Now()
		go tempoPraExecutar(tempoAgora, 1, canal)
		fmt.Println(<-canal)
		fmt.Println(time.Since(tempoAgora).Nanoseconds())

}



// terminal:
// go run main.go no terminal para rodar
// go build . -> compila
// go mod init secomppGO -> guarda pacotes externos, como github (dependência)
// go get -u gorm.io/gorm -> gera a dependência do gorm
// go mod tidy -> remove as dependências q nn estão sendo usadas














