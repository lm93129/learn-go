## Go modules
使用go modules的go程序，会在目录下有一个go.mod和go.sum的文件。
mod这个文件是记录使用了哪些依赖包，sum这个文件是记录依赖包的校验码。
程序下载好之后，可以直接使用go build或者go run来启动程序。此时程序会自动安装mod中的信息下载依赖包
这里推荐设置下下载依赖的代理为goproxy.cn。具体设置方法可以谷歌下。
# 一、go语言学习入门
## 1. 基本语法
### 1.0 HelloWorld
### 1.1 基本数据类型 
### 1.2 基本条件判断
### 1.3 数组和切片
### 1.4 Map和字典类型的使用
### 1.5 Goroutine实现并发
## 2. 算法相关
### 2.1 LeetCode算法题
### 2.2 常见排序算法
## 3. 应用相关
### 3.1 爬虫应用
#### 3.1.1 爬取百度企业信用信息
#### 3.1.2 爬取空气质量信息
#### 3.1.3 爬取免费快代理