package service

import "go_code/project02/customerManager/model"

// 完成对Customer的操作

type CustomerService struct {
	cusomers []model.Customer
	// 表示当前切片有多少个客户
	// 该字段后面还可以作为新客户的编号id+1
	customNum int
}


func NewCustomerService() *CustomerService {
	// 定义切片
	customerService := &CustomerService{}
	// 初始化一个客户
	customer := model.NewCustomer2("小马", "男", 20,
		"135", "xm@sina.com")
	customerService.Add(customer)
	return customerService
}

// 返回客户切片
func (this *CustomerService) List() []model.Customer {
	return this.cusomers
}

// 添加客户到customer切片
func (this *CustomerService) Add(cusomer *model.Customer) bool {
	this.customNum++
	cusomer.Id = this.customNum
	this.cusomers = append(this.cusomers, *cusomer)
	return true
}

// 根据id查找客户在切片中对应的下标，找不到返回-1
func (this *CustomerService) FindById(id int) int {
	index := -1
	for i :=0; i <len(this.cusomers); i++ {
		if this.cusomers[i].Id == id {
			index = i
			break
		}
	}
	return index
}

// 根据id修改客户信息
func (this *CustomerService) Update(customer *model.Customer) bool {
	index := this.FindById(customer.Id)
	if index == -1 {
		return false
	}
	// 修改
	if customer.Name != "" {
		this.cusomers[index].Name = customer.Name
	}
	if customer.Age != 0 {
		this.cusomers[index].Age = customer.Age
	}
	if customer.Phone != "" {
		this.cusomers[index].Phone = customer.Phone
	}
	if customer.Email != "" {
		this.cusomers[index].Email = customer.Email
	}
	return true
}

// 根据id删除客户（从切片中删除）
func (this *CustomerService) Delete(id int) bool {
	index := this.FindById(id)
	if index == -1 {
		return false
	}
	// 从切片中删除一个元素
	this.cusomers = append(this.cusomers[:index], this.cusomers[index+1:]...)
	return true
}
