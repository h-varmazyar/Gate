var options = {
	chart: {
		height: 320,
		type: 'line',
		stacked: false,
		toolbar: {
			show: false,
		},
	},
	dataLabels: {
		enabled: false
	},
	series: [{
		name: 'سفارشات',
		type: 'column',
		data: [4, 2, 2, 5, 6, 8, 8, 7]
	},{
		name: 'فروش ها',
		type: 'column',
		data: [2, 3, 1, 4, 5, 9, 5, 8]
	},{
		name: 'بازگشتی',
		type: 'line',
		data: [20, 10, 15, 36, 44, 45, 50, 58]
	}],
	stroke: {
		width: [1, 1, 3]
	},
	title: {
		text: 'درآمد کلی در میلیون دلار از فروش آنلاین و آفلاین از سال 2010 تا 2018',
		align: 'center',
	},
	colors: ['#1a8e5f', '#262b31', '#cc2626', '#63686f', '#868a90'],
	xaxis: {
		categories: [2010, 2011, 2012, 2013, 2014, 2015, 2016, 2017, 2018],
	},
	yaxis: [{
		axisTicks: {
			show: true,
		},
		axisBorder: {
			show: true,
			color: '#1a8e5f'
		},
		labels: {
			style: {
				color: '#1a8e5f',
			}
		},
		title: {
			text: "سفارشات",
			style: {
				color: '#1a8e5f',
			}
		},
		tooltip: {
			enabled: true
		}
	},{
			seriesName: 'سفارشات',
			opposite: true,
			axisTicks: {
				show: true,
			},
			axisBorder: {
				show: true,
				color: '#262b31'
			},
			labels: {
				style: {
					color: '#262b31',
				}
			},
			title: {
				text: "فروش",
				style: {
					color: '#262b31',
				}
			},
		},{
			seriesName: 'بازگشتی',
			opposite: true,
			axisTicks: {
				show: true,
			},
			axisBorder: {
				show: true,
				color: '#cc2626'
			},
			labels: {
				style: {
					color: '#cc2626',
				},
			},
			title: {
				text: "بازگشتی ",
				style: {
					color: '#cc2626',
				}
			}
		},
	],
	legend: {
		horizontalAlign: 'center',
		offsetY: 10
	}
}

var chart = new ApexCharts(
	document.querySelector("#multiple-yaxis"),
	options
);
chart.render();