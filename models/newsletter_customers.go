package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/beego/beego/v2/client/orm"
)

type Newsletter_customers struct {
	CustomerId   int64     `orm:"auto"`
	FirstName    string    `orm:"size(100)"`
	LastName     string    `orm:"size(100)"`
	Email        string    `orm:"size(100)"`
	DateCreated  time.Time `orm:"type(datetime)"`
	DateModified time.Time `orm:"type(datetime)"`
	CreatedBy    int
	ModifiedBy   int
	Active       int
}

func init() {
	orm.RegisterModel(new(Newsletter_customers))
}

// AddNewsletter_customers insert a new Newsletter_customers into database and returns
// last inserted Id on success.
func AddNewsletter_customers(m *Newsletter_customers) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetNewsletter_customersById retrieves Newsletter_customers by Id. Returns error if
// Id doesn't exist
func GetNewsletter_customersById(id int64) (v *Newsletter_customers, err error) {
	o := orm.NewOrm()
	v = &Newsletter_customers{CustomerId: id}
	if err = o.QueryTable(new(Newsletter_customers)).Filter("CustomerId", id).RelatedSel().One(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllNewsletter_customers retrieves all Newsletter_customers matches certain condition. Returns empty list if
// no records exist
func GetAllNewsletter_customers(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Newsletter_customers))
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

	var l []Newsletter_customers
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

// UpdateNewsletter_customers updates Newsletter_customers by Id and returns error if
// the record to be updated doesn't exist
func UpdateNewsletter_customersById(m *Newsletter_customers) (err error) {
	o := orm.NewOrm()
	v := Newsletter_customers{CustomerId: m.CustomerId}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteNewsletter_customers deletes Newsletter_customers by Id and returns error if
// the record to be deleted doesn't exist
func DeleteNewsletter_customers(id int64) (err error) {
	o := orm.NewOrm()
	v := Newsletter_customers{CustomerId: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Newsletter_customers{CustomerId: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
