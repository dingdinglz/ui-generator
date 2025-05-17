## UI-Generator

![](./image/logo.svg)

让idea快速成型！让我们帮助你让创意和想法跃然纸上！

## 简介

**UI-Generator**是一款应用原型生成软件，你可以理解为生成一个应用的前端界面，当你在开发一款应用之前，可以在UI-Generator中输入你的idea，简短或详细都可，从`一个古诗学习app`到详细的细致到每一个功能的描述都可行，然后UI-Generator会利用大模型自动为你生成待生成的界面列表，然后一次为您生成html的界面，效果参考可以移步生成效果画廊

## Feature

- UI-Generator带有一个易用、简洁的webui
- 多模型支持：openai系，ollama（TODO）
- 支持生成：web端、手机端
- 操作简单，易于上手，简单的一句话即可作为参照生成一个应用
- 支持历史记录，加载历史生成结果，修改历史生成结果
- UI-Generator带有一个代码编辑框，在生成结束后你可以通过点击对应的文件名，在编辑框中对文件进行修改，也可以加载历史生成结果后进行修改
- 支持直接预览生成结果
- 生成的代码结构清晰，一个界面对应一个html文件

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