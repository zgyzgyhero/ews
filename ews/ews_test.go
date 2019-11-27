package ews

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseSoapFault(t *testing.T) {
	soapMessage := `
		<?xml version="1.0" encoding="utf-8"?><s:Envelope
    xmlns:s="http://schemas.xmlsoap.org/soap/envelope/">
    <s:Header>
        <Action s:mustUnderstand="1"
            xmlns="http://schemas.microsoft.com/ws/2005/05/addressing/none">*
        </Action>
    </s:Header>
    <s:Body>
        <s:Fault>
            <faultcode
                xmlns:a="http://schemas.microsoft.com/exchange/services/2006/types">a:ErrorSchemaValidation
            </faultcode>
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

	soapFault, err := parseSoapFault(soapMessage)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, "ErrorSchemaValidation", soapFault.Detail.ResponseCode)
	assert.Contains(t, "The request failed schema validation.", soapFault.Detail.Message)
}

func TestParseSoapFaultHasNoError(t *testing.T) {
	soapMessage := `
		<?xml version="1.0" encoding="utf-8"?>
<s:Envelope
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

	soapFault, err := parseSoapFault(soapMessage)
	if err != nil {
		t.Fatal(err)
	}
	assert.Nil(t, soapFault)
}
