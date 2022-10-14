package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
)

func Intojson(i interface{}) {
	data, err := json.Marshal(i) // map本身为引用。
	if err != nil {
		fmt.Println("Marshal err:", err)
		return
	}
	fmt.Println("===============================================")
	fmt.Println("map序列化后 = ", string(data))
}

func dbDemo() {
	c := g.Cfg()
	fmt.Println(c.SetFileName("configdemo.toml"))
	db := g.DB("defaultdemo")
	// ================链式模型=======================
	// 默认情况下，gdb是非链式安全的 unsafe
	m1 := db.Table("s_enum")
	m1.Where("enum_key in (?) ", g.Slice{"CMDB_FIXED_HEADER"})
	r1, err := m1.All()
	fmt.Println(r1, err)

	m1.Where("enum_id in (?)", g.Slice{"1"})
	r2, err := m1.All()
	fmt.Println(r2, err)
	fmt.Println("==========================unsafe=====================")

	// 调动Clone方法克隆当前模型
	m2 := db.Table("s_enum")
	c1 := m2.Clone()
	c1.Where("enum_key in (?)", g.Slice{"CMDB_FIXED_HEADER"})
	r3, err := c1.Count()
	fmt.Println(r3, err)

	c2 := m2.Clone()
	c2.Where("enum_id in (?)", g.Slice{"1"})
	r4, err := c2.All()
	fmt.Println(r4, err)
	fmt.Println("==========================clone==============================")

	// Safe方法设置当前模型为链式安全的对象
	s := db.Table("s_enum").Safe()
	s1 := s.Where("enum_key in (?)", g.Slice{"CMDB_FIXED_HEADER"})
	res, err := s1.All()
	fmt.Println(res, err)

	s2 := s.Where("enum_id = ?", g.Slice{"1"})
	res1, err := s2.All()
	fmt.Println(res1, err)
	fmt.Println("====================safe===========================")
	// ===============链式操作=================
	result1, err := db.Table("app_develop").Data(g.Map{"app_id": 1, "develop_name": "dev1",
		"create_time": time.Now(),
		"update_time": time.Now(),
	}).Insert()
	fmt.Println(result1, err)
	result2, err := db.Table("app_develop").Insert(g.Map{"app_id": 2, "develop_name": "dev2",
		"create_time": time.Now(),
		"update_time": time.Now(),
	})
	fmt.Println(result2, err)
	result3, err := db.Table("app_develop").Insert(g.List{
		{"app_id": 3, "develop_name": "dev3", "create_time": time.Now(), "update_time": time.Now()},
		{"app_id": 4, "develop_name": "dev4", "create_time": time.Now(), "update_time": time.Now()},
		{"app_id": 5, "develop_name": "dev5", "create_time": time.Now(), "update_time": time.Now()},
		{"app_id": 6, "develop_name": "dev6", "create_time": time.Now(), "update_time": time.Now()},
	})

	fmt.Println(result3, err)
	result4, err := db.Table("app_develop").Data(g.Map{"develop_name": "dev1upd1"}).Where("app_id", 1).Update()
	fmt.Println(result4.RowsAffected())

	result5, err := db.Table("app_develop").Update(g.Map{"develop_name": "dev2upd2"}, "app_id", 2)
	fmt.Println(result5.RowsAffected())

	result6, err := db.Table("app_develop").Where("app_id", 5).Delete()
	fmt.Println(result6.RowsAffected())

	result7, err := db.Table("app_develop").Delete("app_id", 6)
	fmt.Println(result7.RowsAffected())

	type AppDevelop struct {
		DevelopId     int `orm:"develop_id"`
		AppId         int `gconv:"app_id"`
		DevelopName   string
		DevelopRemark string
		DelFlag       int
		CreateTime    string
		UpdateTime    string
	}
	dev := []AppDevelop{}
	err = db.Table("app_develop").Where("develop_name", "dev2upd2").Structs(&dev)
	fmt.Println(dev)

	// ===============普通操作=================
	// 数据查询
	list, _ := db.GetAll("select * from s_enum where enum_key in (?) limit 1", g.Slice{"CMDB_FIXED_HEADER"})
	Intojson(list.List())
	// 数据插入1 如果写入的数据中存在主键或者唯一索引时，返回失败
	i, err := db.Insert("app_attr", gdb.Map{
		"app_attr_id": "1",
	})
	fmt.Println(err)
	fmt.Println(i.LastInsertId())
	// 数据插入2 如果写入的数据中存在主键或者唯一索引时，更新原有数据
	is, err := db.Save("app_attr", gdb.Map{
		"app_attr_id": "2",
	})
	fmt.Println(err)
	fmt.Println(is.LastInsertId())
	// 数据插入3
	bi, err := db.BatchInsert("app_attr", gdb.List{
		{"app_attr_id": "3"},
		{"app_attr_id": "4"},
		{"app_attr_id": "5"},
		{"app_attr_id": "6"},
	}, 10)
	fmt.Println(err)
	fmt.Println(bi.LastInsertId())
	// 数据更新1
	u1, err := db.Update("app_attr", gdb.Map{"attr_key": "test1"}, "app_attr_id=?", 1)
	fmt.Println(err)
	fmt.Println(u1.LastInsertId())
	// 数据更新2
	u2, err := db.Update("app_attr", "attr_key='test2'", "app_attr_id=2")
	fmt.Println(err)
	fmt.Println(u2.LastInsertId())

	// 数据更新3
	u3, err := db.Update("app_attr", "attr_key=?", "app_attr_id=?", "test3", 3)
	fmt.Println(err)
	fmt.Println(u3.LastInsertId())

	// 数据删除
	d1, err := db.Delete("app_attr", gdb.Map{"app_attr_id": "4"})
	fmt.Println(err)
	fmt.Println(d1.LastInsertId())
}
