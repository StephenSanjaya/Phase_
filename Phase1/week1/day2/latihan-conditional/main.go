package main

import "fmt"

/*
    Convert Score to Alphabet

    - Jika value dari score lebih dari 100, maka print Invalid Score
    - Jika value dari score kurang dari 0, maka print Invalid Score
    - Jika value dari score berada diantara 90 hingga 100, print A
    - Jika value dari score berada diantara 80 hingga 89, print B
    - Jika value dari score berada diantara 70 hingga 79, print C
    - Jika value dari score berada diantara 60 hingga 69, print D
    - Jika value dari score berada dibawah 60, print E    
*/

func main()  {
	score := 40
	
	if score > 100 || score < 0 {
		fmt.Println("Invalid Score");
	}else{
		if(score >= 90 && score <= 100){
			fmt.Println("A");
		}else if score >= 80 && score <= 89 {
			fmt.Println("B");
		}else if score >= 70 && score <= 79 {
			fmt.Println("C");
		}else if score >= 60 && score <= 69 {
			fmt.Println("D");
		}else{
			fmt.Println("E");
		}
	}
}