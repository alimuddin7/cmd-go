package main

import (
	"context"
	"log"

	"discord-webhook/grpc/client/protofiles/greetpb"
	"discord-webhook/grpc/client/protofiles/hashpb"
	"discord-webhook/grpc/client/protofiles/mathpb"
	"discord-webhook/grpc/client/protofiles/splitpb"

	"google.golang.org/grpc"
)

func main() {
	cc, err := grpc.Dial("localhost:8080", grpc.WithInsecure())

	if err != nil {
		panic(err)
	}

	defer cc.Close()

	c := greetpb.NewGreetServiceClient(cc)
	m := mathpb.NewMathServiceClient(cc)
	s := splitpb.NewSplitServiceClient(cc)
	h := hashpb.NewHashServiceClient(cc)

	getGreeting("Jack", "us", c)
	getGreeting("Jose", "mx", c)
	getGreeting("Fikri", "us", c)
	getGreeting("Alimuddin", "", c)
	mathOperation(1, 2, "+", m)
	mathOperation(10, 5, "-", m)
	split("Aku,Anak,Pintar", s)
	hash("1234", h)
}

func getGreeting(name, countryCode string, c greetpb.GreetServiceClient) {

	log.Println("creating greeting")

	res, err := c.Greet(context.Background(), &greetpb.GreetRequest{
		CountryCode: countryCode,
		UserName:    name,
	})

	if err != nil {
		log.Println("error: ", err)
		panic(err)
	}

	log.Println(res.Result)
}

func mathOperation(numberOne, numberTwo int64, operation string, c mathpb.MathServiceClient) {

	log.Println("Math Request")

	res, err := c.Math(context.Background(), &mathpb.MathRequest{
		NumberOne: numberOne,
		NumberTwo: numberTwo,
		Operation: operation,
	})

	if err != nil {
		log.Println("error: ", err)
		panic(err)
	}

	log.Println(res.Result)
}

func split(data string, c splitpb.SplitServiceClient) {
	log.Println("Split Request")

	res, err := c.Split(context.Background(), &splitpb.SplitReq{
		Data: data,
	})
	if err != nil {
		log.Println("error: ", err)
		panic(err)
	}

	log.Println(res.Result)

}

func hash(data string, c hashpb.HashServiceClient) {
	log.Println("hash Request")

	res, err := c.Hash(context.Background(), &hashpb.HashReq{
		Data: data,
	})
	if err != nil {
		log.Println("error: ", err)
		panic(err)
	}

	log.Println(res.Result)
}
