/**
 * Created by gernest on 1/10/15.
 */
$(document).ready(function(){
    var footerDown=function(){
        var winHeight=$(window).height();
        var docHeight=$(document).height();
        var footer=$(".footer-section");
        if(docHeight<=winHeight){
            p=winHeight-docHeight;
            footer.offset({
                left:0,
                top: docHeight-footer.height()+p,
            });
        }
    };
    $(document).ready(function(){
        footerDown();
        $(window).resize(function(){
            footerDown();
        });
    });
})