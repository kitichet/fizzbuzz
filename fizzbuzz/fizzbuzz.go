package fizzbuzz
import (
	"strconv"
)
func Fizzbuzz(s string)string {
	i ,_:= strconv.Atoi(s)
	
	if(i%3 ==0 && i%5 ==0){
		 return "FizzBuzz"
	}else {
		if(i%3 ==0){
			return "Fizz"	
		} else if (i%5 == 0){
			return "Buzz"
		}
	}
	
	return s
}