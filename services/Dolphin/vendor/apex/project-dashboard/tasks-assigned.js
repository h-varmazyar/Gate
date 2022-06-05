var options = {
	chart: {
		height: 200,
		width: 250,
		type: 'heatmap',
		toolbar: {
			show: false,
		},
	},
	stroke: {
		width: 2
	},
	plotOptions: {
		heatmap: {
			radius: 0,
			enableShades: false,
			colorScale: {
				ranges: [{
						from: 0,
						to: 4,
						color: '#3f9f78'
					},
					{
						from: 5,
						to: 8,
						color: '#cc2626'
					},
				],
			},

		}
	},
	dataLabels: {
		enabled: true,
		style: {
			colors: ['#fff']
		}
	},
	series: [{
			name: 'جعفر',
			data: [{
				x: 'شنبه',
				y: 5
			}, {
				x: 'یکشنبه',
				y: 7
			}, {
				x: 'سه شنبه',
				y: 3
			}, {
				x: 'چهارشنبه',
				y: 4
			}, {
				x: 'پنجشنبه',
				y: 2
			}]
		},
		{
			name: 'سعید',
			data: [{
				x: 'شنبه',
				y: 5
			}, {
				x: 'یکشنبه',
				y: 3
			}, {
				x: 'سه شنبه',
				y: 1
			}, {
				x: 'چهارشنبه',
				y: 3
			}, {
				x: 'پنجشنبه',
				y: 2
			}]
		},
		{
			name: 'مجتبی',
			data: [{
				x: 'شنبه',
				y: 4
			}, {
				x: 'یکشنبه',
				y: 2
			}, {
				x: 'سه شنبه',
				y: 3
			}, {
				x: 'چهارشنبه',
				y: 5
			}, {
				x: 'پنجشنبه',
				y: 1
			}]
		},
		{
			name: 'طاهر',
			data: [{
				x: 'شنبه',
				y: 2
			}, {
				x: 'یکشنبه',
				y: 4
			}, {
				x: 'سه شنبه',
				y: 5
			}, {
				x: 'چهارشنبه',
				y: 1
			}, {
				x: 'پنجشنبه',
				y: 2
			}]
		},
		{
			name: 'خان عباسی',
			data: [{
				x: 'شنبه',
				y: 1
			}, {
				x: 'یکشنبه',
				y: 2
			}, {
				x: 'سه شنبه',
				y: 4
			}, {
				x: 'چهارشنبه',
				y: 2
			}, {
				x: 'پنجشنبه',
				y: 3
			}]
		},
	],
	tooltip: {
		y: {
			formatter: function(val) {
				return val + ' کار'
			}
		}
	},
	xaxis: {
		type: 'category',
	},
}

var chart = new ApexCharts(
	document.querySelector("#tasks-assigned"),
	options
);

chart.render();