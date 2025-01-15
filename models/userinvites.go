package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/beego/beego/v2/client/orm"
)

type UserInvites struct {
	UserInviteId    int64       `orm:"auto"`
	InvitedBy       *Users      `orm:"rel(fk);column(invited_by)"`
	InvitationToken *UserTokens `orm:"rel(fk);column(invitation_token)"`
	Email           string      `orm:"column(email)"`
	Role            *Roles      `orm:"rel(fk);column(role)"`
	Status          string
	DateCreated     time.Time `orm:"type(datetime)"`
	DateModified    time.Time `orm:"type(datetime)"`
	CreatedBy       int
	ModifiedBy      int
	Active          int
}

func init() {
	orm.RegisterModel(new(UserInvites))
}

// AddUserInvites insert a new UserInvites into database and returns
// last inserted Id on success.
func AddUserInvites(m *UserInvites) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetUserInvitesById retrieves UserInvites by Id. Returns error if
// Id doesn't exist
func GetUserInvitesById(id int64) (v *UserInvites, err error) {
	o := orm.NewOrm()
	v = &UserInvites{UserInviteId: id}
	if err = o.QueryTable(new(UserInvites)).Filter("UserInviteId", id).RelatedSel().One(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetUserInvitesByTokenId retrieves UserInvites by Id. Returns error if
// Id doesn't exist
func GetUserInvitesByToken(token *UserTokens) (v *UserInvites, err error) {
	o := orm.NewOrm()
	v = &UserInvites{InvitationToken: token}
	if err = o.QueryTable(new(UserInvites)).Filter("InvitationToken", token).RelatedSel().One(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetUserInvitesByTokenId retrieves UserInvites by Id. Returns error if
// Id doesn't exist
func GetUserInvitesByEmail(email string) (v *UserInvites, err error) {
	o := orm.NewOrm()
	v = &UserInvites{Email: email}
	if err = o.QueryTable(new(UserInvites)).Filter("Email", email).RelatedSel().One(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllUserInvites retrieves all UserInvites matches certain condition. Returns empty list if
// no records exist
func GetAllUserInvites(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(UserInvites))
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

	var l []UserInvites
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

// UpdateUserInvites updates UserInvites by Id and returns error if
// the record to be updated doesn't exist
func UpdateUserInvitesById(m *UserInvites) (err error) {
	o := orm.NewOrm()
	v := UserInvites{UserInviteId: m.UserInviteId}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteUserInvites deletes UserInvites by Id and returns error if
// the record to be deleted doesn't exist
func DeleteUserInvites(id int64) (err error) {
	o := orm.NewOrm()
	v := UserInvites{UserInviteId: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&UserInvites{UserInviteId: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
