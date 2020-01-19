/**
 * Este código para Node incluye algoritmos para determinar la serie de Fibonacci y aplicar
 * operaciones sobre las mismas series. Estas funciones son para ser expuestas mediante microservicios 
 * con Functions.
 * 
 * @author joseluisrg
 * 
 */

const FIBO_INIT =[0,1] //F0 = 0; F1 = 1
var fiboArray = [0,1]
/**
 * Esta funcióno calcula una serie Fibonacci dado un indice. La técnica usada para calcular la serie 
 * aprovecha recursión y los cálculos anteriores almacenados en un arreglo.
 * El índice es n, el contenido es el valor de la serie Fibonacci para ese n.
 * Se inicializan con los valores de n = 0 y n = 1 que definen a la serie.
 * Cuando hay un valor asignado significa que hay un valor ya calculado.
 * @param n El tamaño de la serie > 0
 *  */ 
function fibonacciNth(n){
    //console.log("n es " + n + " y el tamaño de precálculos es " + fiboArray.length);
    if (n < 0) {
        throw Error("No es posible calcular series con N negativas");
    } else if((n >= 0) && (fiboArray.length - 1  >= n )){ //ya inicié? lo calculé? (para n=0 y n=1 ya se "calcularon")
        //console.log("Regresando número previamente calculado fiboArray["+n+"]="+fiboArray[n]);
        return fiboArray[n]
    } else { // No se ha calculado antes, sumar los últimos 2 números y agregar un nuevo numero (FN-1 + Fn-2) al arreglo       
        fiboArray[n] = fibonacciNth(n-1) + fibonacciNth(n-2) //Recursión para asignar un nuevo numero
        //console.log("Se agregó nuevo número f["+n+"]=" + fiboArray[n] );
        return fiboArray[n] 
    }  
}

/**
 * 
 * Esta función genera el cuadrado del contenido del arreglo
 * @param anArray El arreglo a elevar al cuadrado
 * */
function squaredArray(anArray){
    var squaredSeries = [];
    for (const i in anArray) {
        squaredSeries[i] = anArray[i] * anArray[i]
    }
    return squaredSeries
}

/**
 * Regresa un arreglo con la serie completa, al cuadrado, de la serie Fibonacci dado un N
 * @param n El numero para calcular la serie de Fibonacci 
 */
function fibonacciNthArray(n){
    //Init array as in FIBO_INIT
    fiboArray= [0,1]
    fibonacciNth(n)
    return fiboArray
}

/**
 * Regresa un arreglo con la serie completa, al cuadrado, de la serie Fibonacci dado un N
 * @param n El numero para calcular la serie de Fibonacci 
 */
function squaredFibonacciNthArray(n){
    //Init array as in FIBO_INIT
    fiboArray= [0,1]
    fibonacciNth(n)
    return squaredArray(fiboArray)
}

/**
 * Regresa el numero factorial del numero dado
 * @param n El numero para calcular el factorial
 */
function factorialN(n){
    if (n < 0) {
        throw Error("No es posible calcular con N negativo");
    } else if(n == 0){
        return 1
    } else {
        return (n * factorialN(n-1));
    }
}

/**
 * Regresa un arreglo numérico calculando el factorial de cada elemento del arreglo ingresado
 * @param arrayCalc arreglo a calcular  
 */
function factorialArray(arrayCalc){
    var resultArray = []
    for (const i in arrayCalc) {
        const val = arrayCalc[i];
        resultArray[i] = factorialN(val)
    }
    return resultArray
}

/**
 * Regresa un arreglo factorial de una seria de Fibonacci
 * @param n
 */
function factorialFibonacci(n){
    return factorialArray(squaredFibonacciNthArray(n))
}

/**
 * Responde a request de HTTP.
 *
 * @param {!express:Request} req HTTP request context.
 * @param {!express:Response} res HTTP response context.
 */
exports.squaredFibonacciNode = (req, res) => {
    let nString = req.query.n || req.body.n || 'Indicar n como {"n":"<valor>"} en e request';
    var n = parseInt(nString)
    var result = squaredFibonacciNthArray(n)
    res.status(200).send(result);
  };
  

//Ejemplos
var N = 80
//console.log("El número Fibonaccy para n=" + N + " es " + fibonacciNth(N))
//console.log("La serie completa empezando con cero es : {" + fibonacciNthArray(N)+"}");
console.log("La serie completa al cuadrado: {" +squaredFibonacciNthArray(N) +"}")
//console.log("El factorial de " + N + " es " + factorialN(N))
//console.log("El arreglo factorial: " + factorialFibonacci(N))


