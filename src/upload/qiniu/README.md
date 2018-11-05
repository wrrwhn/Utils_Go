
## 描述
- 将指定 [文件|文件夹|当前文件夹] 的内容上传至 `七牛云`
- 上传完毕后将链接转换为 `Markdown` 格式文本
- 将结果数据更新至`剪贴板`

## 配置
- 将七牛云 AccessKey 和 SecretKey 写入到环境变量 QINIU_ACCESS_KEY、QINIU_SECRET_KEY 中
- 参数形式传入待上传的具体七牛云库的 domain 和 bucket 信息

## 参考
- [Go SDK](https://developer.qiniu.com/kodo/sdk/1238/go)
- [qiniu/api.v7](https://github.com/qiniu/api.v7)
- [package context](https://godoc.org/golang.org/x/net/context)
- [tjgq/clipboard](https://github.com/tjgq/clipboard)
- [package clipboard](https://godoc.org/github.com/tjgq/clipboard)
- [Windows右键菜单设置与应用技巧](http://www.cnblogs.com/russellluo/archive/2011/11/25/2263817.html)
- [命令提示符中同时运行多命令](http://www.45it.com/order/200512/3041.htm)

## 扩展功能
### 右键上传
- 修改注册表
    - 结构
        - `HKEY_CLASSES_ROOT`
            - `*`
                - `shell`
                    - `（╯‵□′）╯︵┴─┴`
                    	- 新增「项」，其中名称为右键菜单要展示的名称
                    	- `command`
                        	- 新增「项」，对应点击后执行的内容
                        		- `默认`
                        			- `cmd.exe /k upload.exe -path=%1 && exit`
                        				- `cmd.exe /k XXXXX` 
                        					- 在 cmd 中执行 `XXXXX` 指令
                        					- `&&` 于前一条指令执行完成后执行
            - `HKEY_CLASSES_ROOT\diretory\shell`
                - 重复上述操作，为文件夹绑定该权限
    - 示例
        - ![微信图片_20170908170222.png](http://otzm88f21.bkt.clouddn.com/425eca95-80af-4988-97fc-a1676e190dd4.png)
        - ![2.png](http://otzm88f21.bkt.clouddn.com/0a666895-6f16-4dc8-8a4b-5493713fc25f.png)


## 测试
- 单文件
```
go run main.go -path=C:\Users\Yao\Desktop\76aaa869ly1fi6n1duxn5j21dw0kak5a.jpg

![76aaa869ly1fi6n1duxn5j21dw0kak5a.jpg][http://otzm88f21.bkt.clouddn.com/04cf05bd-ba52-4dee-8fa9-095555f5c7ec.jpg]
```

- 文件夹
```
go run main.go -path=D:\data\soft\OneDrive\Documents\Write\work\云开

[人员列表.xls][http://otzm88f21.bkt.clouddn.com/a.xls]
[六步搞定实地辅导-20170822.md][http://otzm88f21.bkt.clouddn.com/b.md]
[人员列表.txt][http://otzm88f21.bkt.clouddn.com/c.txt]
```

- Exe 方式调用
    - **异常**
```
go run upload.go -path=C:\Users\Yao\Desktop\76aaa869ly1fi6n1duxn5j21dw0kak5a.jpg

[新建 Microsoft PowerPoint 演示文稿.pptx][http://otzm88f21.bkt.clouddn.com/19118280-cacb-43f3-98b6-32600ef459b9.pptx]
uplaod.exe%!(EXTRA string=http://otzm88f21.bkt.clouddn.com/a932af7f-837a-4df6-92dd-e80b29c25ab6.exe)
```


## 发布
- 配置环境

    ```
    $env:GOPATH = "D:\server\go\lib;D:\work\git\yao\go\Utils_Go";
    ```
    
- 目录跳转
    - `cd D:\work\git\yao\go\Utils_Go\src\upload\qiniu`

- 发布
    - `go build -o upload.exe`