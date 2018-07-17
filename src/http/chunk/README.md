

golang POST data using the Content-Type multipart/form-data
https://stackoverflow.com/questions/20205796/golang-post-data-using-the-content-type-multipart-form-data
golang几种post方式
https://www.cnblogs.com/zhangqingping/p/4598337.html





bodyBuf := bytes.NewBufferString("")
    bodyWriter := multipart.NewWriter(bodyBuf)

    //建立文件的http第一部分数据,文件信息
    _, err := bodyWriter.CreateFormFile(paramName, path)
    if err != nil {
        return nil, err
    }

    //读取文件,当做http第二部分数据
    fh, err := os.Open(path)
    if err != nil {
        return nil, err
    }
    //mulitipart/form-data时,需要获取自己关闭的boundary
    boundary := bodyWriter.Boundary()
    closeBuf := bytes.NewBufferString(fmt.Sprintf("\r\n--%s--\r\n", boundary))

    //建立写入socket的reader对象
    requestReader := io.MultiReader(bodyBuf, fh, closeBuf)

    fi, err := fh.Stat()
    if err != nil {
        return nil, err
    }
    req, err := http.NewRequest("POST", uri, requestReader)
    if err != nil {
        return nil, err
    }
    //设置http头
    req.Header.Add("Content-Type", "multipart/form-data; boundary="+boundary)
    req.ContentLength = fi.Size() + int64(bodyBuf.Len()) + int64(closeBuf.Len())