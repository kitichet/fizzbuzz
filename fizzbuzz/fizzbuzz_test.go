package fizzbuzz
import "testing"


func TestFizzbuzzOne(t *testing.T){
	result := Fizzbuzz("1")
	if(result != "1"){
		t.Fatalf("expect 1 but got" + result)
	}
}

func TestFizzbuzzTwo(t *testing.T){
	result := Fizzbuzz("2")
	if(result != "2"){
		t.Fatalf("expect 2 but got" + result)
	}
}

func TestFizzbuzzThree(t *testing.T){
	result := Fizzbuzz("3")
	if(result != "Fizz"){
		t.Fatalf("expect Fizz but got" + result)
	}
}

func TestFizzbuzzFive(t *testing.T){
	result := Fizzbuzz("5")
	if(result != "Buzz"){
		t.Fatalf("expect Buzz but got" + result)
	}
}

func TestFizzbuzzSix(t *testing.T){
	result := Fizzbuzz("6")
	if(result != "Fizz"){
		t.Fatalf("expect Fizz but got" + result)
	}
}

func TestFizzbuzzNine(t *testing.T){
	result := Fizzbuzz("9")
	if(result != "Fizz"){
		t.Fatalf("expect Fizz but got" + result)
	}
}

func TestFizzbuzzTen(t *testing.T){
	result := Fizzbuzz("10")
	if(result != "Buzz"){
		t.Fatalf("expect Buzz but got" + result)
	}
}

func TestFizzbuzzTwelve(t *testing.T){
	result := Fizzbuzz("12")
	if(result != "Fizz"){
		t.Fatalf("expect Fizz but got" + result)
	}
}


func TestFizzbuzzFifteen(t *testing.T){
	result := Fizzbuzz("15")
	if(result != "FizzBuzz"){
		t.Fatalf("expect FizzBuzz but got" + result)
	}
}

