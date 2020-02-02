## 数据存储
genedata/dump.go 爬取按年级搜索的json 数据 ,可设置cron 定时抓取到本地

redis 按照年级id 存储课程数据(json)
  

## api 
   基于gin 开发api 
   
   按年级  从redis 读数据 , html/template 渲染html，返回
   
## 部署 
 ```apple js
    作为daemon 运行 
    sh run.sh 
```

