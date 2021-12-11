// Calendar 1 ****************************
document.addEventListener('DOMContentLoaded', function() {
	var calendarEl = document.getElementById('calendar');

	var calendar = new FullCalendar.Calendar(calendarEl, {
		plugins: [ 'interaction', 'dayGrid' ],
		header: {
			left: 'prevYear,prev,next,nextYear today',
			center: 'title',
			right: 'dayGridMonth,dayGridWeek,dayGridDay'
		},
		navLinks: true, // can click day/week names to navigate views
		editable: true,
		eventLimit: true, // allow "more" link when too many events
		events: [
			{
				title: 'رویداد یک روز کامل',
				start: '2019-08-01'
			},
			{
				title: 'رویداد طولانی',
				start: '2019-08-07',
				end: '2019-08-10'
			},
			{
				groupId: 999,
				title: 'تکرار رویداد',
				start: '2019-08-09T16:00:00'
			},
			{
				groupId: 999,
				title: 'تکرار رویداد',
				start: '2019-08-16T16:00:00'
			},
			{
				title: 'کنفرانس',
				start: '2019-08-11',
				end: '2019-08-13'
			},
			{
				title: 'ملاقات',
				start: '2019-08-12T10:30:00',
				end: '2019-08-12T12:30:00'
			},
			{
				title: 'ناهار',
				start: '2019-08-12T12:00:00'
			},
			{
				title: 'ملاقات',
				start: '2019-08-12T14:30:00'
			},
			{
				title: 'ساعت استراحت',
				start: '2019-08-12T17:30:00'
			},
			{
				title: 'شام',
				start: '2019-08-12T20:00:00'
			},
			{
				title: 'جشن تولد',
				start: '2019-08-13T07:00:00'
			},
			{
				title: 'کلیک رو داشبورد',
				url: 'index.html',
				start: '2019-08-28'
			}
		]
	});

	calendar.render();
});