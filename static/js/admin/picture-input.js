//提交
function submit() {
    close()

    //获取id
    var id = document.getElementById("id").value;

    //获取表单元素
    var name = document.getElementById("picturename").value;
    if (!name) {
        alertErrorMsg("图片名称为空");
        return;
    }
    var time_place = document.getElementById("picturetime").value;
    if (!time_place) {
        alertErrorMsg("时间地点为空");
        return;
    }
    var address = document.getElementById("pictureaddress").value;
    if (!address) {
        alertErrorMsg("图片地址为空");
        return;
    }
    var description = document.getElementById("picturedescription").value;
    if (!description) {
        alertErrorMsg("图片描述为空");
        return;
    }

    //构建json
    var form_data = {
        "id": id,
        "name": name,
        "address": address,
        "time_place": time_place,
        "description": description
    }

    //新增
    $.ajax({
        url: '/picture/add',
        type: 'POST',
        datatype: 'json',
        data: form_data,
        async: false,
        cache: false,
        success: function (res) {
            if (res.status === 1) {
                alertSuccessMsg();
                setTimeout(function () {
                    window.location.href = "/picture/list/0";
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