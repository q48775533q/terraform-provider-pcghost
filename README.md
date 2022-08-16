本仓库存的是个terraform的“自建供应实例的例子”

```bash
mkdir ${GOPATH}/src
git clone https://github.com/q48775533q/terraform-provider-pcghost.git
# 准备一个mysql数据库用作api服务的模拟。并建立相应的数据库，用户名和密码
yum install -y mariadb-server

systemctl start mariadb 
mysql
create database pets;
grant all on pets.* to pets@'%' identified by 'pets';
delete from mysql.user where User='';
exit

systemctl restart mariadb

export rds_user=pets
export rds_password=pets
export rds_host=127.0.0.1
export rds_port=3306
export rds_database=pets

cd ${GOPATH}/src/terraform-provider-pcghost/server
go run .

cd ${GOPATH}/src/terraform-provider-pcghost/
go mod init
go mod tidy

# makefile需要根据实际情况修改下，根据具体的输出复制粘贴即可
# 根据实际情况修改 example 里面的main.tf 即可做最简单的测试
make install 

yum install -y yum-utils
yum-config-manager --add-repo https://rpm.releases.hashicorp.com/RHEL/hashicorp.repo
yum -y install terraform

cd ${GOPATH}/src/terraform-provider-pcghost/example 
terraform init
terraform plan
terraform apply
```







├── example
│   └── main.tf             < -- terraform 范例
├── go.mod
├── go.sum
├── main.go                    < -- terraform 的 provider入口文件
├── Makefile                   < -- go程序的 build 文件
├── pcghost
│   ├── datasource_ps_pet_ids.go
│   ├── datasource_ps_pet_ids_test.go
│   ├── provider.go
│   ├── provider_test.go
│   ├── resource_ps_pet.go
│   └── resource_ps_pet_test.go
└── server                   < --该文件夹用作模拟api的服务端，默认端口为8000
    ├── action
    │   └── pets
    │       ├── create.go
    │       ├── delete.go
    │       ├── get.go
    │       ├── list.go
    │       └── update.go
    ├── go.mod
    ├── go.sum
    ├── main.go
    ├── model
    │   └── pet
    │       ├── model.go
    │       └── orm.go
    └── README.md



本仓库参考 terraform 实战编写
