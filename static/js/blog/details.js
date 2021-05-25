//注意：导航 依赖 element 模块，否则无法进行功能性操作
layui.use(['form','element','layer','jquery'], function(){

});

//markdown编辑器展示
$(function() {
    var testView = editormd.markdownToHTML("test-markdown-view", {
        //markdown : "[TOC]\n### Hello world!\n## Heading 2", // Also, you can dynamic set Markdown text
        //htmlDecode : true,  // Enable / disable HTML tag encode.
        //htmlDecode : "style,script,iframe",  // Note: If enabled, you should filter some dangerous HTML tags for website security.
    });
});