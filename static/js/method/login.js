//login 登录
function login() {
    var $ = layui.jquery;
    //获取页面信息
    var username = document.getElementById("username").value;
    var password = document.getElementById("password").value;
    if (!username) {
        layer.alert("用户名为空");
        return;
    }
    if (!password) {
        layer.alert("密码为空");
        return;
    }
    $.ajax({
        url: "/login/login?username="+username+"&password="+password,
        type: "get",
        success: function(res) {
            if (res.status === 0) {
                setTimeout(function(){
                    window.onload = "/";
                }, 500);
            }
        }
    })
}