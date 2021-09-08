//提交
function submit() {
    close()

    //获取id
    var id = document.getElementById("id").value;

    //获取表单元素
    var name = document.getElementById("blogname").value;
    if (!name) {
        alertErrorMsg("友链名称为空");
        return;
    }
    var address = document.getElementById("blogaddress").value;
    if (!address) {
        alertErrorMsg("友链地址为空");
        return;
    }
    var img = document.getElementById("pictureaddress").value;
    if (!img) {
        alertErrorMsg("图片地址为空");
        return;
    }

    //构建json
    var form_data = {
        "id": id,
        "name": name,
        "address": address,
        "img": img
    }

    //新增
    $.ajax({
        url: '/link/add',
        type: 'POST',
        datatype: 'json',
        data: form_data,
        async: false,
        cache: false,
        success: function (res) {
            if (res.status === 1) {
                alertSuccessMsg();
                setTimeout(function () {
                    window.location.href = "/link/list/0";
                }, 1000);
            } else {
                alertErrorMsg(res.msg);
            }
        },
        error: function () {
            alertErrorMsg("发送新增请求失败");
        }
    });
}