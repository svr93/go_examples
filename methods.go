package main

import "fmt"

type User struct {

    Name string
    Surname string
}

func (user *User) ToString() string {

    return user.Name + " " + user.Surname
}

type BaseFloat float64

func (baseFloat BaseFloat) getAddress() *BaseFloat {

    return &baseFloat
}

func main() {

    man := User{ "Владимир", "Шамшеев" }
    woman := &User{ "Татьяна", "Николаева" }

    fmt.Println(man.ToString())
    fmt.Println(woman.ToString())

    f := BaseFloat(1.2)
    fmt.Println(f.getAddress())
}
