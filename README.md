# EosTest

## 设计

本测试使用EOSIO 提供的测试合约 currency 实现

构架： 后端gin 前端vuejs+iview

由于官方api问题部分 功能使用命令行工具代替

## 目标

在区块链中提交两个只能合约AWcoin 和GameCoin

每当 GameGoin用户之间进行转账时都会向转账方收取 1 AWcoin 作为手续费

在页面上可以进行：
1. 查询用户合约内余额
2. 用户之间进行交易
3. 系统向用户发行合约Token

## 如何使用
### 流程
1. 启动 nodeos 节点
2. 创建钱包 cleos create wallet
3. 创建测试私钥： cleos create key   
    eg：5Kh9eQZQLoyA4aEknx5MMhLrBTHaZfqd3nEgpJhojkTYkDn2Lki
4. 导入私钥：cleos wallet import 5Kh9eQZQLoyA4aEknx5MMhLrBTHaZfqd3nEgpJhojkTYkDn2Lki
5. 创建awcoin账户：cleos create account eosio awcoin EOS4vfrp1WfL4UEmpGJn4JxfVjAXKrH5sPK5VG8efVU9eCcUYVEbF EOS4vfrp1WfL4UEmpGJn4JxfVjAXKrH5sPK5VG8efVU9eCcUYVEbF
6. 创建gamecoin账户：cleos create account eosio gamecoin EOS4vfrp1WfL4UEmpGJn4JxfVjAXKrH5sPK5VG8efVU9eCcUYVEbF EOS4vfrp1WfL4UEmpGJn4JxfVjAXKrH5sPK5VG8efVU9eCcUYVEbF
7. cleos set contract awcoin currency 向awcoin 账户导入合约
8. cleos set contract gamecoin currency 向gamecoin 账户导入合约
9. cleos push action awcoin create '{"issuer":"awcoin","maximum_supply":"1000000.0000 CUR","can_freeze":"0","can_recall":"0","can_whitelist":"0"}' --permission awcoin@active
10. cleos push action gamecoin create '{"issuer":"gamecoin","maximum_supply":"1000000.0000 CUR","can_freeze":"0","can_recall":"0","can_whitelist":"0"}' --permission gamecoin@active
11. 创建两个 player账户 或其他自己喜欢的账户
cleos create account eosio player1 EOS4vfrp1WfL4UEmpGJn4JxfVjAXKrH5sPK5VG8efVU9eCcUYVEbF EOS4vfrp1WfL4UEmpGJn4JxfVjAXKrH5sPK5VG8efVU9eCcUYVEbF  
cleos create account eosio player2 EOS4vfrp1WfL4UEmpGJn4JxfVjAXKrH5sPK5VG8efVU9eCcUYVEbF EOS4vfrp1WfL4UEmpGJn4JxfVjAXKrH5sPK5VG8efVU9eCcUYVEbF
12. clone 项目 go build 后运行编译好的二进制文件

运行二进制文件 默认8080端口即可看到页面

## 注意
该项目必须和 nodeos 服务器部署在同一台机器上
nodeos cleos 必须在系统环境变量下

