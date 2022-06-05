new Chartist.Line('.areaChart', {
	labels: ['شنبه', 'یکش', 'دوشن', 'سه ش', 'چهار', 'پنج', 'جمعه'],
	series: [
		[
			{meta: 'درآمد', value: 5000},
			{meta: 'درآمد', value: 8000},
			{meta: 'درآمد', value: 7000},
			{meta: 'درآمد', value: 8000},
			{meta: 'درآمد', value: 5000},
			{meta: 'درآمد', value: 3000},
			{meta: 'درآمد', value: 4000},
		]
	]
}, {
	low: 0,
	lineSmooth: true,
	showArea: true,
	showLine: true,
	showPoint: true,
	showLabel: true,
	fullWidth: true,
	height: "240px",
	chartPadding: {
		right: 20,
		left: 20,
		bottom: 20,
		top: 20
	},
	axisX: {
		showGrid: true,
		showLabel: true,
	}, 
	axisY: {
		showGrid: true,
		showLabel: true,
	},
	plugins: [
		Chartist.plugins.tooltip()
	],
});