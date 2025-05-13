## UI-Generator

![](./image/logo.svg)

让idea快速成型！让我们帮助你让创意和想法跃然纸上！

## 使用方法

1. 下载Release中的文件
2. 启动ui，目录下会为您生成config.json，编辑config.json，完成端口和llm配置
3. 再次启动ui，打开localhost:端口号，您就可以开始使用ui-generator了

## 编译方法

1. `go build`得到ui可执行文件
2. 前往[前端仓库](https://github.com/dingdinglz/ui-generator-frontend)获取前端部分代码
3. 前端代码中`npm i`，`npm run build`得到dist文件夹
4. 将dist文件夹复制到与ui文件同一目录，并将dist文件夹更名为web
5. 此时，目录下应该有ui可执行文件，web文件夹
6. 按照使用方法启动即可

## 画廊

![](./image/1.png)