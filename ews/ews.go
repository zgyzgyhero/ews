// 26 august 2016
package ews

import (
	"bytes"
	"net/http"
)

// https://msdn.microsoft.com/en-us/library/office/dd877045(v=exchg.140).aspx
// https://arvinddangra.wordpress.com/2011/09/29/send-email-using-exchange-smtp-and-ews-exchange-web-service/
// https://msdn.microsoft.com/en-us/library/office/dn789003(v=exchg.150).aspx

const (
	soapStart = `<?xml version="1.0" encoding="utf-8" ?>
	<soap:Envelope xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" 
		xmlns:m="http://schemas.microsoft.com/exchange/services/2006/messages" 
		xmlns:t="http://schemas.microsoft.com/exchange/services/2006/types" 
		xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/">
  		<soap:Header>
    		<t:RequestServerVersion Version="Exchange2007_SP1" />
  		</soap:Header>
  		<soap:Body>
`
	soapEnd = `</soap:Body></soap:Envelope>`
)

func Issue(ewsAddr string, username string, password string, body []byte) (*http.Response, error) {

	bb := []byte(soapStart)
	bb = append(bb, body...)
	bb = append(bb, soapEnd...)

	req, err := http.NewRequest("POST", ewsAddr, bytes.NewReader(bb))
	if err != nil {
		return nil, err
	}
	req.SetBasicAuth(username, password)
	req.Header.Set("Content-Type", "text/xml")
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
	return client.Do(req)
}
