import $ from "jquery"
import "../styles/style.scss"

function import_page(template, file) {
    if ($('body[data-template="' + template + '"]').length) require('./pages/' + template)
}

window.$ = $

$(() => {
    $('.delete').on('click', () => {
        return confirm("Are you sure you want to delete this item?")
    })

    import_page('dashboard', 'dashboard')
    import_page('user/account', 'user/account')
})