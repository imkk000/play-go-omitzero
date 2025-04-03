package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func main() {
	a := A{
		ZeroPID: primitive.NilObjectID,
		ZeroC: C{
			EmptyB: 0,
			ZeroB:  0,
		},
	}
	content, err := json.Marshal(a)
	if err != nil {
		log.Fatal().Err(err).Msg("marshal json")
	}

	fmt.Printf("result: %s\n", content)
}

type A struct {
	EmptyPID     primitive.ObjectID `json:",omitempty"` // show
	ZeroPID      primitive.ObjectID `json:",omitzero"`  // no show
	EmptyTime    time.Time          `json:",omitempty"` // show
	ZeroTime     time.Time          `json:",omitzero"`  // no show
	EmptyB       B                  `json:",omitempty"` // show
	ZeroB        B                  `json:",omitzero"`  // no show
	EmptyPtrB    *B                 `json:",omitempty"` // no show
	ZeroPtrB     *B                 `json:",omitzero"`  // no show
	EmptyBoolean bool               `json:",omitempty"` // no show
	ZeroBoolean  bool               `json:",omitzero"`  // no show
	EmptyC       C                  `json:",omitempty"` // show, nested no show
	ZeroC        C                  `json:",omitzero"`  // no show
}

type B struct{}

type C struct {
	EmptyB int `json:",omitempty"`
	ZeroB  int `json:",omitzero"`
}
