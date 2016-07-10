package store

import (
	"golang.org/x/net/context"
	"github.com/wide-field-ethnography/wfe/api/wfe"
)

type Accounts interface {
	Create(ctx context.Context, newUser *wfe.User, email *wfe.EmailAddr) (*wfe.User, error)
}
