package main

import (
	"fmt"

	"github.com/learnGo/09/accounts"
)

func main() {
	account := accounts.NewAccount("coco")
	
	//&{coco 0} : 복사본이 아니라 object라는 이야기
	//fmt.Println(account) 
	
	account.Deposit(10) //account를 복사하는 것이 아니라, Deposit method를 호출한 account를 사용해라
	
	//error handling
	err := account.Withdraw(20)
	if err != nil {
		//log.Fatalln(err)
		fmt.Println(err)
	}

	//fmt.Println(account.Balance(), account.Owner())
	fmt.Println(account)


}
