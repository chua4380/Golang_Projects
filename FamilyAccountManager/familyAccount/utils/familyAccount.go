package utils

import "fmt"

type FamilyAccount struct {
	// 接收用户输入的选项
	key string
	// 控制是否退出for循环
	loop bool
	// 定义账户余额
	balance float64
	// 每次收支的金额
	money float64
	// 每次收支的说明
	note string
	// 收支的详情字符串，当有收支时，对details进行拼接
	details string
	// 定义一个变量，记录是否有收支情况
	flag bool
	// 用户名
	username string
	// 密码
	passwd string
}

// 编写要给工厂模式的构造方法， 返回一个*FamilyAccount
func NewFamilyAccount() *FamilyAccount {
	return &FamilyAccount{
		key : "",
		loop : true,
		balance : 10000.0,
		money : 0.0,
		note : "",
		flag : false,
		details : "收支\t账号余额\t收支金额\t说    明",
		username : "lioney",
		passwd : "qwerr12345",
	}
}

// 给该结构体绑定相应的方法
// 将显示明细写成一个方法
func (this *FamilyAccount) showDetails() {
	fmt.Println("---------------当前收支明细记录---------------")
	if this.flag {
		fmt.Println(this.details)
	} else {
		fmt.Println("当前没有收支明细，来一笔吧")
	}
}

// 登记收入
func (this *FamilyAccount) income() {
	fmt.Println("本次收入金额:")
	fmt.Scanln(&this.money)
	this.balance += this.money
	fmt.Println("本次收入说明:")
	fmt.Scanln(&this.note)
	// 将收入情况拼接到details变量
	// 收入		11000		1000		有人发红包
	this.details += fmt.Sprintf("\n收入\t%v\t%v\t%v", this.balance, this.money, this.note)
	this.flag = true
}

// 登记支出
func (this *FamilyAccount) pay() {
	fmt.Println("本次支出金额:")
	fmt.Scanln(&this.money)
	if this.money > this.balance {
		fmt.Println("余额不足，不能支出")
		//break
	} else {
		this.balance -= this.money
		fmt.Println("本次支出说明:")
		fmt.Scanln(&this.note)
		this.details += fmt.Sprintf("\n支出\t%v\t%v\t%v", this.balance, this.money, this.note)
		this.flag = true
	}
}

// 退出
func (this *FamilyAccount) exit() {
	fmt.Println("你确定要退出吗？ y/n")
	choice := ""
	for {
		fmt.Scanln(&choice)
		if choice == "y" || choice == "n"{
			break
		}
		fmt.Println("你的输入有误，请重新输入")
	}
	if choice == "y" {
		this.loop = false
	}
}

// 登陆验证
func (this *FamilyAccount) login() {
	fmt.Println("欢迎你使用家庭收支软件!")
	username := ""
	passwd := ""
	for {
		fmt.Println("请输入你的用户名:")
		fmt.Scanln(&username)
		fmt.Println("请输入你的密码:")
		fmt.Scanln(&passwd)
		if username != this.username || passwd != this.passwd {
			fmt.Println("你输入的用户名和密码错误，请重新输入....")
		} else {
			break
		}
	}
}

// 显示主菜单
func (this *FamilyAccount) MainMenu() {
	this.login()
	for {
		fmt.Println("\n---------------家庭收支记账软件---------------")
		fmt.Println("                 1.收支明细")
		fmt.Println("                 2.登记收入")
		fmt.Println("                 3.登记支出")
		fmt.Println("                 4.退出软件")
		fmt.Print("请选择(1-4):")
		fmt.Scanln(&this.key)
		switch this.key {
		case "1":
			this.showDetails()
		case "2":
			this.income()
		case "3":
			this.pay()
		case "4":
			this.exit()
		default:
			fmt.Println("请输入正确选项..")
		}
		if !this.loop {
			break
		}
	}
}