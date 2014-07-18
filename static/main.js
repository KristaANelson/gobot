$('.button').on('click', function(e) {
    e.preventDefault();
    console.log('clicked!');
    var src = $(this).attr('href');
    var key = $(this).data('key');
    $.getJSON(src, function(response) {
        var centeredDiv = $('.centered-div');
        var image = centeredDiv.find('.centered-image');
        image.attr('src', response[key]);
        centeredDiv.show();
        centeredDiv.on('click', function(e2) {
            centeredDiv.hide();
            centeredDiv.off('click');
        })
    });
})