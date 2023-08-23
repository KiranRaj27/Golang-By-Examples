package main 

import "fmt"

func main(){
	i:=1
	for i <=3{
		fmt.Println(i)
		i = i + 1
	}

	for j:=1;j<=9;j++{
		fmt.Println(j)
	}

	for {
		fmt.Println("Loop")
		break
	}

	for n:=0;n<=5;n++{
		if n % 2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}