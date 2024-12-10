package main

import (
	"fmt"
	"html/template"
	"net/http"
	"regexp"
	"strconv"
)

func VerifyFormHander(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	fmt.Println("表单内容:", r.Form)
	fmt.Println("method:", r.Method)

	if r.Method == "GET" {
		// 显示静态页面
		t, _ := template.ParseFiles("./form.tpl")
		t.Execute(w, nil) // 解析模板
	} else {
		// 校验表单

		// 必填字段
		if len(r.Form["username"][0]) == 0 {
			fmt.Fprintf(w, "用户名不能为空")
			return
		}

		// 数字
		if m, _ := regexp.MatchString("^[0-9]+$", r.Form.Get("age")); !m {
			fmt.Fprintf(w, "年龄必须是正数")
			return
		}

		// 中文校验
		if m, _ := regexp.MatchString("^[a-zA-Z]+$", r.Form.Get("engname")); !m {
			fmt.Fprintf(w, "英文名必须是英文")
			return
		}

		// 电子邮件地址
		if m, _ := regexp.MatchString(`^([\w\.\_]{2,10})@(\w{1,}).([a-z]{2,4})$`, r.Form.Get("email")); !m {
			fmt.Fprintf(w, "邮箱无效")
			return
		}
		// 手机号码
		if m, _ := regexp.MatchString(`^(1[3|4|5|8][0-9]\d{4,8})$`, r.Form.Get("mobile")); !m {
			fmt.Fprintf(w, "手机号无效")
			return
		}

		// 下拉菜单
		slice := []string{"apple", "pear", "banane"}
		hasfruit := false
		for _, v := range slice {
			if v == r.Form.Get("fruit") {
				hasfruit = true
			}
		}
		if !hasfruit {
			fmt.Fprintf(w, "下拉菜单选项不存在")
			return
		}

		// 单选按钮
		slice2 := []int{1, 2}
		hasgender := false
		for _, v2 := range slice2 {
			getint, _ := strconv.Atoi(r.Form.Get("gender"))
			if v2 == getint {
				hasgender = true
			}
		}
		if !hasgender {
			fmt.Fprintf(w, "性别选项有误")
			return
		}

		// 复选框
		slice3 := []string{"football", "basketball", "tennis"}
		hasinterest := false
		for _, v3 := range slice3 {
			if v3 == r.Form.Get("interest") {
				hasinterest = true
			}
		}
		if !hasinterest {
			fmt.Fprintf(w, "爱好选项有误")
			return
		}

		// 身份证号码
		isusercard := false
		usercard := r.Form.Get("usercard")

		// 验证15位身份证，15位的是全部数字
		m1, _ := regexp.MatchString(`^(\d{15})$`, usercard)

		// 验证18位身份证，18位前17位为数字，最后一位是校验位，可能为数字或字符X。
		m2, _ := regexp.MatchString(`^(\d{17})([0-9]|X)$`, usercard)

		if m1 || m2 {
			isusercard = true
		}

		if !isusercard {
			fmt.Fprintf(w, "身份证号有误")
			return
		}

		fmt.Fprintf(w, "表单验证通过")

	}
}

func main() {
	http.HandleFunc("/verifyForm", VerifyFormHander)
	fmt.Println("服务器已经启动，请在浏览器地址栏中输入 http://localhost:8900/verifyForm")

	err := http.ListenAndServe(":8900", nil)
	if err != nil {
		fmt.Println("服务器启动失败")
	}
}
