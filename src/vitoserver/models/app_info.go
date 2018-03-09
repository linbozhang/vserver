package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type AppInfo struct {
	Id          int       `orm:"column(id);pk;auto"`
	AppName     string    `orm:"column(app_name);size(100)"`
	VersionId   string    `orm:"column(version_id);size(20)"`
	VersionCode int16     `orm:"column(version_code)"`
	Description string    `orm:"column(description);size(1024);null"`
	IconPath    string    `orm:"column(icon_path);size(512);null"`
	UserId      int       `orm:"column(user_id);null"`
	UpdateDate  time.Time `orm:"column(update_date);type(datetime);null"`
	CreateDate  time.Time `orm:"column(create_date);type(datetime);null"`
}

func (t *AppInfo) TableName() string {
	return "app_info"
}

func init() {
	orm.RegisterModel(new(AppInfo))
}

// AddAppInfo insert a new AppInfo into database and returns
// last inserted Id on success.
func AddAppInfo(m *AppInfo) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetAppInfoById retrieves AppInfo by Id. Returns error if
// Id doesn't exist
func GetAppInfoById(id int) (v *AppInfo, err error) {
	o := orm.NewOrm()
	v = &AppInfo{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}
func GetAppListByDevId(devid int)(ml []AppInfo,err error){
	o:=orm.NewOrm()
	var list [] AppInfo
	_,err1:=o.QueryTable("app_info").Filter("user_id",devid).All(&list)
	if err1==nil{
		return list,nil
	}
	err=err1
	return nil,err

}
// GetAllAppInfo retrieves all AppInfo matches certain condition. Returns empty list if
// no records exist
func GetAllAppInfo(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(AppInfo))
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		if strings.Contains(k, "isnull") {
			qs = qs.Filter(k, (v == "true" || v == "1"))
		} else {
			qs = qs.Filter(k, v)
		}
	}
	// order by:
	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			// 1) for each sort field, there is an associated order
			for i, v := range sortby {
				orderby := ""
				if order[i] == "desc" {
					orderby = "-" + v
				} else if order[i] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
			qs = qs.OrderBy(sortFields...)
		} else if len(sortby) != len(order) && len(order) == 1 {
			// 2) there is exactly one order, all the sorted fields will be sorted by this order
			for _, v := range sortby {
				orderby := ""
				if order[0] == "desc" {
					orderby = "-" + v
				} else if order[0] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return nil, errors.New("Error: unused 'order' fields")
		}
	}

	var l []AppInfo
	qs = qs.OrderBy(sortFields...)
	if _, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				ml = append(ml, v)
			}
		} else {
			// trim unused fields
			for _, v := range l {
				m := make(map[string]interface{})
				val := reflect.ValueOf(v)
				for _, fname := range fields {
					m[fname] = val.FieldByName(fname).Interface()
				}
				ml = append(ml, m)
			}
		}
		return ml, nil
	}
	return nil, err
}

// UpdateAppInfo updates AppInfo by Id and returns error if
// the record to be updated doesn't exist
func UpdateAppInfoById(m *AppInfo) (err error) {
	o := orm.NewOrm()
	v := AppInfo{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteAppInfo deletes AppInfo by Id and returns error if
// the record to be deleted doesn't exist
func DeleteAppInfo(id int) (err error) {
	o := orm.NewOrm()
	v := AppInfo{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&AppInfo{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
