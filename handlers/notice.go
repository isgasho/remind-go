package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	sms "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/sms/v20190711"
	"gopkg.in/gomail.v2"
	"log"
	"remind-go/config"
	"remind-go/models"
	"time"
)

type Send interface {
	SendNotice(todo models.Todo)
}

type Email struct {
	Title string
	Body  string
	Email string
}
type Phone struct {
	Phone string
}

var timer *time.Timer

func SendEmailOrPhone(todo models.Todo, email *Email, phone *Phone) {
	//查看这条通知离现在的时间 定时器安排一下发送
	now := GetLocalTimeNow()
	noticeTime, _ := time.ParseInLocation("2006-01-02 15:04:05",
		todo.NoticeTime.Format("2006-01-02 15:04:05"), time.Local)
	diff := noticeTime.Sub(now)
	timer = time.NewTimer(diff)
	<-timer.C
	//暂时没接
	//email.SendNotice()
	phone.SendNotice(todo.Id)
	log.Println("此次发送的手机号是:%d", phone)
	log.Println("主键id是:", todo.Id)
}

func (email *Email) SendNotice() {
	mConfig := config.LoadConfig()
	sendMail := gomail.NewMessage()
	sendMail.SetHeader(`From`, mConfig.Email.User)
	sendMail.SetHeader(`To`, "wuqinqiang050@gmail.com")
	sendMail.SetHeader(`Subject`, "来自吴亲库里的温馨提醒")
	sendMail.SetBody(`text/html`, email.Body)
	err := gomail.NewDialer(
		mConfig.Email.Host, mConfig.Email.Port, mConfig.Email.User,
		mConfig.Email.Pass).DialAndSend(sendMail)
	if err != nil {
		log.Println(err.Error())
		return
	}
}

func (phone *Phone) SendNotice(id int64) {
	config := config.LoadConfig()
	credential := common.NewCredential(
		config.Teng.SECRETID,
		config.Teng.SecretKey,
	)
	cpf := profile.NewClientProfile()
	/* SDK 默认使用 POST 方法
	 * 如需使用 GET 方法，可以在此处设置，但 GET 方法无法处理较大的请求 */
	cpf.HttpProfile.ReqMethod = "POST"
	cpf.SignMethod = "HmacSHA1"
	/* 实例化 SMS 的 client 对象
	 * 第二个参数是地域信息，可以直接填写字符串 ap-guangzhou，或者引用预设的常量 */
	client, _ := sms.NewClient(credential, "ap-guangzhou", cpf)

	/* 实例化一个请求对象，根据调用的接口和实际情况，可以进一步设置请求参数
	   * 您可以直接查询 SDK 源码确定接口有哪些属性可以设置
	    * 属性可能是基本类型，也可能引用了另一个数据结构
	    * 推荐使用 IDE 进行开发，可以方便地跳转查阅各个接口和数据结构的文档说明 */
	request := sms.NewSendSmsRequest()
	/* 基本类型的设置:
	 * SDK 采用的是指针风格指定参数，即使对于基本类型也需要用指针来对参数赋值。
	 * SDK 提供对基本类型的指针引用封装函数
	 * 帮助链接：
	 * 短信控制台：https://console.cloud.tencent.com/smsv2
	 * sms helper：https://cloud.tencent.com/document/product/382/3773 */

	/* 短信应用 ID: 在 [短信控制台] 添加应用后生成的实际 SDKAppID，例如1400006666 */
	request.SmsSdkAppid = common.StringPtr(config.Teng.SDKAppID)
	/* 短信签名内容: 使用 UTF-8 编码，必须填写已审核通过的签名，可登录 [短信控制台] 查看签名信息 */
	request.Sign = common.StringPtr("库里的深夜食堂公众号")
	/* 模板 ID: 必须填写已审核通过的模板 ID，可登录 [短信控制台] 查看模板 ID */
	request.TemplateID = common.StringPtr(config.Teng.TemplateID)
	request.PhoneNumberSet = common.StringPtrs([]string{"+86" + phone.Phone})

	// 通过 client 对象调用想要访问的接口，需要传入请求对象
	response, err := client.SendSms(request)
	// 处理异常
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return
	}
	// 非 SDK 异常，直接失败。实际代码中可以加入其他的处理
	if err != nil {
		panic(err)
	}
	b, _ := json.Marshal(response.Response)
	// 打印返回的 JSON 字符串
	fmt.Printf("%s", b)
	//标识成功
	models.SetSuccessStatus(id)

}
