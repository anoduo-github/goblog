//获取信息对象
var successDom = document.getElementById("success");
var errorDom = document.getElementById("error");
var errorMsg = document.getElementById("errorMsg");
var infoDom = document.getElementById("info");

//正则表达式
var exp = /^[a-zA-Z0-9]{4,10}$/;

//提示错误信息
function alertErrorMsg(msg) {
    successDom.style.display = "none";
    errorMsg.innerHTML = msg;
    errorDom.style.display = "block";
}

//提示成功信息
function alertSuccessMsg() {
    errorDom.style.display = "none";
    successDom.style.display = "block";
}

//关闭信息栏
function close() {
    successDom.style.display = "none";
    errorDom.style.display = "none";
    infoDom.style.display = "none";
}

//注册
function register() {
    close();

    //获取表单元素并检查格式
    var username = document.getElementById("username").value;
    if (!username) {
        alertErrorMsg("用户名不能为空");
        return;
    } else if (!exp.test(username)) {
        alertErrorMsg("用户名格式错误");
        return;
    }
    var password = document.getElementById("password").value;
    if (!password) {
        alertErrorMsg("密码不能为空");
        return;
    } else if (!exp.test(password)) {
        alertErrorMsg("密码格式错误");
        return;
    }
    var password2 = document.getElementById("password2").value;
    if (!password2) {
        alertErrorMsg("请输入确认密码");
        return;
    } else if (!exp.test(password2)) {
        alertErrorMsg("密码格式错误");
        return;
    }
    var code = document.getElementById("code").value;
    if (!code) {
        alertErrorMsg("验证码不能为空");
        return;
    }

    var myForm = {
        "username": username,
        "password": password,
        "password2": password2,
        "code": code
    };

    $.ajax({
        url: "/doregister",
        type: 'POST',
        datatype: 'json',
        data: myForm,
        async: false,
        cache: false,
        success: function (res) {
            if (res.status === 1) {
                alertSuccessMsg();
                setTimeout(function () {
                    window.location.href = "/login"
                }, 1000);
            } else {
                alertErrorMsg(res.msg);
            }
        },
        error: function () {
            alertErrorMsg("发送注册请求失败");
        }
    });
}

//登录
function login() {
    close();

    //获取表单元素并检查格式
    var username = document.getElementById("username").value;
    if (!username) {
        alertErrorMsg("用户名不能为空");
        return;
    } else if (!exp.test(username)) {
        alertErrorMsg("用户名格式错误");
        return;
    }
    var password = document.getElementById("password").value;
    if (!password) {
        alertErrorMsg("密码不能为空");
        return;
    } else if (!exp.test(password)) {
        alertErrorMsg("密码格式错误");
        return;
    }
    var code = document.getElementById("code").value;
    if (!code) {
        alertErrorMsg("验证码不能为空");
        return;
    }

    var myForm = {
        "username": username,
        "password": password,
        "code": code
    };

    $.ajax({
        url: '/dologin',
        type: 'POST',
        datatype: 'json',
        data: myForm,
        async: false,
        cache: false,
        success: function (res) {
            if (res.status === 1) {
                alertSuccessMsg();
                setTimeout(function () {
                    window.location.href = "/"
                }, 1000);
            } else {
                alertErrorMsg(res.msg);
            }
        },
        error: function () {
            alertErrorMsg("发送注册请求失败");
        }
    });
}
