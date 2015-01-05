/**
 * Created by gernest on 1/3/15.
 */
$(document).ready(function(){
    var boxes =$(".sand-box p");
    console.log(boxes)
    pos=boxes.position()
    var rate=0.9;
    boxes.snabbt({
        position:[pos.top,pos.left,0],
        rotation:[0,0,2*Math.PI],
        easing:'spring',
        spring_constant:0.9,
        spring_decceleration:rate,
        loop:1,
        delay:500
    }).then({
        from_position:[pos.top,pos.left,0],
        postition:[0,pos.left,0],
        easing:'linear'
    });
    var titleEl=$(".title")
    var title = titleEl.text();
    var title_height = titleEl.height();

    $(".title span").each(function(idx, element) {
        var x = 20;
        var z = title.length/2 * x - Math.abs((title.length/2 - idx) * x);
        snabbt(element, {
            from_rotation: [0, 0, -8*Math.PI],
            //perspective: 70,
            delay: 1000 + idx * 100,
            duration: 1000,
            easing: 'ease',
            callback: function() {
                if(idx == title.length - 1) {
                    $(".title").snabbt({
                        offset: [0, -title_height, 0],
                        from_position: [0, title_height, 0],
                        position: [0, title_height, 0],
                        rotation: [-Math.PI/4, 0, 0],
                        perspective: 100,
                        easing: 'linear',
                        delay: 400,
                        duration: 1000
                    }).then({
                        offset: [0, -title_height, 0],
                        from_position: [0, title_height, 0],
                        position: [0, title_height, 0],
                        easing: 'spring',
                        perspective: 100,
                        spring_constant: 0.8,
                        spring_deaccelaration: 0.99,
                        spring_mass: 2,
                    });
                }
            }
        });

    });
});