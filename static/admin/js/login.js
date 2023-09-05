const app = {
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
        const that = this;
        $("#captchaImg").click(function () {
            that.getCaptcha();
        })
    }
};

$(function () {
    app.init()
})