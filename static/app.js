$(document).ready(() => {

    let setupBtn = $('#setup');
    setupBtn.on('click', () => {

        let boormarks = [];

        $('div[data-group]').each(function () {

            let groupItem = {}
            groupItem.group = $(this).data('group');
            groupItem.items = [];

            $('a[data-item]', $(this)).each(function () {

                let item = {
                    title: $(this).data('title'),
                    url: $(this).data('url')
                };

                groupItem.items.push(item)

            });

            boormarks.push(groupItem)

        });

        console.log(boormarks);


        $.ajax({
            type: "POST",
            url: "/save",
            data: JSON.stringify(boormarks),
        }).done(function (res) {
            location.reload();
        }).fail(function (res) {
            alert("Error on saving Bookmarks")
            console.log(res);
        });

    });

});