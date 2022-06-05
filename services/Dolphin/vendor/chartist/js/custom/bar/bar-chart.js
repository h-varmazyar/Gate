new Chartist.Bar('.barChart', {
	labels: ['تعداد 1', 'تعداد 2', 'تعداد 3', 'تعداد 4'],
	series: [
		[
			{meta: 'مرد', value: 2000},
			{meta: 'مرد', value: 4000},
			{meta: 'مرد', value: 6000},
			{meta: 'مرد', value: 8000},
		],
		[
			{meta: 'زن', value: 3000},
			{meta: 'زن', value: 5000},
			{meta: 'زن', value: 7000},
			{meta: 'زن', value: 9000},
		],
	],
}, {
	reverseData: true,
	seriesBarDistance: 15,
	height: "240px",
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
