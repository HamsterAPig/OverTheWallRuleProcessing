package proxy_info_map

// ProxyInfoNode 代理信息
type ProxyInfoNode struct {
	Type      string // 代理类型
	Server    string // 服务器名称或者IP
	Port      string // 服务器端口
	Username  string // 用户名
	Password  string // 密码
	Chipter   string // 加密方式
	OrderName string // 节点别名
}
