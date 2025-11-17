package main

import "fmt"

// 电商支付方式
type Payment interface {
	Pay(amount float64) (string, error)
	Refund(transactionId string, amount float64) (string, error)
	Query(transactionId string) (string, error)
}
type Alipay struct {
	AppId      string
	AppSecret  string
	MerchantID string
}

func (a *Alipay) Pay(amount float64) (string, error) {
	return "支付宝支付成功", nil
}
func (a *Alipay) Refund(transactionId string, amount float64) (string, error) {
	return "支付宝退款成功", nil
}
func (a *Alipay) Query(transactionId string) (string, error) {
	return "支付宝查询成功", nil
}

type WechatPay struct {
	AppId      string
	AppSecret  string
	MerchantID string
}

func (w *WechatPay) Pay(amount float64) (string, error) {
	return "微信支付成功", nil
}
func (w *WechatPay) Refund(transactionId string, amount float64) (string, error) {
	return "微信退款成功", nil
}
func (w *WechatPay) Query(transactionId string) (string, error) {
	return "微信查询成功", nil
}

func ProcessPayment(payment Payment, amount float64) (string, error) {
	transactionId, err := payment.Pay(amount)
	if err != nil {
		return "", err
	}
	return transactionId, nil
}
func main() {
	alipay := &Alipay{
		AppId:      "20230824001",
		AppSecret:  "20230824001",
		MerchantID: "20230824001",
	}
	wechatPay := &WechatPay{
		AppId:      "20230824001",
		AppSecret:  "20230824001",
		MerchantID: "20230824001",
	}
	transactionId, err := ProcessPayment(alipay, 100.0)
	if err != nil {
		fmt.Println("支付宝支付失败:", err)
		return
	}
	fmt.Println("支付宝支付成功:", transactionId)
	transactionId, err = ProcessPayment(wechatPay, 100.0)
	if err != nil {
		fmt.Println("微信支付失败:", err)
		return
	}
	fmt.Println("微信支付成功:", transactionId)
	var anything interface{}
	anything = 100
	processEmptyInterface(anything)
	fmt.Println(anything)
}
func processEmptyInterface(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Println("int:", v)
	case string:
		fmt.Println("string:", v)
	default:
		fmt.Println("unknown type")
	}
}
