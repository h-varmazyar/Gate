// Notify examples

var notes = $('#notes').notify({
	removeIcon: '<i class="icon-close"></i>'
});

$('.add-noti').on('click', function() {
	notes.show("من یک اعلان هستم که به سرعت به شما هشدار خواهم داد!", {
		title: 'سلام',
	});
});

$('.add-success-noti').on('click', function() {
	notes.show("من یک اعلان هستم که به سرعت به شما هشدار خواهم داد!", {
		type: 'success',
		title: 'سلام',
		icon: '<i class="icon-sentiment_satisfied"></i>'
	});
});

$('.add-info-noti').on('click', function() {
	notes.show("من یک اعلان هستم که به سرعت به شما هشدار خواهم داد!", {
		type: 'info',
		title: 'سلام',
		icon: '<i class="icon-alert-circle"></i>'
	});
});

$('.add-warning-noti').on('click', function() {
	notes.show("من یک اعلان هستم که به سرعت به شما هشدار خواهم داد!", {
		type: 'warning',
		title: 'سلام',
		icon: '<i class="icon-alert-octagon"></i>'
	});
});

$('.add-danger-noti').on('click', function() {
	notes.show("من یک اعلان هستم که به سرعت به شما هشدار خواهم داد!", {
		type: 'danger',
		title: 'سلام',
		icon: '<i class="icon-alert-triangle"></i>'
	});
});

$('.add-sticky-noti').on('click', function() {
	notes.show("من یک اعلان هستم که به سرعت به شما هشدار خواهم داد!", {
		title: 'سلام',
		icon: '<i class="icon-info-outline"></i>',
		sticky: true
	});
});




/*************************
	*************************
	*************************
	*************************
	Fixed on Top
	*************************
	*************************
	*************************
	*************************/

var messages = $('#messages').notify({
	type: 'messages',
	removeIcon: '<i class="icon-close"></i>'
});

$('.add-message').on('click', function() {
	messages.show("من یک پیام هستم و به سرعت به شما هشدار خواهم داد", {
		title: 'سلام,',
	});
});

$('.add-success-message').on('click', function() {
	messages.show("من یک پیام هستم و به سرعت به شما هشدار خواهم داد", {
		type: 'success',
		title: 'سلام,',
		icon: '<i class="icon-sentiment_satisfied"></i>'
	});
});

$('.add-info-message').on('click', function() {
	messages.show("من یک پیام هستم و به سرعت به شما هشدار خواهم داد", {
		type: 'info',
		title: 'سلام,',
		icon: '<i class="icon-alert-circle"></i>'
	});
});

$('.add-warning-message').on('click', function() {
	messages.show("من یک پیام هستم و به سرعت به شما هشدار خواهم داد", {
		type: 'warning',
		title: 'سلام,',
		icon: '<i class="icon-alert-octagon"></i>'
	});
});

$('.add-danger-message').on('click', function() {
	messages.show("من یک پیام هستم و به سرعت به شما هشدار خواهم داد", {
		type: 'danger',
		title: 'سلام,',
		icon: '<i class="icon-alert-triangle"></i>'
	});
});

$('.add-sticky-message').on('click', function() {
	messages.show("من یک پیام هستم و به سرعت به شما هشدار خواهم داد", {
		title: 'سلام,',
		sticky: true
	});
});

