package director

import (
	"fmt"
	"testing"

	"github.com/mdmdirector/mdmdirector/utils"
)

func TestRedisClientDefaultDB(t *testing.T) {

	host := utils.RedisHost()
	port := utils.RedisPort()
	password := utils.RedisPassword()

	connectionString := fmt.Sprintf("%v:%v", host, port)

	ans := RedisClient()

	if len(host) == 0 {
		t.Errorf("utils.RedisHost() error, no host name found")
	}

	if len(port) == 0 {
		t.Errorf("utils.RedisPort() error, no port found")
	}

	if len(password) == 0 {
		t.Errorf("utils.RedisPassword() error, no passwordfound")
	}

	if len(connectionString) == 0 {
		t.Errorf("connectionString not configured correctly, missing host or port")
	}

	if len(ans.Options().Addr) == 0 {
		t.Errorf("connectionString in new RedisClient not configured correctly, missing host or port")
	}

	if len(ans.Options().Password) == 0 {
		t.Errorf("password in new RedisClient not configured correctly")
	}

	if ans.Options().DB != 0 {
		t.Errorf("database in new RedisClient not configured correctly, should be 0")
	}
}
