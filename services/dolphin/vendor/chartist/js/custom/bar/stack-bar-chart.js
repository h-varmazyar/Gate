new Chartist.Bar('.stacked-bar', {
	labels: ['شنبه', 'یکشنبه', 'دوشنبه', 'سه شنبه', 'چهارشنبه', 'پنجشنبه', 'جمعه'],
	series: [
		[
			{meta: 'سفارش های آنلاین', value: 55},
			{meta: 'سفارش های آنلاین', value: 83},
			{meta: 'سفارش های آنلاین', value: 72},
			{meta: 'سفارش های آنلاین', value: 68},
			{meta: 'سفارش های آنلاین', value: 57},
			{meta: 'سفارش های آنلاین', value: 41},
			{meta: 'سفارش های آنلاین', value: 30}
		],
		[
			{meta: 'سفارش های آفلاین', value: 35},
			{meta: 'سفارش های آفلاین', value: 52},
			{meta: 'سفارش های آفلاین', value: 37},
			{meta: 'سفارش های آفلاین', value: 45},
			{meta: 'سفارش های آفلاین', value: 35},
			{meta: 'سفارش های آفلاین', value: 27},
			{meta: 'سفارش های آفلاین', value: 19}
		],
		[
			{meta: 'بازدید ها', value: 12},
			{meta: 'بازدید ها', value: 25},
			{meta: 'بازدید ها', value: 22},
			{meta: 'بازدید ها', value: 30},
			{meta: 'بازدید ها', value: 43},
			{meta: 'بازدید ها', value: 39},
			{meta: 'بازدید ها', value: 24}
		],
	],
}, {
	stackBars: true,
	seriesBarDistance: 4,
	height: "176px",
	chartPadding: {
		left: 10,
		top: 0,
		bottom: 0,
	},
	axisX: {
		offset: 20,
	}, 
	axisY: {
		showLabel: true,
		showGrid: false,
		offset: 30,
	},
	plugins: [
		Chartist.plugins.tooltip()
	], 
}).on('draw', function(data) {
	if(data.type === 'bar') {
		data.element.attr({
			style: 'stroke-width: 32px'
		});
	}
});