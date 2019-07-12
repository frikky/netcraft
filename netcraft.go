package netcraft

import (
	"encoding/json"
	"fmt"
	"github.com/levigross/grequests"
	//"time"
)

// Stores login data
type Netcraftdata struct {
	Url    string
	Apikey string
	Ro     grequests.RequestOptions
}

type NewTakedown struct {
	Attack                    string `json:"attack"`
	Comment                   string `json:"comment"`
	Region                    string `json:"region,omitempty"`
	Brand                     string `json:"brand,omitempty"`
	Type                      string `json:"type,omitempty"`
	SuspectedFraudulentDomain string `json:"suspected_fraudulent_domain,omitempty"`
	Dropsite                  string `json:"dropsite,omitempty"`
	Evidence                  string `json:"evidence,omitempty"`
	Password                  string `json:"password,omitempty"`
	PhishkitFetchUrl          string `json:"phishkit_fetch_url,omitempty"`
	PhishkitPhishUrl          string `json:"phishkit_phish_url,omitempty"`
}

type Takedown struct {
	ID                 string `json:"id"`
	GroupID            string `json:"group_id"`
	AttackURL          string `json:"attack_url"`
	IP                 string `json:"ip"`
	Domain             string `json:"domain"`
	Hostname           string `json:"hostname"`
	AttackType         string `json:"attack_type"`
	CountryCode        string `json:"country_code"`
	Reporter           string `json:"reporter"`
	ReportSource       string `json:"report_source"`
	DateSubmitted      string `json:"date_submitted"`
	DateAuthed         string `json:"date_authed"`
	LastUpdated        string `json:"last_updated"`
	Region             string `json:"region"`
	TargetBrand        string `json:"target_brand"`
	Status             string `json:"status"`
	Authgiven          string `json:"authgiven"`
	FwdOwner           string `json:"fwd_owner"`
	RevOwner           string `json:"rev_owner"`
	ReverseDNS         string `json:"reverse_dns"`
	Host               string `json:"host"`
	Registrar          string `json:"registrar"`
	WhoisServer        string `json:"whois_server"`
	FirstResolved      string `json:"first_resolved"`
	FinalResolved      string `json:"final_resolved"`
	FirstOutage        string `json:"first_outage"`
	FinalOutage        string `json:"final_outage"`
	StatusChangeUptime string `json:"status_change_uptime"`
	Language           string `json:"language"`
	CustomerTag        string `json:"customer_tag"`
	HasPhishingKit     string `json:"has_phishing_kit"`
	DomainAttack       string `json:"domain_attack"`
	EvidenceURL        string `json:"evidence_url"`
	ScreenshotURL      string `json:"screenshot_url"`
	TargetedURL        string `json:"targeted_url"`
	StopMonitoringDate string `json:"stop_monitoring_date"`
	Certificate        struct {
		Name    string `json:"name"`
		Subject struct {
			CommonName string `json:"commonName"`
		} `json:"subject"`
		Hash   string `json:"hash"`
		Issuer struct {
			CountryName      string `json:"countryName"`
			OrganizationName string `json:"organizationName"`
			CommonName       string `json:"commonName"`
		} `json:"issuer"`
		Version        int    `json:"version"`
		SerialNumber   string `json:"serialNumber"`
		ValidFrom      string `json:"validFrom"`
		ValidTo        string `json:"validTo"`
		ValidFromTimeT int    `json:"validFrom_time_t"`
		ValidToTimeT   int    `json:"validTo_time_t"`
		Purposes       struct {
			Num1 []interface{} `json:"1"`
			Num2 []interface{} `json:"2"`
			Num3 []interface{} `json:"3"`
			Num4 []interface{} `json:"4"`
			Num5 []interface{} `json:"5"`
			Num6 []interface{} `json:"6"`
			Num7 []interface{} `json:"7"`
			Num8 []interface{} `json:"8"`
			Num9 []interface{} `json:"9"`
		} `json:"purposes"`
		Extensions struct {
			KeyUsage               string `json:"keyUsage"`
			ExtendedKeyUsage       string `json:"extendedKeyUsage"`
			BasicConstraints       string `json:"basicConstraints"`
			SubjectKeyIdentifier   string `json:"subjectKeyIdentifier"`
			AuthorityKeyIdentifier string `json:"authorityKeyIdentifier"`
			AuthorityInfoAccess    string `json:"authorityInfoAccess"`
			SubjectAltName         string `json:"subjectAltName"`
			CertificatePolicies    string `json:"certificatePolicies"`
			CtPrecertScts          string `json:"ct_precert_scts"`
		} `json:"extensions"`
	} `json:"certificate"`
	CertificateRevoked string `json:"certificate_revoked"`
	Managed            bool   `json:"managed"`
	DateEscalated      string `json:"date_escalated"`
}

// Defines API login principles that can be reused in requests
// Takes three parameters:
//  1. Username string
//  2. Password string
// Returns Netcraftdata struct
func CreateLogin(username string, password string) Netcraftdata {
	return Netcraftdata{
		Url: "https://takedown.netcraft.com",
		Ro: grequests.RequestOptions{
			Headers: map[string]string{
				"Content-Type": "application/json",
			},
			Auth: []string{username, password},
		},
	}
	//RequestTimeout: time.Duration(30) * time.Second,
}

func (netcraft *Netcraftdata) GetInfo(search map[string]string) ([]Takedown, error) {
	apiurl := fmt.Sprintf("%s/apis/get-info.php", netcraft.Url)

	search["date_from"] = "2019-01-01"
	search["date_to"] = "now"

	netcraft.Ro.Params = search

	ret, err := grequests.Get(apiurl, &netcraft.Ro)
	if err != nil {
		return []Takedown{}, err
	}

	parsedRet := new([]Takedown)
	err = json.Unmarshal(ret.Bytes(), parsedRet)
	if err != nil {
		return []Takedown{}, err
	}

	return *parsedRet, err
}

// No JSON return = returns bytes instead
func (netcraft *Netcraftdata) DoTakedown(takedown map[string]string) ([]byte, error) {
	apiurl := fmt.Sprintf("%s/authorise.php", netcraft.Url)

	netcraft.Ro.Params = takedown

	ret, err := grequests.Post(apiurl, &netcraft.Ro)
	if err != nil {
		return []byte{}, err
	}

	return ret.Bytes(), nil
}
