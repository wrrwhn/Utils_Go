package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"utils"

	"github.com/google/uuid"
	"github.com/qiniu/api.v7/auth/qbox"
	"github.com/qiniu/api.v7/storage"
	"github.com/tjgq/clipboard"
)

var ak string
var sk string
var domain string
var bucket string
var path string
var level int

func init() {

	ak = os.Getenv("QINIU_ACCESS_KEY")
	if len(ak) == 0 {
		flag.StringVar(&ak, "ak", "", "access key")
	}
	sk = os.Getenv("QINIU_SECRET_KEY")
	if len(sk) == 0 {
		flag.StringVar(&sk, "sk", "", "secret key")
	}
	domain = os.Getenv("QINIU_DOMAIN")
	if len(domain) == 0 {
		flag.StringVar(&domain, "domain", "", "domain")
	}
	bucket = os.Getenv("QINIU_BUCKET")
	if len(bucket) == 0 {
		flag.StringVar(&bucket, "bucket", "", "bucket")
	}
	flag.StringVar(&path, "path", "", "path")
	if len(path) == 0 {
		path = utils.GetCurrentPath()
	}
	flag.IntVar(&level, "level", 1, "path level")

	flag.Parse()
}

func main() {

	// init
	check()
	mac, uploader := initUploader()
	files, filenames, formats := filterPath(path)

	// init.params
	var res string
	length := len(*files)
	c := make(chan string, length)

	// upload
	for idx, localFile := range *files {
		go upload(mac, uploader, &c, localFile, (*filenames)[idx], (*formats)[idx])
	}

	// clipboard.copy
	for i := 0; i < length; i++ {
		res += (<-c + "\n")
	}
	log.Println(res)
	clipboard.Set(res)
}

func upload(mac *qbox.Mac, uploader *storage.ResumeUploader, c *chan string, localFile, filename, format string) {

	// init.token
	ret := storage.PutRet{}
	putPolicy := storage.PutPolicy{
		Scope: bucket,
	}
	token := putPolicy.UploadToken(mac)
	// init.key
	var key = uuid.New().String() + localFile[strings.LastIndex(localFile, "."):]

	err := uploader.PutFile(context.Background(), &ret, token, key, localFile, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	*c <- fmt.Sprintf(format, filename, storage.MakePublicURL(domain, key))
}

func check() {

	errFmt := "Fail: Param[%v] lose"
	if len(ak) == 0 {
		log.Panicf(errFmt, "access key")
	}
	if len(sk) == 0 {
		log.Panicf(errFmt, "secret key")
	}
	if len(domain) == 0 {
		log.Panicf(errFmt, "domain")
	}
	if len(bucket) == 0 {
		log.Panicf(errFmt, "bucket")
	}
}

func initUploader() (mac *qbox.Mac, uploader *storage.ResumeUploader) {

	mac = qbox.NewMac(ak, sk)

	cfg := storage.Config{}
	cfg.Zone = &storage.ZoneHuanan
	cfg.UseHTTPS = false
	cfg.UseCdnDomains = false

	uploader = storage.NewResumeUploader(&cfg)

	return
}

func filterPath(path string) (files *[]string, filenames *[]string, formats *[]string) {

	// init
	files = &[]string{}
	filenames = &[]string{}
	formats = &[]string{}

	// filter
	mode, err := utils.CheckMode(path)
	if nil != err {
		log.Panic(err.Error())
	}
	switch mode {
	case utils.FileModeFormat:
		filterFile(path, files, filenames, formats)
	case utils.DicModeFormat:
		filterDic(path, files, filenames, formats)
	case utils.ErrorModeFormat:
		log.Panicf("fail to check path[%v]", path)
	}

	return
}

func filterFile(path string, files *[]string, filenames *[]string, formats *[]string) {

	if len(path) == 0 {
		log.Panic("filterFile.path is blank")
	}

	*files = append(*files, path)
	typeFormat, err := utils.CheckType(path)
	if nil != err {
		log.Printf("fail to get type for path[%v]", path)
	}
	basename, err := utils.GetBasename(path)
	if nil != err {
		log.Printf(err.Error())
		basename = "."
	}
	*filenames = append(*filenames, basename)
	*formats = append(*formats, string(typeFormat))
}

func filterDic(path string, files *[]string, filenames *[]string, formats *[]string) {

	if len(path) == 0 {
		log.Panic("filterDic.path is blank")
	}

	split := "\\"
	curLevel := strings.Count(path, split)
	_ = filepath.Walk(path, func(p string, file os.FileInfo, err error) error {
		if nil == file {
			return nil
		}
		if strings.Count(p, split) > curLevel+level {
			return nil
		}

		mode, err := utils.CheckMode(p)
		if nil != err {
			log.Printf("fail to get path[%v].mode", p)
		}
		switch mode {
		case utils.FileModeFormat:
			filterFile(p, files, filenames, formats)
		}

		return nil
	})
}
