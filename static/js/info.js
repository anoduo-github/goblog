//获取信息对象
var successDom = document.getElementById("success");
var errorDom = document.getElementById("error");
var errorMsg = document.getElementById("errorMsg");

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