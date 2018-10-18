# Yui
_用go语言实现四则运算_

![GitHub](https://img.shields.io/github/license/mashape/apistatus.svg)

下载源代码并编译或者在Release页面直接[下载](https://github.com/RitterHou/yui/releases)已经编译好的文件

测试文件test

    // 测试四则运算
    (1 + -1) * 20 + 50 / 25 * (2.333 - 3.1415926) - (-1) // 计算结果：-0.6171851

打印帮助文档：

    ./yui
  
编译源代码

    ./yui build test
   
运行源代码文件或者字节码文件

    ./yui run test
    ./yui run test.yuicode

反编译字节码文件得到计算指令

    ./yui dec test.yuicode
    
进入交互式的命令行

    ./yui shell

