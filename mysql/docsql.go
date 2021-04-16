package sql

/*
普通SQL语句执行过程：
	客户端对SQL语句进行占位符替换得到完整的SQL语句。
	客户端发送完整SQL语句到MySQL服务端
	MySQL服务端执行完整的SQL语句并将结果返回给客户端。

预处理执行过程：
	把SQL语句分成两部分，命令部分与数据部分。
	先把命令部分发送给MySQL服务端，MySQL服务端进行SQL预处理。
	然后把数据部分发送给MySQL服务端，MySQL服务端对SQL语句进行占位符替换。
	MySQL服务端执行完整的SQL语句并将结果返回给客户端。

为什么要预处理？
	优化MySQL服务器重复执行SQL的方法，可以提升服务器性能，提前让服务器编译，一次编译多次执行，
		节省后续编译的成本。
	避免SQL注入问题。
*/

/*
事务：一个最小的不可再分的工作单元；通常一个事务对应一个完整的业务(例如银行账户转账业务，该业务就是
	一个最小的工作单元)，同时这个完整的业务需要执行多次的DML(insert、update、delete)语句共同联
	合完成。A转账给B，这里面就需要执行两次update操作。
在MySQL中只有使用了Innodb数据库引擎的数据库或表才支持事务。事务处理可以用来维护数据库的完整性，保
	证成批的SQL语句要么全部执行，要么全部不执行。
*/
