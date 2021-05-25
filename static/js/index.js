//注意：导航 依赖 element 模块，否则无法进行功能性操作
layui.use(['form','element','layer','jquery','carousel'], function(){
    var element = layui.element;
    var carousel = layui.carousel;
    //建造实例
    carousel.render({
        elem: '#test1'
        ,width: '100%' //设置容器宽度
        ,arrow: 'always' //始终显示箭头
        //,anim: 'updown' //切换动画方式
    });
});