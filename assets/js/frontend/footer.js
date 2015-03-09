// footer
$(document).ready(function(){
    var footer=function(){
        var windowHeight=$(window).height();
        var documentHeight=$(document).height();
        var currentFooter=$('#footer');
        if (documentHeight<=windowHeight){
            p=windowHeight-documentHeight;
            currentFooter.offset({
                left:0,
                top:documentHeight-currentFooter.height()+p
            })
        }
    }
    footer();
    $(window).resize(function(){
        footer();
    });
});
