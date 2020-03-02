package model

import "fmt"

// 声明一个Customer结构体，表示一个客户信息
type Customer struct {
	Id int
	Name string
	Gender string
	Age int
	Phone string
	Email string
}

// 使用工厂模式，返回Customer实例
func NewCustomer(id int, name string, gender string,
	age int, phone string, email string) *Customer {
		return &Customer{
			id,
			name,
			gender,
			age,
			phone, email,
		}
}
// 返回不带Id的Customer实例
func NewCustomer2(name string, gender string,
	age int, phone string, email string) *Customer {
	return &Customer{
		Name: name,
		Gender: gender,
		Age: age,
		Phone: phone,
		Email: email,
	}
}
// 返回格式化后的客户信息
func (this *Customer) GetInfo() string{
	info := fmt.Sprintf("%v\t%v\t%v\t%v\t%v\t%v", this.Id, this.Name, this.Gender,
		this.Age, this.Phone, this.Email)
	return info
}