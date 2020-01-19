// Package p contains an HTTP Cloud Function.
package p

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

var (
	//FiboInit Base array
	fiboInit  []int = []int{0, 1} //F0 = 0; F1 = 1
	fiboArray []int = []int{0, 1}
	buf       bytes.Buffer
	logger    = log.New(&buf, "INFO: ", log.Lshortfile)
	infof     = func(info string) {
		logger.Output(3, info)
	}
	arrayToString = func(A []int, delim string) string {
		var buffer bytes.Buffer
		for i := 0; i < len(A); i++ {
			buffer.WriteString(strconv.Itoa(A[i]))
			if i != len(A)-1 {
				buffer.WriteString(delim)
			}
		}
		return buffer.String()
	}
)

/**
 * Esta funcióno calcula una serie Fibonacci dado un indice. La técnica usada para calcular la serie
 * aprovecha recursión y los cálculos anteriores almacenados en un arreglo.
 * El índice es n, el contenido es el valor de la serie Fibonacci para ese n.
 * Se inicializan con los valores de n = 0 y n = 1 que definen a la serie.
 * Cuando hay un valor asignado significa que hay un valor ya calculado.
 * @param n El tamaño de la serie > 0
 *  */
func fibonacciNth(n int, err *error) int {
	//fmt.Println("* Entrando a fibonacciNth n es " + strconv.Itoa(n) + " y el tamaño de precálculos es " + strconv.Itoa(len(fiboArray)))
	if n < 0 {
		*err = errors.New("No es posible calcular series con N negativas")
		return -1
	} else if (n >= 0) && ((len(fiboArray) - 1) >= n) { //ya inicié? lo calculé? (para n=0 y n=1 ya se "calcularon")
		//fmt.Println("Regresando número previamente calculado fiboArray[" + strconv.Itoa(n) + "]=" + strconv.Itoa(fiboArray[n]))
		return fiboArray[n]
	} else { // No se ha calculado antes, sumar los últimos 2 números y agregar un nuevo numero (FN-1 + Fn-2) al arreglo
		//fmt.Println("No se ha calculado antes. Se requiere nuevo cálculo con recursión fibonacciNth(" + strconv.Itoa(n-1) + ") + fibonacciNth(" + strconv.Itoa(n-2) + ")")
		fiboArray = append(fiboArray, (fibonacciNth(n-1, nil) + fibonacciNth(n-2, nil))) //Recursión para asignar un nuevo numero
		//fmt.Println("Se agregó nuevo número f[" + strconv.Itoa(n) + "]=" + strconv.Itoa(fiboArray[n]))
		//fmt.Println("fiboArray completo: ", fiboArray)
		return fiboArray[n]
	}
}

/**
 *
 * Esta función genera el cuadrado del contenido del arreglo
 * @param anArray El arreglo a elevar al cuadrado
 * */
func squaredArray(anArray []int) []int {
	squaredSeries := make([]int, len(anArray))

	for i := 0; i < len(anArray); i++ {
		squaredSeries[i] = anArray[i] * anArray[i]
	}
	return squaredSeries
}

/**
 * Regresa un arreglo con la serie completa, al cuadrado, de la serie Fibonacci dado un N
 * @param n El numero para calcular la serie de Fibonacci
 */
func fibonacciNthArray(n int, err *error) []int {
	//fmt.Println("Entrando fibonacciNthArray" + strconv.Itoa(n))
	//Init array as in FIBO_INIT
	//fiboArray := []int{0, 1}
	fibonacciNth(n, err)
	//fmt.Println("Terminando Invocacion fibonacciNth(" + strconv.Itoa(n) + ") con")
	//fmt.Println("fiboArray completo en fibonacciNthArray: ", fiboArray)
	return fiboArray
}

/**
 * Regresa un arreglo con la serie completa, al cuadrado, de la serie Fibonacci dado un N
 * @param n El numero para calcular la serie de Fibonacci
 */
func squaredFibonacciNthArray(n int, err *error) []int {
	//Init array as in FIBO_INIT
	//fiboArray = []int{0, 1}
	resultArray := fibonacciNthArray(n, err)
	return squaredArray(resultArray)
}

/**
 * Regresa el numero factorial del numero dado
 * @param n El numero para calcular el factorial
 */
func factorialN(n int, err *error) int {
	if n < 0 {
		*err = errors.New("No es posible calcular con N negativo")
		return -1
	} else if n == 0 {
		return 1
	} else {
		return (n * factorialN(n-1, err))
	}
}

/**
 * Regresa un arreglo numérico calculando el factorial de cada elemento del arreglo ingresado
 * @param arrayCalc arreglo a calcular
 */
func factorialArray(arrayCalc []int, err *error) []int {
	var resultArray = make([]int, len(arrayCalc))
	for i := 0; i < len(arrayCalc); i++ {
		val := arrayCalc[i]
		resultArray[i] = factorialN(val, err)
	}
	return resultArray
}

/**
 * Regresa un arreglo factorial de una seria de Fibonacci
 * @param n
 */
func factorialFibonacciArray(n int, err *error) []int {
	return factorialArray(fibonacciNthArray(n, err), err)
}

// SquaredFibonacciGo es la función que será invocada como servicio en un Function.
// @w http.ResponseWriter el objeto para escribir la respuesta
// @r http.Request Objeto para enviar respuesta
func SquaredFibonacciGo(w http.ResponseWriter, r *http.Request) {
	var d struct {
		Message string `json:"n"`
	}
	var integerReceived int
	var err error
	var returnArrayString string
	if err := json.NewDecoder(r.Body).Decode(&d); err != nil {
		fmt.Fprint(w, "Para calcular el arreglo al cuadrado de Fibonacci envía una cadena json a este servicio como con un objeto 'message':'<entero>'")
		return
	}
	if d.Message == "" {
		fmt.Fprint(w, "Para calcular el arreglo al cuadrado de Fibonacci envía una cadena json a este servicio como con un objeto 'message':'<entero>'")
		return
	}
	integerReceived, err = strconv.Atoi(d.Message)
	if err != nil {
		fmt.Fprint(w, "Error en formato. Para calcular el arreglo al cuadrado de Fibonacci envía una cadena json a este servicio como con un objeto 'message':'<entero>'")
		return
	}
	fmt.Println("Input received parsed as int: %i", integerReceived)
	returnArrayString = "{" + arrayToString(squaredFibonacciNthArray(integerReceived, nil), ",") + "}"
	fmt.Println("Returning array " + returnArrayString)
	fmt.Fprint(w, returnArrayString)

}

func main() {
	//Ejemplos
	var N = 80
	fmt.Println("Iniciando con N = " + strconv.Itoa(N))
	//fmt.Println("El número Fibonaccy para n=" + strconv.Itoa(N) + " es " + strconv.Itoa(fibonacciNth(N, nil)))
	//fmt.Println("La serie completa empezando con cero es : {" + arrayToString(fibonacciNthArray(N, nil), ",") + "}")
	fmt.Println("La serie completa al cuadrado: {" + arrayToString(squaredFibonacciNthArray(N, nil), ",") + "}")
	//fmt.Println("El factorial de " + strconv.Itoa(N) + " es " + strconv.Itoa(factorialN(N, nil)))
	//fmt.Println("El arreglo factorial: " + arrayToString(factorialFibonacciArray(N, nil), ","))
}
