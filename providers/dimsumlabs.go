package providers

import (
	"log"
	"net/http"
	"net/url"

	"github.com/bitly/oauth2_proxy/api"
)

type DimsumlabsProvider struct {
	*ProviderData
}

func NewDimsumlabsProvider(p *ProviderData) *DimsumlabsProvider {
	const dslHost string = "door.dimsumlabs.com"

	p.ProviderName = "Dimsumlabs"
	if p.LoginURL.String() == "" {
		p.LoginURL = &url.URL{
			Scheme: "https",
			Host:   dslHost,
			Path:   "/oauth/authorize"}
	}
	if p.RedeemURL.String() == "" {
		p.RedeemURL = &url.URL{
			Scheme: "https",
			Host:   dslHost,
			Path:   "/oauth/token/"}
	}
	if p.ProfileURL.String() == "" {
		p.ProfileURL = &url.URL{
			Scheme: "https",
			Host:   dslHost,
			Path:   "/api/v1/profile"}
	}
	if p.ValidateURL.String() == "" {
		p.ValidateURL = &url.URL{
			Scheme: "https",
			Host:   dslHost,
			Path:   "/api/v1/tokeninfo"}
	}
	if p.Scope == "" {
		p.Scope = "read"
	}
	return &DimsumlabsProvider{ProviderData: p}
}

func (p *DimsumlabsProvider) GetEmailAddress(s *SessionState) (string, error) {
	req, err := http.NewRequest("GET",
		p.ProfileURL.String()+"?access_token="+s.AccessToken, nil)
	if err != nil {
		log.Printf("failed building request %s", err)
		return "", err
	}
	json, err := api.Request(req)
	if err != nil {
		log.Printf("failed making request %s", err)
		return "", err
	}
	return json.Get("email").String()
}
