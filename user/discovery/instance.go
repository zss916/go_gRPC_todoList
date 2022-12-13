package discovery

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"google.golang.org/grpc/resolver"
)

// todo ETCD 结构体
type Server struct {
	Name    string `json:"name"`
	Addr    string `json:"addr"`    // 地址（存储服务器地址）
	Version string `json:"version"` // 版本 （方便服务的版本迭代）
	Weight  int64  `json:"weight"`  // 权重 （后续降级熔断）
}

// todo 定义服务名字前缀
func BuildPrefix(server Server) string {
	if server.Version == "" {
		return fmt.Sprintf("/%s/", server.Name)
	}
	return fmt.Sprintf("/%s/%s/", server.Name, server.Version)
}

// todo 定义注册地址
func BuildRegisterPath(server Server) string {
	return fmt.Sprintf("%s%s", BuildPrefix(server), server.Addr)
}

// todo 值反序列化一个注册server服务
func ParseValue(value []byte) (Server, error) {
	server := Server{}
	if err := json.Unmarshal(value, &server); err != nil {
		return server, err
	}

	return server, nil
}

// todo 分割路径,后续作为server地址的更新
func SplitPath(path string) (Server, error) {
	server := Server{}
	strs := strings.Split(path, "/")
	if len(strs) == 0 {
		return server, errors.New("invalid path")
	}

	server.Addr = strs[len(strs)-1]

	return server, nil
}

// todo Exist helper function,判断
func Exist(l []resolver.Address, addr resolver.Address) bool {
	for i := range l {
		if l[i].Addr == addr.Addr {
			return true
		}
	}

	return false
}

// Remove helper function
func Remove(s []resolver.Address, addr resolver.Address) ([]resolver.Address, bool) {
	for i := range s {
		if s[i].Addr == addr.Addr {
			s[i] = s[len(s)-1]
			return s[:len(s)-1], true
		}
	}
	return nil, false
}

func BuildResolverUrl(app string) string {
	return schema + ":///" + app
}
