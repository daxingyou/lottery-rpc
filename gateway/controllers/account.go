package controllers

import (
	"fmt"
	"net/http"

	"github.com/mholt/binding"
	"github.com/unrolled/render"
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

	// find all people in the database
	account := AccountForm{
		Name: "test",
	}

	r.JSON(res, 200, account)
}
