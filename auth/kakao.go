package auth

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/pkg/errors"

	"golang.org/x/crypto/bcrypt"

	"github.com/shjp/shjp-server/constant"
	"github.com/shjp/shjp-server/models"
	"github.com/shjp/shjp-server/utils"
)

type kakaoLoginRequest struct {
	AccountID    string  `json:"accountId"`
	ClientID     string  `json:"clientId"`
	ProfileImage *string `json:"profileImage"`
	Nickname     *string `json:"nickname"`
}

// HandleKakaoLogin handles kakao login request
func HandleKakaoLogin(w http.ResponseWriter, r *http.Request) {
	log.Println("kakao login handler")
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("Error reading body: %s\n", err)
		return
	}
	fmt.Printf("body: %s\n", string(body))

	var requestPayload kakaoLoginRequest
	err = json.Unmarshal(body, &requestPayload)
	if err != nil {
		fmt.Printf("Error deserializing login body: %s\n", err)
		return
	}

	m := models.NewMember()
	m.AccountType = constant.Kakao
	bcrypted, err := bcrypt.GenerateFromPassword([]byte(fmt.Sprintf("kakao:%s", requestPayload.AccountID)), bcrypt.MinCost)
	if err != nil {
		log.Printf("Could not generate hash from password, %s\n", err)
		utils.SendErrorResponse(w, errors.Wrap(err, "could not generate hash from password"), 500)
		return
	}
	m.AccountHash = string(bcrypted)

	if err = m.FindMe(); err != nil {
		log.Printf("Could not find the matching account")
		utils.SendErrorResponse(w, errors.Wrap(err, "could not verify the credentials"), 401)
		return
	}

	token, err := saveToSession(&m)
	if err != nil {
		log.Printf("Failed saving to session")
		utils.SendErrorResponse(w, errors.Wrap(err, "failed saving to session"), 500)
		return
	}

	utils.SendResponse(w, fmt.Sprintf(`{"token": %s}`, token), 200)
}
