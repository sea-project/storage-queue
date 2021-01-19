package queue

import "testing"

func TestQueuePool_Push(t *testing.T) {
	Single().Push("123456")
	t.Log(Single().Len())
	t.Log(Single().Pop())
	t.Log(Single().Len())
}

func BenchmarkPool_Push(b *testing.B) {
	for i:=0; i<b.N; i++ {
		Single().Push("123456")
	}
}