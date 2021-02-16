（一）以下是我刚开始几天学习go中的一点总结和感慨
    一、Go语言诞生的背景
        google旗下的大牛，从复杂有庞大的google项目整理编译运行中的到启发而开发的一种静态强类型、编译型语言。
        Go 语言语法与 C 相近，但功能上有：内存安全，GC（垃圾回收），结构形态及 CSP-style 并发计算
    二、Go基础语法
        菜鸟教程 https://www.runoob.com/go/go-tutorial.html
        cn官网文档 https://golang.google.cn/doc/
    三、IDE开发工具
        这里我试着去用了一下 LiteIDE X  和 VS Code。（工具只是辅助开发）
            ① Go专属工具LiteIDE x自带使用文档，功能极其强大，很适合刚刚起步的新手，我开始用起来感觉很顺手（推荐）
            如果对主题有要求的，可以根据个人喜好自己编写\liteide\share\liteide\liteeditor\color\your.xml
            我个人用的liteeditor编辑主题是选择它提供的sumblime-bold.xml。还有就是有关交叉编译的问题（实际就是根据需求gcc【这里是编译器】编译成可以在指定系统下运行的文件exe）
            ② VS code这款支持多语言的插件丰富的工具，我是看cn官网推荐的，所以也试了一下。它主要是依赖于丰富的插件。
            界面相对简约，项目文件展示窗口（move，rename，create）、调试等个人感觉挺好（若是新手的话，个人不推荐使用它，因为我用的时候是配合terminal中的go command使用的）  
            VS code中有关go插件：gopkgs  go-outline  go-symbols  guru  gorename  gotests  gomodifytags  impl  fillstruct  goplay  godoctor  dlv  gocode-gomod  godef  goreturns 
            VS code中项目启动配置是以Jason文件形式：其中program配置的几种情况
                //${fileDirname}  Debug all files in the directory where the current file resides
                //${workspaceRoot} Configure with the package name as the starting point
                //${workspaceFolder} Debugging VS Code opens all the files in the root directory of the workspace
                //${file} Debug the current file
                "program": "${fileDirname}"
    四、Go使用（Windows）
        I go env 即go环境，这里我主要是参考cn官网边看便体会的。安装go之后有个gopath的概念，简单来看就是一个为了方便我们便是的workspace（包含bin、pkg、src这个顾名思义）
        Ⅱ go项目开发时，需要一些像数据库驱动、外部模块包等等依赖时，就需要用到go get 指令，但是有些依赖没办法下载（国外），所以我们把下载的代理修改一下。修改指令go env -u GO111MODULE=on
            go env -u GORPOXY=https://goproxy.io,direct  修改完后就个以用go get下载依赖了，例如 go get -u github.com/go-sql-driver/mysql下载mysql驱动
        Ⅲ 打开LiteIDE x打开我们的gopath目录,在src下建立自己的项目HelloWorld/hello.go  这里是在不引入外部包和外部依赖的情况下，只是用go的标准库的fmt的情况，我们直接点击Build and Run
        Ⅳ 有关多package的情况：
               ① 外网的一些package，例如rsc.io/quote
               ② 自定义的一些package 
                    1、.go文件之间得调用（同包下可以直接调用） 不同包就请看下面2中内容
            	    2、package、module：其实在go mod init初始化module的时候，单个看来是一个module，但是对于其他调用者来说就是一个package。
                    如果本地初始化但未发布到Internet的module可以通过在调用者的go.mod中replace  modulename => currentdir用指明路径
	                3、go path在run 和build (go path 就是指定的一个工作空间，方便我们下载依赖（模块包，如go get -u github.com/go-sql-driver/mysql），build是编译包和依赖)
    五、感慨
        语言语法上手快，代码编写简洁（简写模式可以了解一下，有助于提高效率）、有关并发编程中goroutine协程配合channel简直爽歪歪。在实际开发当中，要注意的地方也还是有多。我踩过的坑也不少。
        切莫心浮气躁，多看文档多练习几遍就Ok了。  Practice makes perfect！


（二） main.go在multiple里面
