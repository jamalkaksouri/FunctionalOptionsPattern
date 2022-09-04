package main

import (
	approach2 "FunctionalOptionsPattern/Approach2"
	approach3 "FunctionalOptionsPattern/Approach3"
	approach1 "FunctionalOptionsPattern/Approache1"
	"FunctionalOptionsPattern/RealExample"
	"fmt"
	"log"
	"time"
)

func main() {
	svr := approach1.New("https://google.com", 8080)
	//svr2 := approach1.NewWithTimeout("https://google.com", 1234, time.Duration(time.Second))
	//svr3 := approach1.NewWithTimeoutAndMaxConn("https://google.com", 1234, time.Duration(time.Second), 20)
	if err := svr.Start(); err != nil {
		log.Fatal(err)
	}

	s := approach2.New(approach2.Config{
		Host:    "localhost",
		Port:    1234,
		Timeout: 30 * time.Second,
		MaxConn: 10,
	})
	if err := s.Start(); err != nil {
		log.Fatal(err)
	}

	svrApp3 := approach3.New(
		approach3.WithPort(8080),
		approach3.WithHost("https://google.com"),
		approach3.WithMaxConn(20),
		approach3.WithTimeout(time.Second*5000),
	)
	if err := svrApp3.Start(); err != nil {
		log.Fatal(err)
	}

	houseOp := RealExample.NewHouse(RealExample.WithFloors(5), RealExample.WithConcrete())
	fmt.Println(houseOp)

}
