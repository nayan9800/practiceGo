/*package is define using package keyword following its name*/

package arith

/*In go there no private , public and protected keywords
to Export a identifier it's name should start with Capital letter
for example Abs funcition
  structres,structres's method,structres's field,variable,constant
and function can be exported from package
*/

/*Euler's Constant*/
const Euler float64 = 0.5772156649

/*Abs function returns Absoulate value of given integer*/
func Abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

/*Add function adds given two number*/
func Add(a, b int) int { return a + b }

/*Substract function substract given two number*/
func Substract(a, b int) int { return a - b }

/*Multiply function multiply given two number*/
func Multiply(a, b int) int { return a * b }

/*Divide function divide given two number*/
func Divide(a, b int) int { return a / b }

/*Sum function return sum of all given numbers*/
func Sum(a ...int) (sum int) {

	for _, v := range a {
		sum = sum + v
	}
	return
}

/*Product return product of all given numbers*/
func Product(a ...int) (product int) {

	for _, v := range a {
		product = product * v
	}
	return
}
