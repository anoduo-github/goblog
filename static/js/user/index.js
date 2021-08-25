// 显示公众号
$('.tencent').popup({
    popup: $('.tencent-qr'),
    position: 'bottom center'
});

// 显示微信
$('.wechat').popup({
    popup: $('.wechat-qr'),
    position: 'bottom center'
});

// 显示QQ
$('.qq').popup({
    popup: $('.qq-qr'),
    position: 'bottom center'
});

// 显示邮箱
$('.email').popup();

// 显示CSDN
$('.csdn').popup();

// 显示github
$('.github').popup();

// 导航栏显示
var waypoint = new Waypoint({
    element: document.getElementById('waypoint'),
    handler: function (direction) {
        if (direction == 'down') {
            $('#nav').show(500);
        } else {
            $('#nav').hide(500);
        }
        //console.log('Scrolled to waypoint!  ' + direction);
    }
})