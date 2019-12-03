/*
 * Firefox Guardian API
 *
 * API to manage Guardian accounts, devices and servers
 *
 * API version: 0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package server

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/mozilla-services/guardian-vpn-windows/test/integrations/apimock/server/balrog"
	"github.com/mozilla-services/guardian-vpn-windows/test/integrations/apimock/server/fakewg"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Router struct {
	wg    *fakewg.Server
	chain *balrog.Chain
}
type Routes []Route

func NewRouter() (*mux.Router, error) {
	router := mux.NewRouter().StrictSlash(true)
	r := new(Router)
	var err error
	r.wg, err = fakewg.NewServer()
	r.chain, err = balrog.NewChain()
	if err != nil {
		return nil, err
	}
	GET := strings.ToUpper("get")
	POST := strings.ToUpper("post")
	DELETE := strings.ToUpper("delete")
	var routes = Routes{
		{
			"Index",
			GET,
			"/",
			Index,
		},

		{
			"ApiV1VpnAccountGet",
			GET,
			"/api/v1/vpn/account",
			r.ApiV1VpnAccountGet,
		},

		{
			"ApiV1VpnDevicePost",
			POST,
			"/api/v1/vpn/device",
			r.ApiV1VpnDevicePost,
		},

		{
			"ApiV1VpnDevicePubkeyDelete",
			DELETE,
			"/api/v1/vpn/device/{pubkey}",
			r.ApiV1VpnDevicePubkeyDelete,
		},

		{
			"ApiV1VpnLoginPost",
			POST,
			"/api/v1/vpn/login",
			r.ApiV1VpnLoginPost,
		},

		{
			"ApiV1VpnServersGet",
			GET,
			"/api/v1/vpn/servers",
			r.ApiV1VpnServersGet,
		},

		{
			"ApiV1VpnVersionsGet",
			GET,
			"/api/v1/vpn/versions",
			r.ApiV1VpnVersionsGet,
		},

		{
			"V1VpnLoginVerifyTokenGet",
			GET,
			"/v1/vpn/login/verify/{token}",
			r.V1VpnLoginVerifyTokenGet,
		},

		{
			"BalrogVersion",
			GET,
			"/json/1/FirefoxVPN/{version}/{useragent}/release/update.json",
			r.BalrogVersionGet,
		},

		{
			"BalrogRootSignature",
			GET,
			"/rootsig",
			r.BalrogRootSignatureGet,
		},

		{
			"BalrogSigtestChainGet",
			GET,
			"/chains/sigtest.chain",
			r.BalrogSigtestChainGet,
		},
		{
			"BalrogRegenerateCertPost",
			POST,
			"/regenerate",
			r.BalrogRegenerateCertPost,
		},
		{
			"UpdateSubscriptionStatus",
			POST,
			"/update/subscription",
			r.UpdateSubscriptionStatusPost,
		},
		{
			"DownloadMSI",
			GET,
			"/downloads/vpn/FirefoxPrivateNetworkVPN.msi",
			r.DownloadMSI,
		},
	}
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router, nil
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Mock API Server")
}