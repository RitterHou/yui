# Yui
_用go语言实现四则运算_

![GitHub](https://img.shields.io/github/license/mashape/apistatus.svg)

下载源代码并编译或者在Release页面直接[下载](https://github.com/RitterHou/yui/releases)已经编译好的文件

获取测试文件 [yui.test](https://raw.githubusercontent.com/RitterHou/yui/master/yui.test)

打印帮助文档：

    ./yui
  
编译源代码

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

下面是 shell 的演示操作

[![asciicast](https://asciinema.org/a/K4deAJjyqxECZmpjJOR2rorEt.png)](https://asciinema.org/a/K4deAJjyqxECZmpjJOR2rorEt)
