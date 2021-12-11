new Chartist.Bar('.barChartOrders', {
	labels: ['تعداد 1', 'تعداد 2', 'تعداد 3', 'تعداد 4'],
	series: [
		[
			{meta: 'آنلاین', value: 2000},
			{meta: 'آنلاین', value: 4000},
			{meta: 'آنلاین', value: 6000},
			{meta: 'آنلاین', value: 8000},
		],
		[
			{meta: 'مستقیم', value: 3000},
			{meta: 'مستقیم', value: 5000},
			{meta: 'مستقیم', value: 7000},
			{meta: 'مستقیم', value: 9000},
		],
	],
}, {
	reverseData: true,
	seriesBarDistance: 15,
	height: "130px",
	chartPadding: {
		right: 0,
		left: 20,
		top: 0,
		bottom: 0,
	},
	axisY: {
		offset: 30
	},
	plugins: [
		Chartist.plugins.tooltip()
	],
});
