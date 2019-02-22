# Yui
_用go语言实现四则运算_

![GitHub](https://img.shields.io/github/license/mashape/apistatus.svg)

### 安装
在 Mac 上可以通过 homebrew 进行安装

    brew tap ritterhou/tap && brew install yui

其它平台需要下载源代码并编译或者在Release页面直接[下载](https://github.com/RitterHou/yui/releases)已经编译好的文件

### 使用

打印帮助文档：

    ./yui
  
编译源代码（获取测试文件 [yui.test](https://raw.githubusercontent.com/RitterHou/yui/master/yui.test)）

    ./yui build yui.test
   
运行源代码文件或者字节码文件

    ./yui run yui.test
    ./yui run yui.yuicode

反编译字节码文件得到计算指令

    ./yui dec yui.yuicode
    
进入交互式的命令行

    ./yui shell

如果你已经安装了 rlwrap，那么你可以使用 rlwrap 在输入表达式时进行括号匹配

    rlwrap ./yui shell

[> 视频演示](https://www.bilibili.com/video/av34478667)
