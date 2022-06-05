Morris.Donut({
	element: 'donutFormatter',
	data: [
		{value: 155, label: 'شرکت', formatted: 'حداقل 70%' },
		{value: 12, label: 'بازار', formatted: 'تقریباً. 15%' },
		{value: 10, label: 'موفقیت', formatted: 'تقریباً. 10%' },
		{value: 5, label: 'یک برچسب واقعاً طولانی', formatted: 'تقریباً بیشتر 5%' }
	],
	resize: true,
	hideHover: "auto",
	formatter: function (x, data) { return data.formatted; },
	colors:['#1a8e5f', '#262b31', '#434950', '#63686f', '#868a90']
});