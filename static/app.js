$(document).ready(() => {
    console.log("setup loaded");

    let saveBtn = $('.save-btn');

    saveBtn.show();
    saveBtn.on('click', saveBookmarks);


});

function saveBookmarks() {

    $.ajax({
        type: "POST",
        url: '/save',
        data: data,
        success: success,
        dataType: dataType
    });

}