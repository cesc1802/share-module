package appjwt

import (
	"testing"

	"github.com/cesc1802/share-module/tokprovider"
)

func TestJwtProvider(t *testing.T) {
	jwtProvider := NewJwtProvider("secret-key")

	p := tokprovider.AppPayload{
		UserID: "1234",
	}

	appToken, err := jwtProvider.Generate(&p, 1000)
	if err != nil {
		t.Logf("something went wrong: %v", err)
		return
	}

	pay, err := jwtProvider.Extract(appToken.Token)
	if err != nil {
		t.Logf("something went wrong: %v", err)
		return
	}

	t.Log("generate token successfully: ", pay.GetUserID())
}
