cd D:\goproject\src\github.com\testproject\database_server
$env:RABBITMQ_SERVER = "amqp://admin:admin@192.168.229.132:5672/test"
$env:LISTEN_ADDRESS = "127.0.0.1:12341"
$env:STORAGE_ROOT = "./tmp/1"
$env:GOPATH = "D:\goproject\src\github.com\testproject"
go run main.go

cd D:\goproject\src\github.com\testproject\interface_server
$env:RABBITMQ_SERVER = "amqp://admin:admin@192.168.229.132:5672/test"
$env:ES_SERVER="192.168.229.132:9200"
$env:LISTEN_ADDRESS = "127.0.0.1:12345"
$env:STORAGE_ROOT = "./tmp"
$env:GOPATH = "D:\goproject\src\github.com\testproject"
go run main.go



起四个数据服务：12341,12342,12343，12344
两个接口服务：12345,12346
$env:STORAGE_ROOT = "./tmp/1" ==>需要修改成不同的数据服务对应的文件夹目录

创建两个exchange(交换机)，使用消息订阅模式实现一对多的消息分发
1、apiServers
2、dataServers



curl -v localhost:12341/objects/test1 -XPUT -d"this is a test1 object"

curl -v localhost:12341/objects/test1 -XGET
