package ews

import (
	"fmt"
	"reflect"
	"testing"
	"time"
)

func Test_getRFC3339Offset(t *testing.T) {

	riyadh, _ := time.LoadLocation("Asia/Riyadh")
	marquesas, _ := time.LoadLocation("Pacific/Marquesas")

	type args struct {
		time time.Time
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		want    string
	}{
		{
			name:    "test with timezone +03:00",
			args:    args{time: time.Now().In(riyadh)},
			wantErr: false,
			want:    "+03:00",
		},
		{
			name:    "test with timezone -09:30",
			args:    args{time: time.Now().In(marquesas)},
			wantErr: false,
			want:    "-09:30",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getRFC3339Offset(tt.args.time)
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

func TestResponseMarshal(t *testing.T) {
	xmlStr := `<?xml version="1.0" encoding="utf-8"?><s:Envelope xmlns:s="http://schemas.xmlsoap.org/soap/envelope/"><s:Header><h:ServerVersionInfo MajorVersion="15" MinorVersion="1" MajorBuildNumber="2375" MinorBuildNumber="28" Version="V2017_07_11" xmlns:h="http://schemas.microsoft.com/exchange/services/2006/types" xmlns="http://schemas.microsoft.com/exchange/services/2006/types" xmlns:xsd="http://www.w3.org/2001/XMLSchema" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"/></s:Header><s:Body xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xsd="http://www.w3.org/2001/XMLSchema"><m:CreateItemResponse xmlns:m="http://schemas.microsoft.com/exchange/services/2006/messages" xmlns:t="http://schemas.microsoft.com/exchange/services/2006/types"><m:ResponseMessages><m:CreateItemResponseMessage ResponseClass="Success"><m:ResponseCode>NoError</m:ResponseCode><m:Items><t:Message><t:ItemId Id="AAMkADExZTI1MjVmLWQ4MzEtNDBjYS1hMWEwLTYzMDIyNGQ5MWFjNwBGAAAAAAAsm1xVH2inSpA0JbGZEhm1BwBb0+EYLBFiQqC1D2X3n3odAAAA1zbaAABUGpQQZwqURZb0pqd+SpczAAJQbUuOAAA=" ChangeKey="CQAAABYAAABUGpQQZwqURZb0pqd+SpczAAJQdXwY"/></t:Message></m:Items></m:CreateItemResponseMessage></m:ResponseMessages></m:CreateItemResponse></s:Body></s:Envelope>`
	resp,err := checkCreateItemResponse([]byte(xmlStr))
	if err != nil {
		t.Error(err)
	}
	messages := *resp.Items.Message
	fmt.Println(messages[0].ItemId)
}
