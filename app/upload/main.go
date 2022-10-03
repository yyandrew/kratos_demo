package main

import (
	"io"
	"log"
	"os"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/transport/http"
)

func uploadFile(ctx http.Context) error {
	req := ctx.Request()
	fileName := req.FormValue("name")
        name := req.FormValue("name")
        log.Print(name)
	file, handler, err := req.FormFile("file")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	f, err := os.OpenFile(handler.Filename, os.O_WRONLY|os.O_CREATE, 0o666)

	if err != nil {
		return err
	}

	defer f.Close()

	_, _ = io.Copy(f, file)

	return ctx.String(200, "File "+fileName+"Upload successfully")
}

func main() {
	httpSrv := http.NewServer(http.Address(":8000"))
	route := httpSrv.Route("/")
	route.POST("/upload", uploadFile)

	app := kratos.New(
		kratos.Name("upload"),
		kratos.Server(httpSrv),
	)

	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}
