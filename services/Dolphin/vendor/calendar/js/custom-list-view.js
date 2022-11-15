// Calendar List View
document.addEventListener('DOMContentLoaded', function() {
	var calendarEl = document.getElementById('calendarListView');

	var calendar = new FullCalendar.Calendar(calendarEl, {
		plugins: [ 'list' ],

		header: {
			left: 'prev,next today',
			center: 'title',
			right: 'listDay,listWeek,dayGridMonth'
		},

		// customize the button names,
		// otherwise they'd all just say "list"
		views: {
			listDay: { buttonText: 'list day' },
			listWeek: { buttonText: 'list week' }
		},

		defaultView: 'listWeek',
		defaultDate: '2019-12-12',
		navLinks: true, // can click day/week names to navigate views
		editable: true,
		eventLimit: true, // allow "more" link when too many events
		events: [
			{
				title: 'رویداد یک روز کامل',
				start: '2019-12-01'
			},
			{
				title: 'رویداد طولانی',
				start: '2019-12-07',
				end: '2019-12-10'
			},
			{
				groupId: 999,
				title: 'تکرار رویداد',
				start: '2019-12-09T16:00:00'
			},
			{
				groupId: 999,
				title: 'تکرار رویداد',
				start: '2019-12-16T16:00:00'
			},
			{
				title: 'کنفرانس',
				start: '2019-12-11',
				end: '2019-12-13'
			},
			{
				title: 'ملاقات',
				start: '2019-12-12T10:30:00',
				end: '2019-12-12T12:30:00'
			},
			{
				title: 'ناهار',
				start: '2019-12-12T12:00:00'
			},
			{
				title: 'ملاقات',
				start: '2019-12-12T14:30:00'
			},
			{
				title: 'ساعت استراحت',
				start: '2019-12-12T17:30:00'
			},
			{
				title: 'شام',
				start: '2019-12-12T20:00:00'
			},
			{
				title: 'جشن تولد',
				start: '2019-12-13T07:00:00'
			},
			{
				title: 'کلیک رو داشبورد',
				url: 'index.html',
				start: '2019-12-28'
			}
		]
	});

	calendar.render();
});