//搜索
function search() {
    var name = document.getElementById("name").value;
    var url = "/user/list?index=0&name=" + name

    window.location.href = url
}