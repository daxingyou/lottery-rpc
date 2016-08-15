package controllers

import (
	"fmt"
	"net/http"

	"github.com/mholt/binding"
	"github.com/micro/go-micro/client"
	hello "github.com/micro/micro/examples/greeter/server/proto/hello"
	"github.com/unrolled/render"
	"golang.org/x/net/context"
)

type Account struct{}

type AccountForm struct {
	//	ID         string
	Name string
	//	Blance     float32
	//	Block      bool
	//	Createtime time.Time
	//	Updatetime time.Time
}

func (af *AccountForm) FieldMap(r *http.Request) binding.FieldMap {
	return binding.FieldMap{
		//		&af.ID: "id",
		&af.Name: "name",
	}
}

func (a Account) Create(w http.ResponseWriter, r *http.Request) {
	accountForm := new(AccountForm)
	errs := binding.Bind(r, accountForm)
	if errs.Handle(w) {
		return
	}
	fmt.Fprintf(w, "Create account: %s\n", accountForm.Name)
}

func (a Account) Index(res http.ResponseWriter, req *http.Request) {
	r := render.New(render.Options{})

	greeter := hello.NewSayClient("go.micro.srv.greeter", client.DefaultClient)

	// request the Hello method on the Greeter handler
	rsp, err := greeter.Hello(context.TODO(), &hello.Request{
		Name: "John",
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	// find all people in the database
	account := AccountForm{
		Name: rsp.Msg,
	}

	r.JSON(res, 200, account)
}
