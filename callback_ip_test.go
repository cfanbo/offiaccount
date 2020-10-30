package offiaccount

import (
	"reflect"
	"testing"
)

func TestGetCallbackIp(t *testing.T) {
	ac, err := AccessToken(APPID, SECRET)
	if err != nil {
		t.Fatal(err)
	}
	type args struct {
		accessToken string
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "CallbackIp Test",
			args: args{
				accessToken: ac.AccessToken,
			},
			want:    []string{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetCallbackIp(tt.args.accessToken)
			if err != nil {
				t.Errorf("GetCallbackIp() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Log(got)
			t.Log(reflect.TypeOf(got).Kind())

		})
	}
}
