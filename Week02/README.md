### 学习笔记


### 作业：

>
我们在数据库操作的时候，比如 dao 层中当遇到一个 sql.ErrNoRows 的时候，是否应该 Wrap 这个 error，抛给上层。为什么，应该怎么做请写出代码？


#### 答案及想法：

##### 想法：应该抛给上上层，但是不应该wrap

##### 原因：

wrap的副作用如下：

* 1.caller想判定返回的error的类型时，需要强依赖sql.ErrNoRows，也就是依赖sql包。
* 2.强依赖当数据库层做切换时需要到caller层修改

##### 实现方式：

dao层定义noRecordErr，用来统一作为dao层遇到的sql.ErrNoRows时返回的错误，并内部包裹原始的err

在dao层提供IsNoRecordErr方法用来判定是否是noRecordErr类型

在service层封装一层对dao.IsNoRecordErr的Proxy，然后service层和其它caller需要判定时都使用这个proxy

##### 优点：
	    
* 1.service等调用dao的层不直接依赖于sql.ErrNoRows实现解耦
* 2.dao层切换时调用dao的其它地方可以快速切换，改动位置可控

>noRecordErr定义参考net包 https://golang.org/src/net/net.go
