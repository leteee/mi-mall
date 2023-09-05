var app = {
    init: function () {
        this.getCaptcha();
        this.captchaClick();
    },
    getCaptcha: function () {
        //h获取图形验证码
        $.get("/admin/captcha?t=" + Math.random()).then(function (result) {
            $("#captchaId").attr("value", result.captchaId);
            $("#captchaImg").attr("src", result.captchaImage);
        })
    },
    captchaClick: function () {
        $("#captchaImg").click(function () {
            this.getCaptcha();
        }).bind(this)
    }
}

$(function () {
    app.init()
})