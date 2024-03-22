package dis_doi

import (
	"bufio"
	"bytes"
	"context"
	"fmt"
	"io"
	"strings"
	"sync"

	"github.com/miekg/dns/dnsutil"
)

var (
	whiteList = sync.Map{}
)

// 加载data下的文件，将文件内容解析成
func init() {
	if loadNum, err := initWhiteList("whitelist"); err != nil {
		panic(err)
	} else {
		fmt.Printf("load whitelist:%d \n", loadNum)
	}
}

func initWhiteList(sourceName string) (int64, error) {
	whiteList = sync.Map{}
	// 使用go-bindata库来将敏感词压缩成二进制文件，这样可以直接使用库，不需要部署文件。
	whitelist, err := Asset(sourceName)
	if err != nil {
		return 0, err
	}

	// 使用bufio读取sensitiveWords里面内容
	buf := bytes.NewBuffer(whitelist)
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

		whiteList.Store(strings.TrimSpace(string(line)), struct{}{})
		size++
	}

	return size, nil
}

// IsReserved 判断DOI是否是被保存的，只有某种情况下才可以注册
// 24.03.19 一期的策略是6位以下的doi都需要被保留，如果zone信息不为空，则先删除zone信息，
// 24.03.22 新增whitelist，如果在whitelist内的，直接返回false
// 然后判断剩下的部分是否是6位以下，如果小于等于6位，则返回true
func IsReserved(ctx context.Context, doi string, zone string) bool {
	// 首先doi删除zone的部分
	id := dnsutil.TrimDomainName(doi, zone)
	// 策略1，判断是否在whitelist内
	_, ok := whiteList.Load(id)
	if ok {
		return false
	}
	// 策略2，如果小于6位，返回true
	return len(id) <= 6
}
