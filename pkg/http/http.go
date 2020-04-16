package http

import (
	h "net/http"

	"github.com/gorilla/mux"
	"github.com/microapis/transbank-sdk-golang/pkg/webpay"
)

// Response ...
type Response struct {
	Data  interface{} `json:"data,omitempty"`
	Meta  interface{} `json:"meta,omitempty"`
	Error interface{} `json:"error,omitempty"`
}

type Services struct {
	PlusNormal *webpay.PlusNormal
	Patpass *webpay.Patpass
}

type plusNormalContext struct {
	PlusNormal *webpay.PlusNormal
}

type patpassContext struct {
	Patpass *webpay.Patpass
}

// Routes ...
func Routes(r *mux.Router, s Services) {
	s := r.PathPrefix("/api/v1").Subrouter()

	// check if is PlusNormal service is defined
	if s.PlusNormal != nil {
		// define context
		ctx := plusNormalContext{PlusNormal: s.PlusNormal}

		// POST /api/v1/plus-normal/transactions
		s.HandleFunc("/plus-normal/transactions", plusNormalInitTransaction(ctx)).Methods(h.MethodPost, h.MethodOptions)
	
		// GET /api/v1/plus-normal/transactions/:token
		s.HandleFunc("/plus-normal/transactions/{token}", plusNormalResultTransaction(ctx)).Methods(h.MethodGet, h.MethodOptions)
	}

	// check if is Patpass service is defined
	if s.Patpass != nil {
		// define context
		ctx := patpassContext{Patpass: s.Patpass}

		// POST /api/v1/patpass/transactions
		s.HandleFunc("/patpass/transactions", patpassInitTransaction(ctx)).Methods(h.MethodPost, h.MethodOptions)

		// GET /api/v1/patpass/transactions/:token
		s.HandleFunc("/patpass/transactions/{token}", patpassResultTransaction(ctx)).Methods(h.MethodGet, h.MethodOptions)
	}
}
