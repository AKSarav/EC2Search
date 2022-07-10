package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
)

var name string
var stdout []byte
var err2 error
var err1 error

func main() {
	argsWithoutProg := os.Args[1:]

	if len(argsWithoutProg) < 1 {
		fmt.Println("Need atleast one input")
		fmt.Println("\nTo use the specific named profile\n-------------------------------")
		fmt.Println("./ec2list SEARCHSTRING PROFILE")
		fmt.Println("\nTo use the default profile configured\n-------------------------------")
		fmt.Println("./ec2list SEARCHSTRING ")
		os.Exit(1)
	}

	if len(argsWithoutProg) == 2 {
		cmd := exec.Command("aws", "ec2", "describe-instances", "--query", "Reservations[*].Instances[*].{PublicIP:PublicIpAddress,PrivateIP:PrivateIpAddress,Name:Tags[?Key=='Name'].Value|[0],Status:State.Name,VpcId:VpcId,InstanceID:InstanceId,InstanceType:InstanceType}", "--filters", "Name=instance-state-name,Values=running", "Name=tag:Name,Values=*"+argsWithoutProg[0]+"*", "--profile", argsWithoutProg[1], "--output", "table")
		stdout, err2 = cmd.Output()
		if err2 != nil {
			log.Fatal(err2)
		}
	}

	if len(argsWithoutProg) == 1 {
		cmd := exec.Command("aws", "ec2", "describe-instances", "--query", "Reservations[*].Instances[*].{PublicIP:PublicIpAddress,PrivateIP:PrivateIpAddress,Name:Tags[?Key=='Name'].Value|[0],Status:State.Name,VpcId:VpcId,InstanceID:InstanceId,InstanceType:InstanceType}", "--filters", "Name=instance-state-name,Values=running", "Name=tag:Name,Values=*"+argsWithoutProg[0]+"*", "--output", "table")
		stdout, err1 = cmd.Output()
		if err1 != nil {
			log.Fatal(err1)
		}
	}

	stscmd := exec.Command("aws", "sts", "get-caller-identity", "--output", "json")
	stsout, errsts := stscmd.Output()

	if errsts != nil {
		log.Fatal(errsts)
	}

	fmt.Println("# EC2Search")
	fmt.Println("# https://gritfy.com/ec2search")

	type stsobj struct {
		UserID  string `json:"UserId"`
		Account string `json:"Account"`
		Arn     string `json:"Arn"`
	}

	var stsdata stsobj
	json.Unmarshal(stsout, &stsdata)

	fmt.Println("#", stsdata.Arn)
	fmt.Println(string(stdout))

}
