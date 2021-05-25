//注意：导航 依赖 element 模块，否则无法进行功能性操作
layui.use(['form','element','layer','jquery'], function(){
    
});

//markdown编辑器
$(function() {
    var editor = editormd("editor", {
        width: "100%",
        height: "100%",
        path : "../static/editor.md-master/lib/" , // Autoload modules mode, codemirror, marked... dependents libs path
        saveHTMLToTextarea : true,

        /**上传图片相关配置如下*/
        imageUpload : true,
        imageFormats : ["jpg", "jpeg", "gif", "png", "bmp", "webp"],
        imageUploadURL : "/blog/upload/picture",//注意你后端的上传图片服务地址
    });
});

//submit 提交
function submit() {
    layui.use(['form','element','layer','jquery'], function(){
      var $ = layui.jquery;
      var layer = layui.layer;

      //获取参数
      var title=document.getElementById("title").value;
      if (title.length===0) {
        layer.alert("标题不能为空！");
        return;
      }
      var kind=document.getElementById("kind").value;
      if (title.length===0) {
        layer.alert("标题不能为空！");
        return;
      }
      var text=document.getElementById("text").value;
      if (title.length===0) {
        layer.alert("内容不能为空！");
        return;
      }
      alert(text);
      //创建js对象
      var jsonObject={
        "title":title,
        "kind":kind,
        "text":text
      }
      var json_data=JSON.stringify(jsonObject);

      $.ajax({
        url: '/blog/submit',
        type: 'POST',
        datatype: 'json',
        data: json_data,
        async: false,
        cache: false,
        success: function(res){
            if (res.status === 1) {
                layer.alert(res.msg);
                setTimeout(function(){
                    window.location.href = "/";
                }, 1000);
            } else {
                layer.alert(res.msg);
            }
        }
      });
    });
}