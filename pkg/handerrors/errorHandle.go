/*Error handling in Go*/
package handerrors

import (
	"errors"
	"fmt"
)

/*Go has type error in errors package which is used has Error value */

//To Create new error New function is used
func RetrunsError() (int, error) {

	return 0, errors.New("this my own error")
}

//error can also be created using fmt.Errorf function
func CreteFmtError() error {
	return fmt.Errorf("%s", "this error is created using fmt")
}

/*divide function divides two numbers and
return answer and error*/
func Divide(x, y int) (int, error) {

	if y == 0 {
		return 0, errors.New("cannot divide by 0")
	}
	return x / y, nil
}

func RunHandleError() {

	/*error handling in Go*/
	//check errro in if block
	if ans, err := Divide(4, 0); err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(ans)
	}
	v, err := Divide(4, 0)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(v)

	//to ignore errors
	a, _ := Divide(8, 4)
	fmt.Println(a)

}
