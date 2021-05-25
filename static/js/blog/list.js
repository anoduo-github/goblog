//注意：导航 依赖 element 模块，否则无法进行功能性操作
layui.use(['form','element','layer','jquery'], function(){
    
});

//writeBlog 写博客
function writeBlog() {
    window.location.href="/blog/add";
}