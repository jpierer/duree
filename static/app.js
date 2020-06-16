var duree = window.duree;

$(document).ready(() => {

    var editMode = false;

    $('body').on('keydown', function (e) {

        // CMD + e toggles between edit and save mode
        if (e.metaKey && e.key == "e") {

            // save
            if (editMode) {
                duree.disableEditMode();
                let bookmarks = duree.collectBookmarks();
                duree.saveBoomarks(bookmarks).done(function () {
                    location.reload();
                }).fail(function (res) {
                    console.log(res);
                    alert("Error on saving the Bookmarks")
                });
                return;
            }

            // edit
            editMode = true;
            duree.enableEditMode();
            return
        }

    });

});

window.duree = {

    editor: $('.editor'),
    titleEditor: $('input[name=title]', this.editor),
    urlEditor: $('input[name=url]', this.editor),

    enableEditMode: function () {
        $('h1').css({ "cursor": "pointer" });
        $('.editmode').show()
        $('[data-editable]').on('click', function (e) {
            e.preventDefault();
            let element = $(this);
            duree.titleEditor.focus();
            duree.urlEditor.val("").prop("disabled", false);

            // assign current item id to the editor
            duree.editor.data('editid', element.data(('editable')));

            // edit link items
            if (element.hasClass('link')) {
                duree.titleEditor.val(element.data('title'))
                duree.urlEditor.val(element.data('url'))
                return;
            }
            // edit group items
            duree.titleEditor.val(element.data('group'))
            duree.urlEditor.val("").prop("disabled", true);
        });

        $('body').on('keydown', function (e) {
            if (e.key == 'Enter') {
                let updateElm = $('[data-editable="' + duree.editor.data('editid') + '"]');

                // assign editor values to the item data

                // link item
                if (updateElm.hasClass('link')) {
                    updateElm.text(duree.titleEditor.val());
                    updateElm.data('title', duree.titleEditor.val());
                    updateElm.attr('href', duree.urlEditor.val()); // not needed cause of reload after save, but hey.. :)
                    updateElm.data('url', duree.urlEditor.val());
                } else {
                    updateElm.text(duree.titleEditor.val());
                    updateElm.data('group', duree.titleEditor.val());
                }

                // clean the editor
                duree.titleEditor.val("");
                duree.urlEditor.val("");
            }
        });

    },
    disableEditMode: function () {
        $('h1').css({ "cursor": "inherit" });

    },
    // collect all bookmarks on the page
    collectBookmarks: function () {
        let bookmarks = [];
        $('.group-container').each(function () {
            let groupItem = {}
            groupItem.group = $('[data-group]', $(this)).data('group');
            groupItem.items = [];
            $('[data-item]', $(this)).each(function () {
                let item = {
                    title: $(this).data('title'),
                    url: $(this).data('url')
                };
                groupItem.items.push(item)
            });
            bookmarks.push(groupItem)
        });
        return bookmarks;
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