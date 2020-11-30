### 学习笔记


#### 作业：

>
我们在数据库操作的时候，比如 dao 层中当遇到一个 sql.ErrNoRows 的时候，是否应该 Wrap 这个 error，抛给上层。为什么，应该怎么做请写出代码？


#### 答案及想法：

dao层定义noRecordErr，用来统一作为dao层遇到的sql.ErrNoRows时的返回值，并在dao层提供IsNoRecordErr方法
用来判定是否是noRecordErr类型，这样的优点：
	    
* 1.service等调用dao的层不直接依赖于sql.ErrNoRows实现解耦
* 2.dao层切换时调用dao的其它地方可以快速切换

>noRecordErr定义参考net包 https://golang.org/src/net/net.go
