package services

import (
	"github.com/citwild/wfe/api"
	"github.com/citwild/wfe/services/testserver"
	"testing"
)

func TestCreate_lg(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	t.Parallel()

	s := testserver.New()
	err := s.Start()
	if err != nil {
		t.Fatal(err)
	}
	defer s.Close()

	c, err := s.NewClient()
	if err != nil {
		t.Fatal(err)
	}

	_, err = c.Accounts.Create(s.Context, &api.NewAccount{Login: "me", Password: "pass", Email: "e@mail.com"})
	//if err != nil {
	//	t.Fatal(err)
	//}
	//
	//if a == nil {
	//	t.Errorf("expected created account != nil")
	//}
}
