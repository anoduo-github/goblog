//注意：导航 依赖 element 模块，否则无法进行功能性操作
layui.use(['form','element'], function(){
});

//login 登录
function login() {
    layui.use(['form','element','layer','jquery'], function(){
      var $ = layui.jquery;
      var layer = layui.layer;
      var form_data = $("form").serialize();

      //检查
      var arr=form_data.split("&");
      for(var i=0;i<arr.length;i++){
          var v=arr[i].split("=");
          switch (v[0]) {
              case "username":
                  if(v[1]==""){
                      layer.alert("用户名不能为空");
                      return;
                  }
                  break;
              case "password":
                  if(v[1]==""){
                      layer.alert("密码不能为空");
                      return;
                  }
                  break;
              case "code":
                  if(v[1]==""){
                      layer.alert("验证码不能为空");
                      return;
                  }
                  break;
          }
      }

      $.ajax({
          url: '/login',
          type: 'POST',
          datatype: 'json',
          data: form_data,
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

//redirect 跳转
function redirect() {
    window.location.href="/registerpage";
}