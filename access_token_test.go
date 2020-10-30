package offiaccount

import (
	"log"
	"testing"
)

const (
	APPID  = "wx8eccfdb653e0f210"
	SECRET = "b9e9bb0cd045e226a547eaf2f170da71"
)

func TestAccessToken(t *testing.T) {
	type args struct {
		appId  string
		secret string
	}
	tests := []struct {
		name    string
		args    args
		want    AccessTokenResult
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "AccessToken Test",
			args: args{appId: APPID, secret: SECRET},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := AccessToken(tt.args.appId, tt.args.secret)
			if err != nil {
				t.Errorf("AccessToken() error = %v", err)
				return
			}

			log.Printf("%#v", got)
		})
	}
}
