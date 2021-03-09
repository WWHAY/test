[toc]
# 栈
## 特点
先入后出：可以联想成蒸包子的现象
## go实现栈
### 入栈
golang有一个自带的函数append()，可以实现入栈的操作
### 出栈
其实出栈的话，就利用了golang的切片了，因为是后入先出，所以只需要截取切片就可以了
### 代码实现
#### 题目描述--leetcode1047
```
给出由小写字母组成的字符串 S，重复项删除操作会选择两个相邻且相同的字母，并删除它们。

在 S 上反复执行重复项删除操作，直到无法继续删除。

在完成所有重复项删除操作后返回最终的字符串。答案保证唯一。

```
#### 代码
```
func removeDuplicates(S string) string {
	res := []byte{}
	for i := range S {
		if len(res) != 0 && res[len(res)-1] == S[i] {
			res = res[:len(res)-1]
		} else {
			res = append(res, S[i])
		}
	}
	return string(res)
}
```
## 写在最后的话
网上关于栈的内在原理有很多，所以我就没有很详细的写，所以以上是我简单的理解