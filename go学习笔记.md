# go 学习笔记

## 1 common 随记

1. `goimports` 可以根据代码需要, 自动地添加或删除`import`声明。这个工具并没有包含在标准的分发包中，可以用下面的命令安装：`go get golang.org/x/tools/cmd/goimports`

2. 对于`string`类型，`+ `运算符连接字符串

3. Printf 格式

    1. | 格式       | 备注                                            |
        | ---------- | ----------------------------------------------- |
        | %d         | 十进制数据                                      |
        | %x, %o, %b | 十六进制，八进制，二进制整数。                  |
        | %f, %g, %e | 浮点数：3.141593 3.141592653589793 3.141593e+00 |
        | %t         | 布尔（true，false）                             |
        | %c         | 字符（rune）（Unicode码点）                     |
        | %s         | 字符串                                          |
        | %q         | 字符串（带双引号）、字符（带单引号）            |
        | %v         | 变量的自然醒时（natual format）                 |
        | %T         | 变量的类型                                      |
        | %%         | 百分号                                          |

4. `go list`命令的作用是列出指定的代码包的信息。

5. 使用`go test`的`-args`标签，该标签会把其后的所有字符串当做参数传入

6. go 编译速度快的 3 个原因

    1. 第一点，所有导入的包必须在每个文件的开头显式声明，这样的话编译器就没有必要读取和分析整个源文件来判断包的依赖关系。
    2. 第二点，禁止包的环状依赖，因为没有循环依赖，包的依赖关系形成一个有向无环图，每个包可以被独立编译，而且很可能是被并发编译。
    3. 第三点，编译后包的目标文件不仅仅记录包本身的导出信息，目标文件同时还记录了包的依赖关系。

7. go 包 package

    1. 默认包名一般采用导入路径名的`最后一段`的约定
        1. 例外1，包对应一个可执行程序，也就是main包
        2. 例外2，包所在的目录中可能有一些文件名是以`_test.go`为后缀的Go源文件
        3. 例外3，一些依赖版本号的管理工具会在导入路径后追加版本号信息，例如"gopkg.in/yaml.v2"。这种情况下包的名字并不包含版本号后缀，而是yaml。

8. go 源码文件

    1. 命令源码文件 ：含 main 的 go 文件，package 声明为 main
    2. 库源码文件
    3. 测试源码文件：以 _test.go 结尾的，里面函数必须包含 Test 开头或者 Benchmark 开头

9. go run 命令只能接受一个命令源码文件以及若干个库源码文件（必须同属于 main 包）作为文件参数，且**不能接受测试源码文件**。

10. 建超：在使用模块的时候，GOPATH 是无意义的，不过它还是会把下载的依赖储存在 $GOPATH/src/mod 中，也会把 go install 的结果放在 \$GOPATH/bin 中。

