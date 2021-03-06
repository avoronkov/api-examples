// Copyright 1999-2014. Parallels IP Holdings GmbH. All Rights Reserved.
package main

import (
	"fmt"
	"log"
	"os"
)

const request = `
<packet>
	<server>
		<get_protos/>
	</server>
</packet>
`

func main() {
	host := os.Getenv("REMOTE_HOST")
	login := os.Getenv("REMOTE_LOGIN")
	if login == "" { login = "admin" }
	password := os.Getenv("REMOTE_PASSWORD")

	cli := NewPleskApiClient(host)
	cli.SetCredentials(login, password)

	// make work with self-signed sertificate:
	// cli.InsecureSkipVerify = true

	responseData, err := cli.Request(request)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(responseData)
}
