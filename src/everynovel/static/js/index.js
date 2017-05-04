
/*layui*/
// layui.use('element', function(){
//     element = layui.element();
// });
// layui.use('layer', function(){
//     var layer = layui.layer;
// });
// layui.use('form', function(){
//     var form = layui.form();
// });

;!function(){
    layer = layui.layer;
    from = layui.form();
    element = layui.element();
}();

/*vuejs*/
var HeadData = new Vue({
    el: '#head-container',
    data: {
        isLogin : 0,
        username : 'Lisi'
    },
    methods: {
        show_log : function(){
            // this.isLogin = 1;
            layer.open({
                type: 1,
                title: false,
                closeBtn: 1,
                shadeClose: true,
                resize: false,
                area: '500px',
                offset: '100px',
                content: $('#popups-regist')
            });
            element.tabChange('reglog', 'log');
        },
        show_reg : function(){
            layer.open({
                type: 1,
                title: false,
                closeBtn: 1,
                shadeClose: true,
                resize: false,
                area: '500px',
                offset: '100px',
                content: $('#popups-regist')
            });
            element.tabChange('reglog', 'reg');
        },
        create_new : function(){
            layer.open({
                type: 1,
                title: false,
                closeBtn: 1,
                shadeClose: true,
                resize: false,
                area: '500px',
                // skin: 'yourclass',
                content: $('#popups-create')
            });
        },
        read_last : function(){},
    }
})

var MainData = new Vue({
    el: '#main-container',
    data: {
        list : [],
        last : [],
        curr : '',
        // title : '',
        // model : 0,
    },
    methods: {
        word_click: function(){
            var t = event.currentTarget;
            var p = $(t).parent('.item');

            var ii = $(p).data('i');
            var id = $(p).data('id')

            // $(t).parent('.item').siblings().fadeOut();
            this.curr = this.list[ii];
            this.last.push(this.curr);

            $('#popups-create form input[name="pid"]').val(id);
            data(id);
        }
    }
})

$(document).ready(function(){
    init_layout();
    data();
});

$(window).scroll(function(){
    var s = $(window).scrollTop();
    if(s>20){
        $('#head-wrap').addClass('scroll');
    }else{
        $('#head-wrap').removeClass('scroll');
    }
})

function init_layout(){
    var w = $(window).width();
    var h = $(window).height();

    var h_h = $('#head-container').height();
    var h_f = $('#foot-container').height();

    $('#main-container').css('min-height',h-h_h-h_f-20);
}

function data(id){
    console.log(id);

    if(!id){
        MainData.last = [];
        MainData.curr = '';
    }
    
    $.ajax({
        type: "GET",
        url: "/GetWord",
        data: {id:id},
        timeout : 6000,
        dataType: "JSON",
        beforeSend: function(xhr){
            // do something
            layer.open({type:3});
        },
        success: function (data) {
            console.log(data);
            MainData.list = data.Data;

            layer.closeAll();
        },
        error: function(xhr,msg){
            alert('ERROR : '+msg);
        },
        complete: function(){
            // do something
        }
    });
}

function create_word(id){
    $.ajax({
        type: "POST",
        url: "/PostWord",
        data: $('#popups-create form').serialize(),
        timeout : 6000,
        dataType: "JSON",
        beforeSend: function(xhr){
            // do something
        },
        success: function (data) {
            console.log(data);
            layer.close(layer.index)
        },
        error: function(xhr,msg){
            alert('ERROR : '+msg);
        },
        complete: function(){
            // do something
        }
    });
}



function login(){
    $.ajax({
        type: "GET",
        url: "/GetUser",
        data: $('#popups-regist form').serialize(),
        timeout : 6000,
        dataType: "JSON",
        beforeSend: function(xhr){
            // do something
        },
        success: function (data) {
            console.log(data);
            if(!data || data.Code == 1002) return;
            HeadData.isLogin = 1;
            HeadData.username = data.Data.username;
            layer.close(layer.index);
        },
        error: function(xhr,msg){
            alert('ERROR : '+msg);
        },
        complete: function(){
            // do something
        }
    });
}

function regist(){}