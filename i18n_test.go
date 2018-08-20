package i18n

import (
	"testing"
)

func Test_Tr(t *testing.T) {

	if err := Init("testdata", "zh_CN"); err != nil {
		t.Error(err)
	}

	desc := Tr("zh_CN", "hi")
	if desc != "你好" {
		t.Errorf("expect '你好', got '%s'", desc)
	}

	desc = Tr("en_US", "hi")
	if desc != "hi" {
		t.Errorf("expect 'hi', got '%s'", desc)
	}

	desc = Tr("en_US", "test", "world")
	if desc != "hello,world" {
		t.Errorf("expect 'hello,world', got '%s'", desc)
	}

	desc = Tr("", "bye")
	if desc != "再见" {
		t.Errorf("expect '再见', got '%s'", desc)
	}

	desc = Tr("en_US", "byebye")
	if desc != "再见再见" {
		t.Errorf("expect '再见再见', got '%s'", desc)
	}

	desc = Tr("en_US", "byebyebye")
	if desc != "byebyebye" {
		t.Errorf("expect 'byebyebye', got '%s'", desc)
	}

	desc = Tr("zh_CN", "section.hi")
	if desc != "你好啊" {
		t.Errorf("expect '你好啊', got '%s'", desc)
	}

	desc = Tr("en_US", "section.hi")
	if desc != "你好啊" {
		t.Errorf("expect '你好啊', got '%s'", desc)
	}
}

func Benchmark_Tr(b *testing.B) {
	Init("testdata", "zh_CN")
	for i := 0; i < b.N; i++ {
		Tr("zh-CN", "hi")
	}
}
