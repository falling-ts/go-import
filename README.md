## Go 包的导入顺序

这是关于 Go 包复杂导入时, 导入顺序的学习

### 下面是执行的结果

```shell
f
e
c
g
d
a
h
b
main
exec main
```

### 代码示例

```go
package a

import (
	"fmt"
	_ "import/c"
	_ "import/d"
)

var name = "a"

func init() {
	fmt.Println(name)
}
```

### 关于网络上的导入图示

![图片](https://tse3-mm.cn.bing.net/th/id/OIP-C.vdC_CSJUyoIEyqqf4wNPuAHaCq?pid=ImgDet&rs=1)

### 结论

以项目为例, 导入顺序如下

> main->a->c->e->f
> 
> d->g
> 
> b->h

因此 init 执行顺序为

> f->e->c->g->d->a->h->b->main->exec main

```shell
注: 修改导入顺序后, init 执行顺序也跟着变化了
```