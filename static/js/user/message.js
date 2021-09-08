$('.wechat').popup({
    popup: $('.wechat-qr'),
    position: 'bottom center'
});

$('.qq').popup();

//回复评论
function reply(obj) {
    //获取父评论id
    var messageId = $(obj).data('messageid');
    //获取父评论用户名
    var messageNickname = $(obj).data('messagenickname');
    //将父评论用户名显示在文本输入框中
    $("[name='content']").attr("placeholder", "@" + messageNickname).focus();
    //将父评论id赋值到隐藏的父id输入框中
    $("[name='parent_id']").val(messageId);

    //将页面滚动到指定位置
    var blog_id = document.getElementById("blog_id").value
    if (blog_id != -1) {
        //跳到博客评论区底部
        window.scrollTo(0, 1600)
    } else {
        //跳到留言板发布区上面
        window.scrollTo(0, 0)
    }
}

//提交评论
function submit() {
    close()

    //获取参数
    var parent_id = document.getElementById("parent_id").value
    var blog_id = document.getElementById("blog_id").value
    var content = document.getElementById("content").value

    if (!content) {
        alertErrorMsg("评论不能为空！")
        return
    }

    //构建参数
    var form_data = {
        "parent_id": parent_id,
        "blog_id": blog_id,
        "content": content
    }

    $.ajax({
        url: "/comment/add",
        type: 'POST',
        datatype: 'json',
        data: form_data,
        async: false,
        cache: false,
        success: function (res) {
            if (res.status === 1) {
                alertSuccessMsg();
                setTimeout(function () {
                    if (blog_id != -1) {
                        window.location.href = "/blog/details/" + blog_id
                    } else {
                        window.location.href = "/comment/0"
                    }
                }, 1000);
            } else {
                alertErrorMsg(res.msg);
            }
        },
        error: function () {
            alertErrorMsg("发送提交评论请求失败");
        }
    });
}