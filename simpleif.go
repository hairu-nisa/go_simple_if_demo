package main
import(
		"fmt"
		)
	func main(){
	fmt.Println("\nSimple if statment")
	fmt.Println("\n==================")
	fmt.Println("Enter a number:")
	var number int
	fmt.Scanln(&number)
	fmt.Println("\nEnter the number you are guessing")
	var guess int
	fmt.Scanln(&guess)
	fmt.Print("\nnumber and guess are ",number,guess)
	
		if guess < number{
				fmt.Println("\ntoo low")
				}
		if guess > number{
				fmt.Println("\ntoo high")
				}
		if guess == number{
				fmt.Println("\nyou got it")
				}
	}
