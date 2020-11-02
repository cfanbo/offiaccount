package offiaccount

import (
	"sync"
	"testing"
)

func BenchmarkNewTemplateMessage(b *testing.B) {
	for i := 1; i < b.N; i++ {
		t := NewTemplateMessage("templateId")
		t.SetUser("openId")
		t.AddDataItem("first", TemplateMessageDataItem{"first", ""}).
			AddDataItem("keyword1", TemplateMessageDataItem{"kw", ""}).
			AddDataItem("keyword2", TemplateMessageDataItem{"kw", ""}).
			AddDataItem("keyword1", TemplateMessageDataItem{"kw", ""}).
			AddDataItem("keyword2", TemplateMessageDataItem{"kw", ""}).
			AddDataItem("remark", TemplateMessageDataItem{"remark", ""})

		t.Json()
	}
}

// 客户端口自定义的pool
// 性能最好的方法
func BenchmarkNewTemplateMessagePool(b *testing.B) {
	pool := sync.Pool{
		New: func() interface{} {
			return NewTemplateMessage("")
		},
	}

	for i := 1; i < b.N; i++ {
		t := pool.Get().(*TemplateMessage)
		t.SetTemplateId("templateId").SetUser("openId")
		t.AddDataItem("first", TemplateMessageDataItem{"first", ""}).
			AddDataItem("keyword1", TemplateMessageDataItem{"kw", ""}).
			AddDataItem("keyword2", TemplateMessageDataItem{"kw", ""}).
			AddDataItem("keyword3", TemplateMessageDataItem{"kw", ""}).
			AddDataItem("keyword4", TemplateMessageDataItem{"kw", ""}).
			AddDataItem("keyword5", TemplateMessageDataItem{"kw", ""}).
			AddDataItem("remark", TemplateMessageDataItem{"remark", ""})
		t.Json()

		t.TemplateId = ""
		if t.Miniprogram != nil {
			t.Miniprogram = make(map[string]string)
		}
		t.Data = make(map[string]TemplateMessageDataItem)
		pool.Put(t)
	}
}

func BenchmarkTemmplateMessagePool2(b *testing.B) {
	p := NewTemplateMessagePool()

	for i := 1; i < b.N; i++ {
		t := p.Get()

		t.SetTemplateId("templateId").SetUser("openId")
		t.AddDataItem("first", TemplateMessageDataItem{"first", ""}).
			AddDataItem("keyword1", TemplateMessageDataItem{"kw", ""}).
			AddDataItem("keyword2", TemplateMessageDataItem{"kw", ""}).
			AddDataItem("keyword3", TemplateMessageDataItem{"kw", ""}).
			AddDataItem("keyword4", TemplateMessageDataItem{"kw", ""}).
			AddDataItem("keyword5", TemplateMessageDataItem{"kw", ""}).
			AddDataItem("remark", TemplateMessageDataItem{"remark", ""})
		t.Json()

		p.Put(t)
	}

}
