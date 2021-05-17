# 迎风
by jialanli 2021 持续更新中...

so surprise 自研的一款配置框架，欢迎支持和使用

    go get -t github.com/jialanli/windward/wind

1, 支持同时批量读取多个不同类型的配置文件
2, 支持json、yaml文件类型
3, 支持获取单条/多条配置、全文件注入

eg:

	import (
       "fmt"
       windward "github.com/jialanli/windward/wind"
    )
    
    func main() {
        name, name1 := "./s.json", "./c.yml"
        w := windward.GetWindward()
        w.InitConf([]string{name, name1})
        c4 := w.GetVal(name, "class.data")
        fmt.Println("class.data-->", c4)
        port := w.GetVal(name1, "spec.ports.port")
        fmt.Println("spec.ports.port-->", port)
    }
    
    ...多用途方式详见单元测试
持续改进和更新！
可详见大本营：https://lan6193.blog.csdn.net/article/details/116940283
  
  欢迎围观、感谢支持