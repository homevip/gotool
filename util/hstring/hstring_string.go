package hstring

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"os"
	"sync"
	"sync/atomic"
	"time"

	"github.com/bwmarrin/snowflake"
	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/guuid"
)

// 获取 UUID
func UUID() (uuid string) {

	u := guuid.New()
	uuid = fmt.Sprint(u)

	return
}

// 更好的随机值
func EnUniqid(unique ...string) (s string) {
	if len(unique) == 0 {
		uuid := UUID()
		s, _ = gmd5.Encrypt(uuid)
	} else {
		s, _ = gmd5.Encrypt(unique[0] + gtime.Date())
	}

	return
}

// Sha256 加密
func Sha256(src string) (res string) {
	m := sha256.New()
	m.Write([]byte(src))
	res = hex.EncodeToString(m.Sum(nil))
	return
}

// MD5 加密 (32位)
func MD5(src string) (res string) {
	b := md5.Sum([]byte(src))
	s := fmt.Sprintf("%x", b)
	res = string(s)

	return
}

// MD5 加密 (16位)
func MD5To16(src string) string {
	return MD5(src)[8:24]
}

// 生成24位订单号
// 前面17位代表时间精确到毫秒，中间3位代表进程id，最后4位代表序号
// 生成订单号
// 生成位数,19位日期模板:20060102(年月日), 24位日期模板:20060102150405(年月日时分秒)
var increment int64 // 增量

func BuilderOrderSn() string {

	var (
		t = time.Now()
	)

	s := time.Now().Format("20060102150405")
	m := t.UnixNano()/1e6 - t.UnixNano()/1e9*1e3
	ms := sup(m, 3)
	p := os.Getpid() % 1000
	ps := sup(int64(p), 3)
	i := atomic.AddInt64(&increment, 1)
	r := i % 10000
	rs := sup(r, 4)
	n := fmt.Sprintf("%s%s%s%s", s, ms, ps, rs)
	return n
}

// 全局雪花节点 + 互斥锁(保证线程安全）
var (
	snowNode *snowflake.Node
	once     sync.Once // 保证节点只初始化一次
)

func SnowflakeID() (unique int64, err error) {

	// 使用sync.Once保证节点只初始化一次，线程安全
	once.Do(func() {
		var initErr error
		snowNode, initErr = snowflake.NewNode(1)
		if initErr != nil {
			err = initErr // 将初始化错误传递出去
		}
	})

	// 检查初始化是否失败
	if err != nil {
		return 0, fmt.Errorf("初始化雪花节点失败: %w", err)
	}

	if snowNode == nil {
		return 0, fmt.Errorf("雪花节点初始化后为空")
	}

	// 生成唯一ID
	unique = snowNode.Generate().Int64()

	return
}

// 对长度不足n的数字前面补0
func sup(i int64, n int) string {
	m := fmt.Sprintf("%d", i)
	for len(m) < n {
		m = fmt.Sprintf("0%s", m)
	}
	return m
}
