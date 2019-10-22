package function

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/red-gold/ts-serverless/src/core/utils"
)

type EmailInput struct {
	RefEmail string `json:"email"`
	Password string `json:"password"`
}

func Handle(w http.ResponseWriter, r *http.Request) {
	var input []byte

	if r.Body != nil {
		defer r.Body.Close()

		body, _ := ioutil.ReadAll(r.Body)

		input = body
	}

	var emailInput EmailInput

	if err := json.Unmarshal(input, &emailInput); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("Umarshaling %s", err.Error())))
	}
	email := utils.NewEmail(emailInput.RefEmail, emailInput.Password, "smtp.gmail.com")
	req := utils.NewEmailRequest([]string{"amir.gholzam@live.com"}, "test subject", "")
	templateData := struct {
		Name    string
		Code    string
		AppName string
	}{
		Name:    "Amir",
		Code:    "1293740",
		AppName: "Playground",
	}
	status, err := email.SendEmail(req, "tmpl.html", templateData)
	if err != nil {
		fmt.Printf("Error in email: %s", err)
	}

	fmt.Printf("Email status is %v", status)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("Hello world, input was: %s", string(input))))
}
