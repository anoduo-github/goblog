//获取用户登录信息
$.ajax({
    url: "/user/info",
    type: 'GET',
    async: false,
    cache: false,
    success: function (res) {
        if (res.status === 1) {
            if (res.msg != "admin") {
                document.getElementById("yonghu_img").src = "../../../static/images/avatar2.png"
            } else {
                document.getElementById("yonghu_admin").style.display = "block"
            }
            document.getElementById("yonghu_name").innerHTML = res.msg
            document.getElementById("yonghu_href").innerHTML = "注销"
            document.getElementById("yonghu_href").href = "/logout"
        }
    },
    error: function () {
        alertErrorMsg("获取用户登录信息的请求，发送失败");
    }
});



//搜索
function search() {
    var query = document.getElementById("query").value
    window.location.href = "/blog/list/search?index=0&title=" + query
}