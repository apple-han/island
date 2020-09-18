package bloom

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"github.com/gomodule/redigo/redis"
	"island/crawler_distributed/gredis"
	"strconv"
)
var (
	BitSize uint = 1<<31
	Seeds = []uint{1,2,3,5}
	BlockNum uint = 1
	Key = "bloomfilter"
	Ret = 1
)
type SimpleHash struct {
	Cap uint
	Seed uint
}

type BloomFilter struct {
	BitSize uint
	Seeds   []uint
	Key     string
	BlockNum uint
	HashFunc []*SimpleHash
}

func NewBloomFilter() *BloomFilter{
	var HashMiddleFunc []*SimpleHash
	for _, v := range Seeds{
		HashMiddleFunc = append(HashMiddleFunc, &SimpleHash{
			Cap:  BitSize,
			Seed: v,
		})
	}
	return &BloomFilter{
		BitSize:  BitSize,
		Seeds:    Seeds,
		Key:      Key,
		BlockNum: BlockNum,
		HashFunc: HashMiddleFunc,
	}
}

// IsContains int 1 represent has already existed
func (b *BloomFilter) IsContains(s string) (int, error){
	if len(s) == 0{
		return 0, errors.New("input is must")
	}
	input := b.MD5(s)
	n, err := strconv.ParseUint(input[0:2], 16, 64)
	if err != nil {
		return 0, errors.New("string to uint64 fail")
	}
	name := b.Key + strconv.Itoa(int(uint(n) % b.BlockNum))
	for _, f := range b.HashFunc{
		loc := f.Hash(input)
		r, err := redis.Int(gredis.GetBit(name, loc))
		if err != nil{
			return 0, err
		}
		Ret = Ret & r
	}

	return Ret, nil
}

func (b *BloomFilter) Insert(s string) error{
	if len(s) == 0{
		return errors.New("input is must")
	}
	input := b.MD5(s)
	n, err := strconv.ParseUint(input[0:2], 16, 64)
	if err != nil {
		return errors.New("string to uint64 fail")
	}
	name := b.Key + strconv.Itoa(int(uint(n) % b.BlockNum))
	for _, f := range b.HashFunc{
		loc := f.Hash(input)
		gredis.SetBit(name, loc, 1)
	}
	return  nil
}

func (s *SimpleHash) Hash(value string) uint{
	var ret uint = 0
	for _, v := range value{
		ret += s.Seed * ret + uint(v)
	}
	var bz = BitSize-1
	r := bz & ret
 	return r
}

func (b *BloomFilter) MD5(value string) string {
	m := md5.New()
	m.Write([]byte(value))
	return hex.EncodeToString(m.Sum(nil))
}
