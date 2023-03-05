package main

import (
	"bufio"
	"doremi/handler"
	"fmt"
	"os"
	"strings"
)

func main() {
	cliArgs := os.Args[1:]

	if len(cliArgs) == 0 {
		fmt.Println("Please provide the input file path")
		return
	}

	filePath := cliArgs[0]
	file, err := os.Open(filePath)

	if err != nil {
		fmt.Println("Error opening the input file")
		return
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	if err != nil {
		fmt.Println("something wrong", err.Error())
		return
	}
	user := &handler.User{}
	var subscriptionHandler handler.SubscriptionHandler = user
	var planHandler handler.PlanHandler = user
	var TopUpHandler handler.TopUpHandler = user
	processCommand(scanner, subscriptionHandler, planHandler, TopUpHandler)
}

func processCommand(scanner *bufio.Scanner, subs handler.SubscriptionHandler, plan handler.PlanHandler, topUp handler.TopUpHandler) {
	for scanner.Scan() {
		args := scanner.Text()
		argList := strings.Fields(args)

		switch argList[0] {
		case "START_SUBSCRIPTION":
			_ = subs.StartSubscription(argList[1])
		case "ADD_SUBSCRIPTION":
			_ = plan.AddPlan(argList[1], argList[2])
		case "ADD_TOPUP":
			_ = topUp.AddTopUp(argList[1], argList[2])
		case "PRINT_RENEWAL_DETAILS":
			_ = subs.GetRenewalDetails()
			_ = subs.GetRenewalAmount()
		}
	}
}
