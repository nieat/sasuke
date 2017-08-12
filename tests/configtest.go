package tests

import(
	"github.com/revel/revel/testing"
	"net/url"
)

type ConfigTest struct{
	testing.TestSuite
}

func (t *ConfigTest) Before(){}

func (t *ConfigTest) After(){}

func (t *ConfigTest) TestConfigPageWorks(){
	t.Get("/config")
	t.AssertOk()
	t.AssertContentType("text/html; charset=utf-8")
}

func (t *ConfigTest) TestConfigPostFormSuccessData(){
	configData := url.Values{}
	configData.Add("db", "mysql")
	configData.Add("host", "192.168.192.168")
	configData.Add("port", "3306")
	configData.Add("dbuser", "testuser")
	configData.Add("dbname", "testdb")
	configData.Add("password", "P@ssw0rd")

	t.PostForm("/config/save", configData)

	t.AssertOk()
	t.AssertContentType("text/html; charset=utf-8")
}
