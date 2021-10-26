###### 示例程序
[import](../)

###### import

Go通过import导入Go标准库以及外部的包

使用方法：

- 单行导入（一个包）

场景：当导入一个包时
```
//导入系统格式化库文件
import "fmt"

//使用fmt格式化输出
fmt.Println("hello world")
```
输出：
hello world

- 单行导入（多个包）

场景：当导入多个包时
```
//导入系统格式化库文件
import "fmt"
//导入时间库文件
import "time"

func main() {
    //使用fmt格式化输出
	fmt.Println("hello world")
	//输出当前时间
	fmt.Println(time.Now())
}
```
输出：
hello world
2021-10-10 13:30:36.3718097 +0800 CST m=+0.006640701

- 多行导入

场景：当导入多个包时省略多个import命令
```
import (
    
    //导入系统格式化库文件
	"fmt"
	//导入时间库文件
	"time"
)

func main() {
	fmt.Println("hello world")
	fmt.Println(time.Now())
}
```
输出：
hello world
2021-10-10 13:30:36.3718097 +0800 CST m=+0.006640701

- 点操作

场景：当想省略使用包时的包名前缀时，使用点操作
```
import . "fmt"

func main() {
	Println("hello world")
}
```
输出：
hello world
使用 . 点操作使用在后续使用fmt包内定义的方法时，省略掉 `fmt.` 这样的前缀
TODO: . 的其他操作还有那些？

- 别名操作

场景：当把包名更换为另一个（容易记忆的）名字 或者 导入多个包但是包名冲突（相同）
```
import fmtAlias "fmt"

func main() {
fmtAlias.Println("hello world")
}
```
输出：
hello world

- _ 操作
场景：当只想引入包执行包内的init()函数，做一些初始化操作，而并非需要其他的方法，可以使用 _ 操作

上面的示例都是引入Go标准库，引入本地项目的包以及外部第三方的包

```
import (
	"fmt"
	"github.com/sirupsen/logrus"
	"import/lib"
)

func main() {
	fmt.Println("hello world")
	lib.Example()
	logrus.Debug("debug")
}
```
上面的代码引入Go标准库fmt，引入了自己本地项目实现的import/lib包，引入了github上面的日志包logrus

引入了包，那么包在哪里? 是如何查找到的？，就有了包管理的概念。
###### 包管理

- GOPATH模式
在GOPATH下查找包按照GOROOT和多GOPATH目录下的src/xxx依次查找
- GOMOD模式
在GOMOD下查找包，mod包名就是包的前缀，里面的目录就是后续的路径。
在GOMOD模式下，查找包就不去GOPATH查找，只是GOMOD包缓存在gopath/pkg/mod里面。

GOPATH模式 是旧的方式，缺点 TODO:
目前推荐GOMOD模式


https://zhuanlan.zhihu.com/p/142176662
https://www.cnblogs.com/foxhappy/p/14547415.html
https://www.cnblogs.com/Console-LIJIE/p/14855055.html







