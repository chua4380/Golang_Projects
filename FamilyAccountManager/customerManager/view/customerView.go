package main

import (
	"fmt"
	"go_code/project02/customerManager/service"
	"go_code/project02/customerManager/model"
)

type customerView struct {
	key string  // 接收用户输入
	loop bool   // 是否循环显示主菜单
	customerService *service.CustomerService
}

// 显示所有的客户信息
func (this *customerView) List() {
	// 获得客户信息的切片
	customers := this.customerService.List()
	// 显示
	fmt.Println("--------------------客户列表--------------------")
	fmt.Println("编号\t姓名\t性别\t年龄\t电话\t邮箱")
	for i:=0; i<len(customers); i++ {
		fmt.Println(customers[i].GetInfo())
	}
	fmt.Println("-----------------------------------------------")
}

// 得到用户输入，构建新的客户，并完成添加
func (this *customerView) add() {
	fmt.Println("----------添加客户----------")
	fmt.Println("姓名:")
	name := ""
	fmt.Scanln(&name)
	fmt.Println("性别:")
	gender := ""
	fmt.Scanln(&gender)
	fmt.Println("年龄:")
	age := 0
	fmt.Scanln(&age)
	fmt.Println("电话:")
	phone := ""
	fmt.Scanln(&phone)
	fmt.Println("邮箱")
	email := ""
	fmt.Scanln(&email)

	// 构建新的customer实例，customerservice分配编号
	customer := model.NewCustomer2(name, gender, age, phone, email)

	if this.customerService.Add(customer) {
		fmt.Println("----------添加成功!----------")
	} else {
		fmt.Println("----------添加失败!----------")
	}
}

// 得到用户输入的id,修改id对应的客户的信息(只允许修改客户姓名，年龄，电话，邮箱)
func (this *customerView) update() {
	fmt.Println("----------修改客户信息----------")
	fmt.Println("请选择待修改的客户编号(-1退出)")
	id := -1
	for {
		fmt.Scanln(&id)
		if id == -1 {
			return // 放弃删操作
		}
		if this.customerService.FindById(id) == -1 {
			fmt.Println("输入的id号不存在!请重新输入...")
		} else {
			break
		}
	}
	name := ""
	fmt.Println("姓名:")
	fmt.Scanln(&name)
	age := 0
	fmt.Println("年龄:")
	fmt.Scanln(&age)
	phone := ""
	fmt.Println("电话:")
	fmt.Scanln(&phone)
	email := ""
	fmt.Println("邮箱:")
	fmt.Scanln(&email)

	temp := model.NewCustomer(id, name, "", age, phone, email)
	if this.customerService.Update(temp) {
		fmt.Println("----------修改成功!----------")
	} else {
		fmt.Println("----------修改失败!----------")
	}
}

// 得到用户输入的id，删除该id对应的客户
func (this *customerView) delete() {
	fmt.Println("----------删除客户----------")
	fmt.Println("请选择待删除的客户编号(-1退出)")
	id := -1
	fmt.Scanln(&id)
	if id == -1 {
		return // 放弃删操作
	}
	fmt.Println("确认是否删除？y/n")
	choice := ""
	for {
		fmt.Scanln(&choice)
		if(choice == "y" || choice == "n") {
			break
		} else {
			fmt.Println("输入有误!请重新输入...")
		}
	}
	if(choice == "n") {
		return  // 放弃本次删除操作
	}
	// 调用customerService的Delete方法执行删除操作
	if this.customerService.Delete(id) {
		fmt.Println("----------删除成功!----------")
	} else {
		fmt.Println("----------删除失败,输入的id号不存在!!!----------")
	}
}

// 退出系统
func (this *customerView) exit() {
	fmt.Println("确认是否退出？y/n")
	choice := ""
	for {
		fmt.Scanln(&choice)
		if(choice == "y" || choice == "n") {
			break
		} else {
			fmt.Println("输入有误!请重新输入...")
		}
	}
	if choice == "n" {
		return // 不退出
	} else {
		this.loop = false
	}
}

// 显示主菜单
func (this *customerView) mainMenu() {
	for {
		fmt.Println("\n----------客户信息管理软件----------")
		fmt.Println("            1.添加客户")
		fmt.Println("            2.修改客户")
		fmt.Println("            3.删除客户")
		fmt.Println("            4.客户列表")
		fmt.Println("            5.退   出")
		fmt.Print("请选择(1-5):")
		fmt.Scanln(&this.key)
		switch this.key {
		case "1":
			this.add()
		case "2":
			this.update()
		case "3":
			this.delete()
		case "4":
			this.List()
		case "5":
			this.exit()
		default:
			fmt.Println("你的输入有误，请重新输入...")
		}
		if !this.loop {
			fmt.Println("已退出客户管理系统")
			break
		}
	}
}

func main() {
	// 创建customerView实例
	customerView := customerView{
		key:"",
		loop: true,
	}
	customerView.customerService = service.NewCustomerService()
	// 显示主菜单
	customerView.mainMenu()
}
