// Copyright 2015 Geofrey Ernest a.k.a gernest, All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"): you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations
// under the License.

$(document).ready(function () {
    $(".lead-section").geopattern("loraa");
    $('.shout-out a').geopattern("lora");
    $('.lead-title').geopattern('titles')
// Animate the home page

    var animateParagraph = function () {
        var boxes = $(".sand-box p");
        var pos = boxes.position();
        boxes.snabbt({
            position: [pos.top, pos.left, 0],
            rotation: [0, 0, 2 * Math.PI],
            easing: 'spring',
            spring_constant: 0.9,
            spring_decceleration: 0.1,
            loop: 1,
            delay: 500
        }).then({
            from_position: [pos.top, pos.left, 0],
            postition: [0, pos.left, 0],
            easing: 'linear'
        });
    }
    var animateTitles = function () {
        $(".title span").each(function (idx, element) {
            var x = 20;
            var title = $(".title")
            var title_height = title.height
            var z = title.length / 2 * x - Math.abs((title.length / 2 - idx) * x);
            snabbt(element, {
                from_rotation: [0, 0, -8 * Math.PI],
                //perspective: 70,
                delay: 1000 + idx * 100,
                duration: 1000,
                easing: 'ease',
                callback: function () {
                    if (idx == title.length - 1) {
                        $(".title").snabbt({
                            offset: [0, -title_height, 0],
                            from_position: [0, title_height, 0],
                            position: [0, title_height, 0],
                            rotation: [-Math.PI / 4, 0, 0],
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
    }
//    animateParagraph();
    animateTitles();
});

