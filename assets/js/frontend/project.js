$(document).ready(function(){
    var install=function(){
        var template=$('.template')
        var theme=$('.theme')

        template.hide();
        theme.hide();

        $('.pick-theme').click(function(){
            var base=$(this).parents(".theme-box");
            theme.val(base.attr("t-id"));
            $(this).text("Selected");
            $(this).parents("p").geopattern(base.attr('t-id'));
        });
        $(".pick-template").click(function(){
            var base=$(this).parents(".template-box");
            template.val( base.attr("t-id"));
            $(this).text("Selected");
        });
    }
    install();
});