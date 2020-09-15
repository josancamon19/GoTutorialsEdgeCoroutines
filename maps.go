package main

import "fmt"

type Service interface {
	sayHi()
}

type Service1 struct {
}

type Service2 struct {
}

func (s Service1) sayHi() {
	fmt.Println("Hi from service type 1")
}

func (s Service2) sayHi() {
	fmt.Println("Hi from service type 2")
}

func main1() {
	mapExample := make(map[string]int)
	mapExample["testKey"] = 2
	mapExample["testKe1"] = 3
	mapExample["testKe2"] = 4

	for key, value := range mapExample {
		fmt.Printf("%s: %d\n", key, value)
	}

	interfacesExample := make(map[string]Service)
	interfacesExample["s1"] = Service1{}
	interfacesExample["s2"] = Service2{}

	for key, value := range interfacesExample {
		fmt.Printf("%s: %d\n", key, value)
	}


}
