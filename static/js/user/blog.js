$('#payButton').popup({
    popup: $('.payQR.popup'),
    on: 'click',
    position: 'bottom center'
});

tocbot.init({
    // Where to render the table of contents.
    tocSelector: '.js-toc',
    // Where to grab the headings to build the table of contents.
    contentSelector: '.js-toc-content',
    // Which headings to grab inside of the contentSelector element.
    headingSelector: 'h1, h2, h3',
});

$('.toc.button').popup({
    popup: $('.toc-container.popup'),
    on: 'click',
    position: 'left center'
});

$('.wechat').popup({
    popup: $('.wechat-qr'),
    position: 'left center'
});

var serurl = /*[[#{blog.serurl}]]*/"127.0.0.1:8080";
var url = /*[[@{/blog/{id}(id=${blog.id})}]]*/"";
var qrcode = new QRCode("qrcode", {
    text: serurl + url,
    width: 110,
    height: 110,
    colorDark: "#000000",
    colorLight: "#ffffff",
    correctLevel: QRCode.CorrectLevel.H
});

$('#toTop-button').click(function () {
    $(window).scrollTo(0, 500);
});


var waypoint = new Waypoint({
    element: document.getElementById('waypoint'),
    handler: function (direction) {
        if (direction == 'down') {
            $('#toolbar').show(100);
        } else {
            $('#toolbar').hide(500);
        }
        console.log('Scrolled to waypoint!  ' + direction);
    }
})


//评论表单验证
$('.ui.form').form({
    fields: {
        title: {
            identifier: 'content',
            rules: [{
                type: 'empty',
                prompt: '请输入评论内容'
            }
            ]
        },
        content: {
            identifier: 'nickname',
            rules: [{
                type: 'empty',
                prompt: '请输入你的大名'
            }]
        },
        type: {
            identifier: 'email',
            rules: [{
                type: 'email',
                prompt: '请填写正确的邮箱地址'
            }]
        }
    }
});



$('#commentpost-btn').click(function () {
    var boo = $('.ui.form').form('validate form');
    if (boo) {
        console.log('校验成功');
        postData();
    } else {
        console.log('校验失败');
    }
});

function clearContent() {
    $("[name='nickname']").val('');
    $("[name='email']").val('');
    $("[name='content']").val('');
    $("[name='parentComment.id']").val(-1);
    $("[name='content']").attr("placeholder", "请输入评论信息...");
}

//markdown编辑器展示
$(function () {
    var testView = editormd.markdownToHTML("test-markdown-view", {
        //markdown : "[TOC]\n### Hello world!\n## Heading 2", // Also, you can dynamic set Markdown text
        //htmlDecode : true,  // Enable / disable HTML tag encode.
        //htmlDecode : "style,script,iframe",  // Note: If enabled, you should filter some dangerous HTML tags for website security.
    });
});