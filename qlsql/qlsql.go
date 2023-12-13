package qlsql

import (
	"database/sql"
	"fmt"

	"github.com/sirupsen/logrus"
)

// Add_Table_field 添加字段  表 名称 类型 默认值 备注
func Add_Table_field(db sql.DB, table, name, typ string) error {
	sqlstr := fmt.Sprintf("ALTER TABLE %s ADD COLUMN %s "+typ, table, name)
	//sqlstr := "ALTER TABLE " + table + " ADD COLUMN "+name+" int(11) DEFAULT NULL COMMENT '年龄'"
	_, err := db.Exec(sqlstr)
	if err != nil {
		logrus.Error(err)
		return err
	}
	return nil
}

// Edit_Table_field 编辑字段  表 名称1 to 名称2 类型
func Edit_Table_field(db sql.DB, table, name1, name2, typ string) error {
	sqlstr := fmt.Sprintf("ALTER TABLE %s CHANGE %s %s %s", table, name1, name2, typ)
	//sqlstr := "ALTER TABLE tb_emp1 CHANGE col1 col3 CHAR(30)"
	_, err := db.Exec(sqlstr)
	if err != nil {
		fmt.Printf("%s failed, err:%v\n", sqlstr, err)
		return err
	}
	return nil
}

// Del_Table_field 删除字段
func Del_Table_field(db sql.DB, table, name string) error {
	sqlstr := fmt.Sprintf("ALTER TABLE %s DROP %s", table, name)
	_, err := db.Exec(sqlstr)
	if err != nil {
		logrus.Error(err)
		return err
	}
	return nil
}

//ALTER TABLE tkaqkcz add unique 法法 (FgSM8lo, FI3iWrf)
//ALTER TABLE tkaqkcz DROP INDEX 法法

//Addfield_UNIQUE unique_keys
func Addfield_UNIQUE(db sql.DB, Tname, fields, key string) error {
	sqlstr := "ALTER TABLE %s ADD CONSTRAINT %s UNIQUE(%s)"
	sqlstr = fmt.Sprintf(sqlstr, Tname, key, fields)
	_, err := db.Exec(sqlstr)
	if err != nil {
		logrus.Error(err)
		return err
	}
	return nil
}

func Delfield_UNIQUE(db sql.DB, Tname, key string) error {
	sqlstr := "ALTER TABLE %s DROP INDEX %s"
	sqlstr = fmt.Sprintf(sqlstr, Tname, key)
	_, err := db.Exec(sqlstr)
	if err != nil {
		logrus.Error(err)
		return err
	}
	return nil
}

/* 	设置默认值
ALTER TABLE tb_dept3
    -> CHANGE COLUMN location
	-> location VARCHAR(50) DEFAULT 'Shanghai';

指定部门位置不能为空
ALTER TABLE tb_dept4
    -> CHANGE COLUMN location
	-> location VARCHAR(50) NOT NULL;

删除非空约束
	ALTER TABLE tb_dept4
-> CHANGE COLUMN location
-> location VARCHAR(50) NULL;

http://c.biancheng.net/view/7617.html
*/

/*
https://www.cnblogs.com/sweet521/p/6203360.html
创建索引
ALTER TABLE table_name ADD INDEX index_name (column_list)

ALTER TABLE table_name ADD UNIQUE (column_list)

ALTER TABLE table_name ADD PRIMARY KEY (column_list)

删除索引
DROP INDEX index_name ON talbe_name

ALTER TABLE table_name DROP INDEX index_name

ALTER TABLE table_name DROP PRIMARY KEY
*/
