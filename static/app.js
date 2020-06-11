var duree = window.duree;

$(document).ready(() => {

    $('#edit').on('click', () => {
        // todo implement the edit mode
    });

    $('#save').on('click', () => {
        let bookmarks = duree.collectBookmarks();
        duree.saveBoomarks(bookmarks).done(function () {
            location.reload();
        }).fail(function (res) {
            console.log(res);
            alert("Error on saving the Bookmarks")
        });
    });

});

window.duree = {

    // collect all bookmarks on the page
    collectBookmarks: function () {
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
        return boormarks;
    },

    //sends the booksmarks to our backend
    saveBoomarks: function (bookmarks) {
        return $.ajax({
            type: "POST",
            url: "/save",
            data: JSON.stringify(bookmarks),
        });
    },

};