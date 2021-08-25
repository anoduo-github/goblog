//获取信息对象
var successDom = document.getElementById("success");
var errorDom = document.getElementById("error");
var errorMsg = document.getElementById("errorMsg");
var editMsg = document.getElementById("edit");

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
}

//提交
function submit() {
    close()

    //获取表单元素
    var name = document.getElementById("name").value;
    if (!name) {
        alertErrorMsg("分类名称为空");
        return
    }
    //获取id
    var id = editMsg.value;
    console.log("id:", id)

    if (id == "-1") {
        console.log("add")
        //新增
        $.ajax({
            url: '/type/add/' + name,
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
                alertErrorMsg("发送新增请求失败");
            }
        });
    } else {
        console.log("edit")
        //修改
        $.ajax({
            url: '/type/edit/' + id + "/" + name,
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
                alertErrorMsg("发送修改请求失败");
            }
        });
    }
}