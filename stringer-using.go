package main

import (
    "fmt"
    "strconv"
)

type Driver struct {

    Name string
    Surname string
    Experience int
}

func (driver *Driver) String() string {

    expStr := "стаж (полных лет): " + strconv.Itoa(driver.Experience)

    return "Водитель: " + driver.Name + " " + driver.Surname + "; " + expStr
}

func main() {

    driver := &Driver{ "Владимир", "Шамшеев", 1 }

    fmt.Println(driver)
}
