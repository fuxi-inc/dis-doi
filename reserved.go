package dis_doi

import (
	"context"

	"github.com/miekg/dns/dnsutil"
)

// IsReserved 判断DOI是否是被保存的，只有某种情况下才可以注册
// 24.03.19 一期的策略是6位以下的doi都需要被保留，如果zone信息不为空，则先删除zone信息，
// 然后判断剩下的部分是否是6位以下，如果小于等于6位，则返回true
func IsReserved(ctx context.Context, doi string, zone string) bool {
	// 首先doi删除zone的部分
	// 策略1，如果小于6位，返回true
	return len(dnsutil.TrimDomainName(doi, zone)) <= 6
}
