$(document).ready(function(){
    editor();
});

var editor=function(){
    var content =$(".content");
    var sections=$(".edit-section");
    var secBtn=$(".edit-section-btn");
    sections.hide();
    secBtn.click(function(){
        sections.show();
    });
    var opts={
        btns: [
            'viewHTML',
            '|', 'formatting',
        ]
    }
    content.trumbowyg(opts);
    $("#pageForm").submit(function(){
        content.val(toMarkdown(content.val()));
    });
}