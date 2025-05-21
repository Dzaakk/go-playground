package main

import (
	"fmt"

	"auction-anonymizer/anonymizer"
)

func main() {
	emails := []string{"john.doe@example.com", "a@b.com"}
	usernames := []string{"auctioneer123", "xy"}
	ips := []string{"192.168.0.123", "10.0.0.77"}

	emailAnonymizer := anonymizer.EmailAnonymizer{}
	usernameAnonymizer := anonymizer.UsernameAnonymizer{}
	ipAnonymizer := anonymizer.IPAnonymizer{}

	fmt.Println("Email Anonymization:")
	for _, e := range emails {
		fmt.Printf("%s -> %s\n", e, emailAnonymizer.Anonymize(e))
	}

	fmt.Println("\nUsername Anonymization:")
	for _, u := range usernames {
		fmt.Printf("%s -> %s\n", u, usernameAnonymizer.Anonymize(u))
	}

	fmt.Println("\nIP Anonymization:")
	for _, ip := range ips {
		fmt.Printf("%s -> %s\n", ip, ipAnonymizer.Anonymize(ip))
	}
}
