# Golang调用dll

1. golang调用报错： %1 not a valid win32 application
   * XXX.dll为32位dll，在64位电脑上运行golang默认生成的是64位程序，故需要设置环境变量GOARCH=386

2. syscall.Syscall参数：必须按照原型填满所有参数

```
func Syscall(trap, nargs, a1, a2, a3 uintptr) (r1, r2 uintptr, err Errno)
func Syscall12(trap, nargs, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12 uintptr) (r1, r2 uintptr, err Errno)
func Syscall15(trap, nargs, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15 uintptr) (r1, r2 uintptr, err Errno)
func Syscall6(trap, nargs, a1, a2, a3, a4, a5, a6 uintptr) (r1, r2 uintptr, err Errno)
func Syscall9(trap, nargs, a1, a2, a3, a4, a5, a6, a7, a8, a9 uintptr) (r1, r2 uintptr, err Errno)
```
   * 参数1：dll函数
   * 参数如果包含指针：查询下uintptr

# 参考

* https://golang.org/pkg/syscall/#Syscall
* https://www.cnblogs.com/keanuyaoo/p/3357816.html