package notify

import (
	"github.com/Blue-Berrys/Tiktok_e_commerce/rpc_gen/kitex_gen/email"
	"github.com/kr/pretty"
)

type NoopEmail struct {
}

func (e *NoopEmail) Send(req *email.EmailReq) error {
	_, _ = pretty.Printf("%v\n", req)
	return nil
}

func NewNoopEmail() NoopEmail {
	return NoopEmail{}
}
