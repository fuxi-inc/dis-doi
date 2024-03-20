package dis_doi

import (
	"bufio"
	"bytes"
	"context"
	"fmt"
	"io"
	"strings"
	"sync"

	"github.com/miekg/dns"
)

var (
	validateMap = sync.Map{}
)

// 加载data下的文件，将文件内容解析成
func init() {
	if loadNum, err := initSensitiveWords("sensitive_words"); err != nil {
		panic(err)
	} else {
		fmt.Printf("load sensitive_words:%d \n", loadNum)
	}
}

func initSensitiveWords(sourceName string) (int64, error) {
	validateMap = sync.Map{}
	// 使用go-bindata库来将敏感词压缩成二进制文件，这样可以直接使用库，不需要部署文件。
	sensitiveWords, err := Asset(sourceName)
	if err != nil {
		return 0, err
	}

	// 使用bufio读取sensitiveWords里面内容
	buf := bytes.NewBuffer(sensitiveWords)
	rd := bufio.NewReader(buf)
	size := int64(0)

	for {
		line, _, err := rd.ReadLine()
		if err == io.EOF {
			break
		}

		if err != nil {
			return 0, err
		}

		validateMap.Store(strings.TrimSpace(string(line)), struct{}{})
		size++
	}

	return size, nil
}

// IsValidated 验证doi是否包含敏感词，如果包含敏感词，不让注册
func IsValidated(ctx context.Context, doi string, zone string) bool {
	// 首先doi删除zone的部分
	fqdnName := dns.Fqdn(strings.ToLower(doi))
	doiPre := strings.TrimSuffix(fqdnName, "."+dns.Fqdn(strings.ToLower(zone)))

	// 查看doiPre是否命中敏感词，命中返回false，未命中返回true
	_, ok := validateMap.Load(doiPre)
	return !ok
}
