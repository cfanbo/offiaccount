package offiaccount

import (
	"testing"
)

func TestNewRequestParse(t *testing.T) {
	appId := "wx9749dc353f606262"
	key := "as8dfafyasl2kjq8giajdklfaslkdjfklasdfklja37"
	xmlRaw := `<xml>
    <AppId><![CDATA[wx9749dc353f606262]]></AppId>
    <Encrypt><![CDATA[G9rg9BXXHWEShtSpac+T6JsSmnUUHc0l6r+wWf+Bm6ulrvXEg8wHjDGneKr9lhiuyiIcYjXnMw6G6p2pVeB11hrHx64Xz4BPi763S8favRoisECZ5F2Ryndd2Bg4DO2m9kiNHfwo3QzW7F+e5U/cnuz4t8dmsAz03+QFFrexukkjbS82gfM9OLjqZMcuiwSrvRz29/F9GzVSG
bgCE9G0alFX3h4+PkCrBWo2LU2Vo9INB6mcK3Gtuo2Lq3MUB/quCgovfLI+meUWQtOfWZ5W5FBbuDMOhdQ8dhwgOF84Pg1G+SOrnblRrHCI6S+bIhjLo26sC/1+1bmJGKxrmiMUbwfCGqdWXkouh41D2VlntbWF2TueHNi+W3pZPKUC2DbtG5GETBURCrM6SAXJc7G1+1UyS+W78VgtHcv8zUz7Q6sETGz+
TIorE5KXtGaQJlkkuIRgclihHBTgt8CyzsGopA==]]></Encrypt>
</xml>`

	r := NewRequest(appId, key)
	ticket, err := r.NewTicket([]byte(xmlRaw))
	if err != nil {
		t.Errorf("%s", err)
	}
	if ticket.AppId != appId {
		t.Errorf("解析失败")
	}
}
