/**
 * Created by gernest on 1/3/15.
 */
$(document).ready(function(){
    var boxes =$(".sand-box .title");
    console.log(boxes)
    pos=boxes.position()
    var rate=0.9;
    boxes.snabbt({
        position:[pos.top,pos.left,0],
        rotation:[0,0,2*Math.PI],
        easing:'spring',
        spring_constant:0.1,
        spring_decceleration:rate,
//        loop:3,
        delay:500
    });
//    boxes.each(function(){
//        var elem= $(this);
//        var pos =elem.position();
//        elem.snabbt({
//            position:[pos.top,pos.left,0],
//            rotation:[0,0,2*Math.PI],
//            easing:'spring',
//            spring_constant:0.1,
//            spring_decceleration:rate,
//            delay:500
//        });
//    });
});