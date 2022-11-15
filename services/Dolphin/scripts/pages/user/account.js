import $ from 'jquery'

$('.verify-merchant').on('click', function () {
    let companyID = $(this).attr('data-company-id');
    let verified = ($(this).attr('data-verified').toLowerCase() === 'true');

    $.post(`./company/${companyID}/verify`, JSON.stringify(verified)).then(() => {
        window.location.reload();
    }).fail((r) => {
        alert(r.responseText)
    })
})

$('.verify-vitrine').on('click', function () {
    let vitrineID = $(this).attr('data-vitrine-id');

    $.post(`./vitrine/${vitrineID}/verify`).then(() => {
        window.location.reload();
    }).fail((r) => {
        alert(r.responseText)
    })
})

$('.edit-slug').on('click', function () {
    let companyID = $(this).attr('data-company-id');
    let slug = $(this).attr('data-slug');

    let new_slug = prompt("Slug:", slug)
    if (new_slug !== null) {
        $.post(`./company/${companyID}/edit`, JSON.stringify({
            slug: new_slug
        })).then(() => {
            window.location.reload();
        }).fail((r) => {
            alert(r.responseText)
        })
    }
})
