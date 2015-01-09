/**
 * Created by gernest on 1/9/15.
 */
$(document).ready(function(){
    $(".pick-theme").click(function(){
        var base=$(this).parents(".theme-box");
        var theme=$(".theme")
        var themeName=base.attr("t-id");
        var box =$(".theme-box");


        console.log(themeName)
        theme.val(themeName);
        box.hide();
        $(".theme-select").prepend("<div class=\"theme-notice\">"+"You Picked "+themeName+" Theme <div>")
        $(".theme-select-btn").click(function(){
            $(".theme-notice").remove();
            box.show();
        });
    });
    $(".pick-template").click(function(){
        var base=$(this).parents(".template-box");
        var template=$(".template")
        var templateName=base.attr("t-id");
        var box =$(".template-box");


        console.log(templateName)
        template.val(templateName);
        box.hide();
        $(".template-select").prepend("<div class=\"template-notice\">"+"You Picked "+templateName+" Template <div>")
        $(".template-select-btn").click(function(){
            $(".template-notice").remove();
            box.show();
        });
    });
});
