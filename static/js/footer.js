/**
 * Created by gernest on 1/10/15.
 */
$(document).ready(function () {
    var footerDown = function () {
        var winHeight = $(window).height();
        var docHeight = $(document).height();
        var footer = $(".footer-section");
        if (docHeight <= winHeight) {
            p = winHeight - docHeight;
            footer.offset({
                left: 0,
                top: docHeight - footer.height() + p,
            });
        }
    };

    var thumbNail=function(){
        var thumb=$(".thumbnail")
        var thumbImg=$(".thumbnail img")
        cPos=thumb.position();
        if(thumbImg.position()){
            iPos=thumbImg.position();
            console.log(iPos)
            console.log(cPos)
            diff=(thumb.height()-thumbImg.height())/2;
//        thumbImg.offset({ top: cPos.top+diff, left:iPos.left})
            l=iPos.left
            t=cPos.top+diff;
            console.log(l);
            off=thumbImg.offset();
            thumbImg.offset({top:off.top+diff, let: off.left});

            console.log(thumbImg.offset());
        }
        console.log("Hatarii");

    }
    footerDown();
    thumbNail();
    $(window).resize(function () {
        footerDown();
    });



})
