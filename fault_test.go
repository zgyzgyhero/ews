package ews

import (
	"errors"
	"io"
	"net/http"
	"reflect"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	soapMessage = `
		<?xml version="1.0" encoding="utf-8"?><s:Envelope
    xmlns:s="http://schemas.xmlsoap.org/soap/envelope/">
    <s:Header>
        <h:ServerVersionInfo MajorVersion="15" MinorVersion="20" MajorBuildNumber="2495" MinorBuildNumber="20" Version="V2018_01_08"
            xmlns:h="http://schemas.microsoft.com/exchange/services/2006/types"
            xmlns:xsd="http://www.w3.org/2001/XMLSchema"
            xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"/>
        </s:Header>
        <s:Body>
            <m:CreateItemResponse
                xmlns:m="http://schemas.microsoft.com/exchange/services/2006/messages"
                xmlns:xsd="http://www.w3.org/2001/XMLSchema"
                xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
                xmlns:t="http://schemas.microsoft.com/exchange/services/2006/types">
                <m:ResponseMessages>
                    <m:CreateItemResponseMessage ResponseClass="Success">
                        <m:ResponseCode>NoError</m:ResponseCode>
                        <m:Items/>
                    </m:CreateItemResponseMessage>
                </m:ResponseMessages>
            </m:CreateItemResponse>
        </s:Body>
    </s:Envelope>
	`
	soapMessageWithFault = `
		<?xml version="1.0" encoding="utf-8"?><s:Envelope xmlns:s="http://schemas.xmlsoap.org/soap/envelope/">
    <s:Header>
        <Action s:mustUnderstand="1"
            xmlns="http://schemas.microsoft.com/ws/2005/05/addressing/none">*
        </Action>
    </s:Header>
    <s:Body>
        <s:Fault>
            <faultcode
                xmlns:a="http://schemas.microsoft.com/exchange/services/2006/types">a:ErrorSchemaValidation</faultcode>
            <faultstring xml:lang="en-US">The request failed schema validation: End element 'Envelope' from namespace 'http://schemas.xmlsoap.org/soap/envelope/' expected. Found text '\n'. Line 4, position 60.</faultstring>
            <detail>
                <e:ResponseCode
                    xmlns:e="http://schemas.microsoft.com/exchange/services/2006/errors">ErrorSchemaValidation</e:ResponseCode>
                <e:Message
                    xmlns:e="http://schemas.microsoft.com/exchange/services/2006/errors">The request failed schema validation.</e:Message>
                <t:MessageXml
                    xmlns:t="http://schemas.microsoft.com/exchange/services/2006/types">
                    <t:LineNumber>0</t:LineNumber>
                    <t:LinePosition>0</t:LinePosition>
                    <t:Violation>End element 'Envelope' from namespace 'http://schemas.xmlsoap.org/soap/envelope/' expected. Found text '\n'. Line 4, position 60.</t:Violation>
                </t:MessageXml>
            </detail>
        </s:Fault>
    </s:Body>
</s:Envelope>
	`
)

var (
	soapErr = &SoapError{Fault: &Fault{
		Faultcode:   "a:ErrorSchemaValidation",
		Faultstring: `The request failed schema validation: End element 'Envelope' from namespace 'http://schemas.xmlsoap.org/soap/envelope/' expected. Found text '\n'. Line 4, position 60.`,
		Detail: detail{
			ResponseCode: "ErrorSchemaValidation",
			Message:      "The request failed schema validation.",
			MessageXml: faultMessageXml{
				LineNumber:   "0",
				LinePosition: "0",
				Violation:    `End element 'Envelope' from namespace 'http://schemas.xmlsoap.org/soap/envelope/' expected. Found text '\n'. Line 4, position 60.`,
			},
		},
	}}
)

func Test_parseSoapFault(t *testing.T) {
	type args struct {
		soapMessage string
	}
	tests := []struct {
		name    string
		args    args
		want    *Fault
		wantErr bool
	}{
		{
			name: "test with soap fault",
			args: args{soapMessage: soapMessageWithFault},
			want: soapErr.Fault,
		},
		{
			name: "test with no soap fault",
			args: args{soapMessage: soapMessage},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseSoapFault(tt.args.soapMessage)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseSoapFault() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseSoapFault() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewSoapError(t *testing.T) {
	type args struct {
		resp *http.Response
	}
	tests := []struct {
		name     string
		args     args
		expected error
	}{
		{
			name:     "test when reading from response return error",
			args:     args{resp: &http.Response{Body: fakeErrorReadCloser{}}},
			expected: errors.New("error while reading"),
		},
		{
			name: "test when reading from response without soap fault",
			args: args{resp: &http.Response{
				Body: &FakeReadCloser{contents: strings.NewReader(soapMessageWithFault)}},
			},
			expected: soapErr,
		},
		{
			name: "test no soap fault but http error",
			args: args{resp: &http.Response{
				Status:     "gateway error",
				StatusCode: http.StatusBadGateway,
				Body:       &FakeReadCloser{contents: strings.NewReader(soapMessage)}},
			},
			expected: &HTTPError{Status: "gateway error", StatusCode: http.StatusBadGateway},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := NewError(tt.args.resp)
			assert.Equal(t, tt.expected, err)
		})
	}
}

// --------------
type fakeErrorReadCloser struct{}

func (f fakeErrorReadCloser) Read(p []byte) (n int, err error) {
	return 0, errors.New("error while reading")
}
func (f fakeErrorReadCloser) Close() error {
	return nil
}

// --------------

type FakeReadCloser struct {
	contents io.Reader
}

func (f *FakeReadCloser) Read(p []byte) (n int, err error) {
	return f.contents.Read(p)
}
func (f *FakeReadCloser) Close() error {
	return nil
}
