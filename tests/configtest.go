package tests

import(
	"github.com/revel/revel/testing"
	"net/url"
	"sasuke/app/service/file"
	"os"
)

type ConfigTest struct{
	testing.TestSuite
}

var(
	testdb		string
	testhost	string
	testport	string
	testuser	string
	testname	string
	testpass	string
)

func (t *ConfigTest) Before(){}

func (t *ConfigTest) After(){}

func (t *ConfigTest) TestConfigPageWorks(){
	t.Get("/config")
	t.AssertOk()
	t.AssertContentType("text/html; charset=utf-8")
}

func (t *ConfigTest) TestConfigPostFormSuccessData(){

	testdb = "mysql"
	testhost = "192.168.192.168"
	testport = "3306"
	testuser = "testuser"
	testname = "testdb"
	testpass = "P@ssw0rd"

	configData := url.Values{}
	configData.Add("db", testdb)
	configData.Add("host", testhost)
	configData.Add("port", testport)
	configData.Add("dbuser", testuser)
	configData.Add("dbname", testdb)
	configData.Add("password", testpass)

	t.PostForm("/config/save", configData)

	t.AssertOk()
	t.AssertContentType("text/html; charset=utf-8")

	f := &file.Handler{}
	f.LoadEnv()
	t.AssertEqual(testdb, os.Getenv("db"))
	t.AssertEqual(testhost, os.Getenv("host"))
	t.AssertEqual(testport, os.Getenv("port"))
	t.AssertEqual(testuser, os.Getenv("dbuser"))
	t.AssertEqual(testdb, os.Getenv("dbname"))
	t.AssertEqual(testpass, os.Getenv("password"))
}
