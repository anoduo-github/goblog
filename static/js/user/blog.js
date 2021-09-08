//markdown编辑器展示
$(function () {
    var testView = editormd.markdownToHTML("test-markdown-view", {
        //markdown : "[TOC]\n### Hello world!\n## Heading 2", // Also, you can dynamic set Markdown text
        //htmlDecode : true,  // Enable / disable HTML tag encode.
        //htmlDecode : "style,script,iframe",  // Note: If enabled, you should filter some dangerous HTML tags for website security.
    });
});

//赞赏
function appreciate() {
    var userId = document.getElementById("userId").value
    var blogId = document.getElementById("blogId").value

    if (!userId || userId == 0) {
        alertErrorMsg("未获取到当前用户信息,请先<a href='/login'>登录</a>")
        return
    }

    if (!blogId) {
        alertErrorMsg("未获取到当前博客信息")
        return
    }

    var form_data = {
        "userId": userId,
        "blogId": blogId
    }

    $.ajax({
        url: "/like/add",
        type: 'POST',
        datatype: 'json',
        data: form_data,
        async: false,
        cache: false,
        success: function (res) {
            if (res.status === 1) {
                alertSuccessMsg();
                setTimeout(function () {
                    window.location.href = "/blog/details/" + blogId
                }, 1000);
            } else {
                alertErrorMsg(res.msg);
            }
        },
        error: function () {
            alertErrorMsg("发送点赞请求失败");
        }
    });
}