package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/beego/beego/v2/client/orm"
)

type Customer_guarantors struct {
	CustomerGuarantorId int64      `orm:"auto"`
	Name                string     `orm:"size(120)"`
	Contact             string     `orm:"size(50)"`
	Customer            *Customers `orm:"rel(fk)"`
	DateCreated         time.Time  `orm:"type(datetime)"`
	DateModified        time.Time  `orm:"type(datetime)"`
	CreatedBy           int
	ModifiedBy          int
}

func init() {
	orm.RegisterModel(new(Customer_guarantors))
}

// AddCustomer_guarantors insert a new Customer_guarantors into database and returns
// last inserted Id on success.
func AddCustomer_guarantors(m *Customer_guarantors) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetCustomer_guarantorsById retrieves Customer_guarantors by Id. Returns error if
// Id doesn't exist
func GetCustomer_guarantorsById(id int64) (v *Customer_guarantors, err error) {
	o := orm.NewOrm()
	v = &Customer_guarantors{CustomerGuarantorId: id}
	if err = o.QueryTable(new(Customer_guarantors)).Filter("CustomerGuarantorId", id).RelatedSel().One(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllCustomer_guarantors retrieves all Customer_guarantors matches certain condition. Returns empty list if
// no records exist
func GetAllCustomer_guarantors(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Customer_guarantors))
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		qs = qs.Filter(k, v)
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

	var l []Customer_guarantors
	qs = qs.OrderBy(sortFields...).RelatedSel()
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

// UpdateCustomer_guarantors updates Customer_guarantors by Id and returns error if
// the record to be updated doesn't exist
func UpdateCustomer_guarantorsById(m *Customer_guarantors) (err error) {
	o := orm.NewOrm()
	v := Customer_guarantors{CustomerGuarantorId: m.CustomerGuarantorId}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteCustomer_guarantors deletes Customer_guarantors by Id and returns error if
// the record to be deleted doesn't exist
func DeleteCustomer_guarantors(id int64) (err error) {
	o := orm.NewOrm()
	v := Customer_guarantors{CustomerGuarantorId: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Customer_guarantors{CustomerGuarantorId: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
