package bloom

import (
	"island/crawler_distributed/gredis"
	"strconv"
	"testing"
)

func TestBloomFilter_Insert(t *testing.T) {
	gredis.Setup()
	i := 0
	for {
		err := NewBloomFilter().Insert("http://www.baidu.com" + strconv.Itoa(i))
		if err != nil{
			t.Errorf("err: %s", err.Error())
		}
		i++
		if i==1000{
			break
		}
	}

}

func TestBloomFilter_IsContains(t *testing.T) {
	gredis.Setup()
	i := 0
	for {
		b, err := NewBloomFilter().IsContains("http://www.baidu.com" + strconv.Itoa(i))
		if b == 0 || err != nil{
			t.Errorf("result: %v; err: %s; i: %d",
				b, "fail", i)
		}
		i++
		if i==1000{
			break
		}
	}
}
