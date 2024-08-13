package email_template

import (
	"strings"
)

func GetResetPasswordEmailHTML(account string, verifyCode string) string {
	template := `
<meta http-equiv="Content-Type" content="text/html; charset=utf-8" />
<div class="content-div">
    <div style="border-style: solid; border-width: thin; border-color:#dadce0; border-radius: 8px; padding: 40px 20px;" align="center" class="mdv2rw">
        <img src="https://jsd.onmicrosoft.cn/gh/LanceHE6/PicGo@main/imgs/WMS-Logo.png" width="74" height="74" aria-hidden="true" style="margin-bottom: 16px;">
        <div style="">
            <div style="font-size: 24px; color: red">重置密码通知</div>
    </div>
    <div style="font-family: Roboto-Regular,Helvetica,Arial,sans-serif; font-size: 14px; color: rgba(0,0,0,0.87); line-height: 20px;padding-top: 20px; text-align: left;">简行云仓库 收到了为账号 <span style="font-weight: bold;">${account}</span> 的重置密码的请求。<br><br>请使用此验证码完成重置密码的操作：<br>
        <div style="text-align: center; font-size: 36px; margin-top: 20px; line-height: 44px;">${code}</div><br>此验证码将在 5 分钟后失效。<br><br>如果不是您本人操作，您的账号和邮箱可能已经泄露，请忽略这封电子邮件。
        <br><br><br>系统邮件 请勿回复
    </div>
</div>
</div>

<style>
    .content-div {
        position: relative;
        font-size: 14px;
        height: auto;
        padding: 15px 15px 10px 15px;
        z-index: 1;
        zoom: 1;
        line-height: 1.7;
        width: 550px;
        min-width: 500px;
        margin: 10px auto;
    }
</style>
`
	template = strings.Replace(template, "${account}", account, -1)
	template = strings.Replace(template, "${code}", verifyCode, -1)
	return template
}
