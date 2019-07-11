
## mysql 查询过程分析

`select city,name,age from t where city='杭州' order by name limit 1000;`

1、初始化sort_buffer,确定放入 city,name,age 这个三个字段

2、从索引city找到第一个满足 city='杭州'条件的主键id，也就是图中的ID_x

3、到主键id 索引取出整行，取出 city,name,age 这个三个字段，存入sort_buffer中

4、从索引city 取下一个记录的主键id

5、重复步骤3，4直到city的值不满足查询条件位止，对应的主键id 也就是图中的id_y

6、对sort_buffer 中的数据按照字段name 做快速排序

7、按照排序结果取前1000 行返回给客户端


### 比如一张表foo 有主键id和索引bar 
`select * from foo where bar > 10 order by id limit 10`

1、初始化sort_buffer, 确定放入所有查找字段

2、从索引bar找到第一个满足 bar > 10 条件的主键id

3、到主键id 索引取出整行，取出所有查找字段值，存入sort_buffer中

4、从索引bar 取下一个记录的主键id

5、重复步骤3，4直到bar的值不满足查询条件位止，对应的主键id

6、对sort_buffer 中的数据按照字段id做快速排序

7、按照排序结果取前10 行返回给客户端