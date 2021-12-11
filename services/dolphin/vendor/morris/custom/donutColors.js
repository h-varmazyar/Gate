// Morris Donut
Morris.Donut({
	element: 'donutColors',
	data: [
		{value: 30, label: 'بازار'},
		{value: 15, label: 'شرکت'},
		{value: 10, label: 'موفقیت'},
		{value: 5, label: ' برچسب واقعاً طولانی'}
	],
	backgroundColor: '#ffffff',
	labelColor: '#666666',
	colors:['#1a8e5f', '#262b31', '#434950', '#63686f', '#868a90'],
	resize: true,
	hideHover: "auto",
	gridLineColor: "#e4e6f2",
	formatter: function (x) { return x + "%"}
});