11. golang 包 package

    1. 两种 golang 安装第三方包的方法 [文档地址](<https://my.oschina.net/u/3045933/blog/1861700>)

        1. go get 安装
        2. 源码包安装【手动下载源码包，拷贝到 $GOPATH/src/ 目录下】
            ```sh
            # 比如要安装"github.com/golang/protobuf/proto"
            # 去github.com/golang/protobuf下载源码包，
            # 拷贝到 $GOPATH/src/github.com/golang/protobuf

            $ cd $GOPATH/src/github.com/golang/protobuf
            $ go install
            ```

    2. 使用第三方包
       1. 直接在源码里 import 即可
       ```sh
       // 比如要使用  "github.com/gin-gonic/gin"
       import "github.com/gin-gonic/gin"
       ```

       2. `注意`  golang 的所有 import 的包必须在 \$GOPATH 路径下，如果直接导入源码路径下的包编译就会报错。因为 golang 只会在 \$GOPATH 下找包的依赖，而不会在当前工程下找。
       3. 理解 golang 包的导入 [参考文档](<https://www.cnblogs.com/sevenyuan/p/4548748.html>)
           1. 在使用第三方包的时候，实际上是 link 了以该源码编译的临时目录的`.a文件`
           2. import 后面的最后一个元素应该是路径，即目录，并非包名
           3. fmt.Println() 中的 fmt 是包名
           4. import m "lib/math" 中的 m 是包名

12. go mod 命令

    1. 查看帮助 `go help mod` `go help modules` `go help go.mod`

    ```yaml
    download    下载依赖的module到本地cache
    edit        编辑go.mod文件
    graph       打印模块依赖图
    init        再当前文件夹下初始化一个新的module, 创建go.mod文件
    tidy        增加丢失的module，去掉未用的module
    vendor      将依赖复制到vendor下
    verify      校验依赖
    why         解释为什么需要依赖
    ```

    1. 包含 go.mod 的目录被称为 模块根（module root）

13. 理解 golang 的 make 和 new 的区别 [参考文章](<https://sanyuesha.com/2017/07/26/go-make-and-new/>)

    1. new(T) 返回的是 T 的指针，即 *T，并初始化为 T 的零值
    2. make 只能用于 slice、map 和 channel
    3. 被 make 初始化的值跟直接初始化是不同的，如 slice 的初始化是 nil，但被 make 初始化的值跟类型相关 ，如果是 int 则为 0，map 和 channel 也类似如此 。
    4. make(T, args) 返回的是 T 的引用
    5. 好少需要使用 new

14. fmt.Printf **格式化动作（verb）**

```yaml
# 通用
%v	值的默认格式表示
%+v	类似%v，但输出结构体时会添加字段名
%#v	值的Go语法表示
%T	值的类型的Go语法表示

# 布尔值
%t	单词true或false

# 整数
%b	表示为二进制
%c	该值对应的unicode码值
%d	表示为十进制
%o	表示为八进制
%q	该值对应的单引号括起来的go语法字符字面值，必要时会采用安全的转义表示
%x	表示为十六进制，使用a-f
%X	表示为十六进制，使用A-F
%U	表示为Unicode格式：U+1234，等价于"U+%04X"

# 字符串和 []byte
%s	直接输出字符串或者[]byte
%q	该值对应的双引号括起来的go语法字符串字面值，必要时会采用安全的转义表示
%x	每个字节用两字符十六进制数表示（使用a-f）
%X	每个字节用两字符十六进制数表示（使用A-F） 

# 指针
%p	表示为十六进制，并加上前导的0x    

```

15. ## make 命令，makefile 文件编写 ？？

16. golang 的 reflect 包

    1. 获取 Value 对象： `reflect.ValueOf(s)`
    2. reflect.TypeOf(s) 返回 动态类型

17. interface 理解 [参考文章](<https://sanyuesha.com/2017/07/22/how-to-understand-go-interface/>)

    1. interface 是一种类型，具有一组方法的类型，不带任何方法的 interface 叫 empty interface，如果一个类型实现了一个 interface 中的所有方法，我们说该类型实现了该 interface，所以所有类型都实现了 empty interface
    2. interface 存储的是实现者的值：如果有多个类型实现了某个 interface，这些类型的值都可以直接使用 interface 的变量存储
    3. 如何判断 interface 变量存储的哪种类型： go 可以使用 comma, ok 的形式做区分 `value, ok := em.(T)`：em 是 interface 类型的变量，T代表要断言的类型，value 是 interface 变量存储的值，ok 是 bool 类型表示是否为该断言的类型 T。
    4. empty interface ：go 不会对类型是 interface{} 的 slice 进行转换

18. 在 golang 中，可以为一个类型编写名为 String 的方法，用以表示该类型的字符串表示形式。这个 String 不需要声明任何参数，但需要有一个 string 类型的结果声明。在 fmt.Print 的时候，无需显式的调用 String 方法，fmt.Print 函数会自己去寻找它。

19. 日志相关
    [参考文章](<https://www.flysnow.org/2017/05/06/go-in-action-go-log.html>)

    ```go
    func init() {
        log.SetPrefix("[INFO]") //设置前缀
        log.SetFlags(log.Flags() | log.Lshortfile) //设置文件名及行数
    }
    log.Fatal //相当于 log.Print，再 os.Exit(1)
    ```

20. delve debug main.go 出现如下错误
    1. `could not launch process: listen tcp: lookup localhost: no such host`
    2. 解决办法：在 hosts 文件里添加 `127.0.0.1 localhost` ，该问题同时也导致 goland ide 调试不成功。

21. 零值

    | 类型             | 零值               |
    | ---------------- | ------------------ |
    | int float32 rune | 0                  |
    | string           | 空字符串           |
    | interface        | nil                |
    | struct           | {}                 |
    | 指针             | nil                |
    | 数组             | 按数组的类型初始化 |
    | slice （切片 ）  | []                 |
    | map              | map[]              |

    

## 2. 编码规范

1. 命名
    1. 变量命令尽量用驼峰命名，不要用下划线，多个变量命名放一起
    2. 在函数外部用 var 命名，不要 采用 := 命名，容易踩坑
    3. package 名字尽量与目录名一致

2. 尽量处理任何有 err 返回的调用，不要用 `_` 丢弃，实在不行就 panic，或者用 log 记录之

3. 闭包调用

    **在循环中调用函数或者 goroutine 方法，一定要采用显式的变量调用，不要在闭包函数里调用循环的参数如**

    ```go
    for i:=0; i<10; i++ {
        go func() {DoSomething(i)} //错误的做法
        go func(i int) {DoSomething(i)}(i) //正确的做法
    }
    ```

4. 在 main 包里只有当如`文件无法打开、数据库无法连接导致程序无法正常运行`的请开启用 panic，对于其他的 package 包对外的接口不能有 panic，只能在包内采用。

5. struct 规范

    1. struct 声明和初始化采用多行

         ```go
        // 定义
        type User struct {
            Username string
        }
        // 初始化
        u := User {
            Username: "zhangsan"
        }
        ```
     2. received 是指类型还是指针类型，参考如下原则

        ```go
        // case 1: w 不会有任何改变（值类型）
        func (w Win) Tally(playerPlayer) int 
        // case 2: w 会改变数据 （指针类型）
        func (w *Win) Tally(playerPlayer) int 
        //case 3: 带 mutex 的 struct 必须是指针 receivers
        
        ```

        

## 3. 结构体相关随笔记录

1. 结构体可以不包含任何字段，这也是有意义的，因为我们可以为类型关联上**方法**

    >方法需要有名字，不能被当作值来看待，更重要的是：**它必须隶属于某个类型**。方法所属的类型会通过其声明中的**接受者（receiver）**声明体现出来。
    >
    >*接受者的类型就是当前方法所属的类型；接受者的名称，则用于当前方法中所引用它所属的类型的当前值*

2. 方法隶属于类型并不局限于结构体类型，但必须是某个自定义的数据类型，并且不能是任何接口类型。

3. 如果结构体类型的某个字段声明只有一个类型名，那么该字段代表了什么？(摘自极客时间）

    > **这个问题的典型回答是：**字段声明AnimalCategory代表了Animal类型的一个嵌入字段。Go 语言规范规定，如果一个字段的声明中只有字段的类型名而没有字段的名称，那么它就是一个嵌入字段，也可以被称为匿名字段。我们可以通过此类型变量的名称后跟“.”，再后跟嵌入字段类型的方式引用到该字段。也就是说，嵌入字段的类型既是类型也是名称。

4. *知识点*：如果处于统一层级的多个嵌入字段拥同名的字段或方法，调用被嵌入类型的值的时候，如果调用了该同名字段或方法时，编译器则会报一个编译错误：`ambiguous selector `，很好理解，因为编译器不知道该选哪个

5. golang 里没有继承的概念，只是通过嵌入字段实现了类型的**组合**，并不需要显式的声明某个类型实现了某个接口。*接口之间的组合更加常见*

## 4. 函数相关随笔记录 

1. 函数是一等（first-class）公民，函数类型也是一等的数据类型

## 5. 接口（interface）相关随笔记录

1. 接口无法被实例化（值化）
2. 接口包裹的是方法定义，结构体包裹的是字段声明
3. 对于任何类型，只要它的方法集合完全包含了一个接口的全部方法，那么它就一定是这个接口的实现类型。
4. **一条通用的规则：如果我们使用一个变量给另外一个变量赋值，那么真正赋值给后者的，并不是前者持有的那个值，而是该值的一个副本。**
5. 当我们给一个接口变量赋值的时候，该变量的动态类型会与它的动态值一起被存储在一个专用的数据结构（iface）中。


