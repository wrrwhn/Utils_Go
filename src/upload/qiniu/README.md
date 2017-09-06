
## 描述
- 将指定 [文件|文件夹|当前文件夹] 的内容上传至 `七牛云`
- 上传完毕后将链接转换为 `Markdown` 格式文本
- 将结果数据更新至`剪贴板`

## 参考
- [Go SDK](https://developer.qiniu.com/kodo/sdk/1238/go)
- [qiniu/api.v7](https://github.com/qiniu/api.v7)
- [package context](https://godoc.org/golang.org/x/net/context)
- [tjgq/clipboard](https://github.com/tjgq/clipboard)
- [package clipboard](https://godoc.org/github.com/tjgq/clipboard)

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