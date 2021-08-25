//删除
function del(id) {
    close()
    var flag = confirm('确定要删除该文章吗？')
    if (flag) {
        console.log("delete")
        //删除
        $.ajax({
            url: '/blog/delete/' + id,
            type: 'GET',
            async: false,
            cache: false,
            success: function (res) {
                if (res.status === 1) {
                    alertSuccessMsg();
                    setTimeout(function () {
                        window.location.reload();
                    }, 1500);
                } else {
                    alertErrorMsg(res.msg);
                }
            },
            error: function () {
                alertErrorMsg("发送删除请求失败");
            }
        });
    }
}

//新增
function add() {
    window.location.href = "/blog/edit"
}

//搜索
function search() {
    close()
    //获取标题
    var title = document.getElementById("title").value;

    //获取分类
    var type = document.getElementById("type").value;

    var url = "/blog/list?title=" + title + "&type=" + type

    window.location.href = url
}

//清空搜索框
function clear() {
    document.getElementById("title").value = "";
    document.getElementById("type").value = "";
}