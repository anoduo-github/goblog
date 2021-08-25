//初始化Markdown编辑器
var contentEditor;
$(function () {
    contentEditor = editormd("md-content", {
        width: "100%",
        height: 640,
        syncScrolling: "single",
        path: "../../static/lib/editormd/lib/",

        /**上传图片相关配置如下*/
        imageUpload: true,
        imageFormats: ["jpg", "jpeg", "gif", "png", "bmp", "webp"],
        imageUploadURL: "/blog/upload/picture",//注意你后端的上传图片服务地址
    });
});

$('#save-btn').click(function () {
    $('[name="published"]').val(false);
    $('#blog-form').submit();
});


$('#publish-btn').click(function () {
    $('[name="published"]').val(true);
    $('#blog-form').submit();
});

//提交
function submit() {
    close()
    //获取标签
    var tag = document.getElementById("tag").value;
    if (!tag) {
        alertErrorMsg("标签不能为空！")
        return
    }
    //获取标题
    var title = document.getElementById("title").value;
    if (!title) {
        alertErrorMsg("标题不能为空！")
        return
    }
    //获取内容
    var content = document.getElementById("content").value;
    if (!content) {
        alertErrorMsg("内容不能为空！")
        return
    }
    //获取分类
    var type = document.getElementById("type").value;
    if (!type) {
        alertErrorMsg("分类不能为空！")
        return
    }
    //获取首图地址
    var fp = document.getElementById("firstPicture").value;
    if (!fp) {
        alertErrorMsg("首图不能为空！")
        return
    }
    //获取描述
    var description = document.getElementById("description").value;
    if (!description) {
        alertErrorMsg("描述不能为空！")
        return
    }
    //获取id
    var id = document.getElementById("id").value;

    //创建post参数
    var form_data

    if (!id) {
        form_data = {
            "tag": tag,
            "title": title,
            "content": content,
            "type": type,
            "firstPicture": fp,
            "description": description
        }
    } else {
        form_data = {
            "tag": tag,
            "title": title,
            "content": content,
            "type": type,
            "firstPicture": fp,
            "description": description,
            "id": id
        }
    }



    console.log("form_data: ", form_data)

    $.ajax({
        url: "/blog/submit",
        type: 'POST',
        datatype: 'json',
        data: form_data,
        async: false,
        cache: false,
        success: function (res) {
            if (res.status === 1) {
                alertSuccessMsg();
                setTimeout(function () {
                    window.location.href = "/blog/list"
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