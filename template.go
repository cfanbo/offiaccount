package offiaccount

import (
	"encoding/json"
	"errors"
	"sync"
)

type TemplateMessageDataItem struct {
	Value string `json:"value"`
	Color string `json:"color"`
}

type TemplateMessage struct {
	mu          sync.Mutex
	ToUser      string                             `json:"touser"`
	TemplateId  string                             `json:"template_id"`
	Url         string                             `json:"url,omitempty"`
	Miniprogram map[string]string                  `json:"miniprogram,omitempty"`
	Data        map[string]TemplateMessageDataItem `json:"data"`
}

func NewTemplateMessage(templateID string) *TemplateMessage {
	return &TemplateMessage{
		TemplateId: templateID,
		Data:       make(map[string]TemplateMessageDataItem),
	}
}

// SetTemplateId 设置模板ID
func (t *TemplateMessage) SetTemplateId(tid string) *TemplateMessage {
	t.TemplateId = tid
	return t
}

// SetUser 设置用户OPENID
func (t *TemplateMessage) SetUser(openID string) *TemplateMessage {
	t.ToUser = openID
	return t
}

// SetUrl 设置跳转URL
func (t *TemplateMessage) SetUrl(url string) *TemplateMessage {
	t.Url = url
	return t
}

// SetMiniprogram 设置小程序
func (t *TemplateMessage) SetMiniprogram(appid, pagepath string) *TemplateMessage {
	t.mu.Lock()
	defer t.mu.Unlock()

	t.Miniprogram["appid"] = appid
	t.Miniprogram["pagepath"] = pagepath

	return t
}

// NewItem 创建消息模板数据项
func (t *TemplateMessage) NewDataItem(value, color string) TemplateMessageDataItem {
	if color == "" {
		color = "#000000"
	}
	return TemplateMessageDataItem{
		Value: value,
		Color: color,
	}
}

// addDataItem 添加消息内容
func (t *TemplateMessage) AddDataItem(key string, value TemplateMessageDataItem) *TemplateMessage {
	t.mu.Lock()
	defer t.mu.Unlock()

	t.Data[key] = value

	return t
}

// Json 返回模板消息JSON数据格式
func (t *TemplateMessage) Json() ([]byte, error) {
	if len(t.Data) < 0 || len(t.Data) > 5 {
		return nil, errors.New("模板消息个数不合法")
	}
	return json.Marshal(t)
}

type TemplateMessagePool struct {
	pool sync.Pool
}

func NewTemplateMessagePool() *TemplateMessagePool {
	pool := sync.Pool{
		New: func() interface{} {
			return NewTemplateMessage("")
		},
	}

	return &TemplateMessagePool{
		pool: pool,
	}
}

// Get 获取
func (p *TemplateMessagePool) Get() *TemplateMessage {
	return p.pool.Get().(*TemplateMessage)
}

// Put 放入池
func (p *TemplateMessagePool) Put(t *TemplateMessage) {
	t.SetTemplateId("").SetUser("").SetUrl("")
	if t.Miniprogram != nil {
		t.SetMiniprogram("", "")
	}
	t.Data = make(map[string]TemplateMessageDataItem)
	t.Miniprogram = make(map[string]string)

	p.pool.Put(t)
}
