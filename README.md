# dis-doi
dis系统中关于专为DOI服务的lib，用于dis-r、dis-chaincode等

sensitive_words里存的是所有的敏感词，需要使用go-bindata生成静态文件(.go)，编译进bin文件
执行命令
go-bindata -pkg dis_doi -o static.go ./sensitive_words ./sensitive_words_test ./whitelist