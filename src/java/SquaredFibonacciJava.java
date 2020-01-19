package java;

class SquaredFibonacciJava{
//FiboInit Base array
  	private final int[] fiboInit = { 0, 1 }; // F0 = 0; F1 = 1
	private static int[] fiboArray = { 0, 1 };

	private static int[] appendToArray(final int[] existingArray, final int toAppend) {
		final int[] newArray = new int[existingArray.length + 1];
		for (final int i : existingArray) {
			newArray[i] = existingArray[i];
		}
		newArray[existingArray.length+1] = toAppend;
		return newArray;
	}

	/**
	 * Esta función calcula una serie Fibonacci dado un indice. La técnica usada para calcular la serie
	 * aprovecha recursión y los cálculos anteriores almacenados en un arreglo.
	 * El índice es n, el contenido es el valor de la serie Fibonacci para ese n.
	 * Se inicializan con los valores de n = 0 y n = 1 que definen a la serie.
	 * Cuando hay un valor asignado significa que hay un valor ya calculado.
	 * @param n El tamaño de la serie > 0
	 *  */
	private static int fibonacciNth( final int n)  {
		//fmt.Println("* Entrando a fibonacciNth n es " + strconv.Itoa(n) + " y el tamaño de precálculos es " + strconv.Itoa(len(fiboArray)))
		if (n < 0 ){
			throw new Error("No es posible calcular series con N negativas");
		} else if ((n >= 0) && (fiboArray.length - 1) >= n) { //ya inicié? lo calculé? (para n=0 y n=1 ya se "calcularon")
			//fmt.Println("Regresando número previamente calculado fiboArray[" + strconv.Itoa(n) + "]=" + strconv.Itoa(fiboArray[n]))
			return fiboArray[n];
		} else { // No se ha calculado antes, sumar los últimos 2 números y agregar un nuevo numero (FN-1 + Fn-2) al arreglo
			//fmt.Println("No se ha calculado antes. Se requiere nuevo cálculo con recursión fibonacciNth(" + strconv.Itoa(n-1) + ") + fibonacciNth(" + strconv.Itoa(n-2) + ")")
			fiboArray = appendToArray(fiboArray,(fibonacciNth(n-1) + fibonacciNth(n-2)));//Recursión para asignar un nuevo numero
			//fmt.Println("Se agregó nuevo número f[" + strconv.Itoa(n) + "]=" + strconv.Itoa(fiboArray[n]))
			//fmt.Println("fiboArray completo: ", fiboArray)
			return fiboArray[n];
		}
	}

/**
 *
 * Esta función genera el cuadrado del contenido del arreglo
 * @param anArray El arreglo a elevar al cuadrado
 * */
private static int[] squaredArray(final int[] anArray) {
	final int[] squaredSeries = new int[anArray.length];
	for (final int j : anArray) {
		squaredSeries[j] = anArray[j] * anArray[j];
	}
	return squaredSeries;
}

/**
 * Regresa un arreglo con la serie completa, al cuadrado, de la serie Fibonacci dado un N
 * @param n El numero para calcular la serie de Fibonacci
 */
private static int[] fibonacciNthArray(final int n) {
	//fmt.Println("Entrando fibonacciNthArray" + strconv.Itoa(n))
	//Init array as in FIBO_INIT
	//fiboArray := []int{0, 1}
	fibonacciNth(n);
	//fmt.Println("Terminando Invocacion fibonacciNth(" + strconv.Itoa(n) + ") con")
	//fmt.Println("fiboArray completo en fibonacciNthArray: ", fiboArray)
	return fiboArray;
}

/**
 * Regresa un arreglo con la serie completa, al cuadrado, de la serie Fibonacci dado un N
 * @param n El numero para calcular la serie de Fibonacci
 */
private static int[] squaredFibonacciNthArray(final int n) {
	// Init array as in FIBO_INIT
	// fiboArray = []int{0, 1}
		final int[] resultArray = fibonacciNthArray(n);
		return squaredArray(resultArray);
	}

	/**
	 * Regresa el numero factorial del numero dado
	 * 
	 * @param n El numero para calcular el factorial
	 */
	private int factorialN(final int n) {
	if (n < 0) {
		throw new Error("No es posible calcular con N negativo");
	} else if (n == 0) {
		return 1;
	} else {
		return (n * factorialN(n-1));
	}
}

/**
 * Regresa un arreglo numérico calculando el factorial de cada elemento del arreglo ingresado
 * @param arrayCalc arreglo a calcular
 */
int[] factorialArray(int[] arrayCalc) {
	int[] resultArray = new int[arrayCalc.length];
	for (int j : arrayCalc) {
		int val = arrayCalc[j];
		resultArray[j] = factorialN(val);
	}
	return resultArray;
}

/**
 * Regresa un arreglo factorial de una seria de Fibonacci
 * @param n
 */
int[] factorialFibonacciArray(int n)  {
	return factorialArray(fibonacciNthArray(n));
}

	// SquaredFibonacciGo es la función que será invocada como servicio en un
	// Function.
	// @w http.ResponseWriter el objeto para escribir la respuesta
	// @r http.Request Objeto para enviar respuesta
/* void SquaredFibonacciGo(final w http.ResponseWriter, r *http.Request) {
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
	fmt.Fprint(w, html.EscapeString("{'result':'"+returnArrayString+"'}"))

} */

public static void main(String[] args) {
	//Ejemplos
	final int N = 80;
	System.out.println("Iniciando con N = " + N);
	//fmt.Println("El número Fibonaccy para n=" + strconv.Itoa(N) + " es " + strconv.Itoa(fibonacciNth(N, nil)))
	//fmt.Println("La serie completa empezando con cero es : {" + arrayToString(fibonacciNthArray(N, nil), ",") + "}")
	System.out.println("La serie completa al cuadrado: {" + squaredFibonacciNthArray(N) + "}");
	//fmt.Println("El factorial de " + strconv.Itoa(N) + " es " + strconv.Itoa(factorialN(N, nil)))
	//fmt.Printn("El arreglo factorial: " + arrayToString(factorialFibonacciArray(N, nil), ","))
}
}