package main

import (
	"github.com/srleyva/CertAPI/pkg/api"
	"net/http"
	"fmt"
)

func main()  {
	http.ListenAndServe(":8080", api.NewPkiServer())
	fmt.Println("Server Started and running at :8080")
}
