package accounts

import (
	"errors"
	"fmt"
)

//Account struct
type Account struct {
	//소문자인 경우 private으로 외부에서 접근이 불가.
	owner   string
	balance int
}

var errNoMoney = errors.New("can't withdraw")


// NewAccount creates Account
func NewAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}


	//복사 후 새로운 object를 return하고 싶은 데, 
	//또 새로운 object를 만들긴 싫으니 복사본 자체를 return.
	return &account
}

//method
//(a Account) : receiver -> GO 에서는 method랑 동등한 관계
// a의 타입은 Account, struct의 첫글자를 따서 소문자로 해야 함.

//Deposit x amount on your account
func (a *Account) Deposit(amount int){
	//fmt.Println("Gonna deposit", amount)
	a.balance += amount
}

////////////////////////////////////////////////////////////
//receiver : GO 에서는 method랑 동등한 관계
//method는 struct 안 에 같이 작성되는 것이 아니라,
//func를 만들고 func는 struct를 가진다.
//////////////////////////////////////////////////////////////

// Balance of your account
func (a Account) Balance() int {
	return a.balance
}

// Wtihdraw x amount from your account
func (a *Account) Withdraw(amount int) error {

	if a.balance < amount {
		//return errors.New("Can't withdraw you are poor")
		return errNoMoney
	} 
	a.balance -= amount

	return nil
}

//ChangeOwner of the account
func (a *Account) ChangeOwner(newOwner string) {
	a.owner = newOwner
}

//Owner of the account
func (a Account) Owner() string {
	return a.owner
}

//fmt.Println(account) 사용시 결과 값이 &{coco 0} 였으나,
//아래를 작성해주니 결과값이 coco's account. Has: 10 로 나옴.
//GO 가 자동으로 method를 호출
func (a Account) String() string{
	return fmt.Sprint(a.Owner(),"'s account.\nHas: ", a.Balance())
}