var options = {
	chart: {
		height: 450,
		type: 'line',
		zoom: {
			enabled: false
		},
	},
	dataLabels: {
		enabled: false
	},
	colors: ['#1a8e5f', '#262b31', '#434950', '#63686f', '#868a90'],
	stroke: {
		width: [3, 3, 3],
		curve: 'straight',
		dashArray: [0, 8, 5]
	},
	series: [
		{
			name: "مدت زمان جلسه",
			data: [45, 52, 38, 24, 33, 26, 21, 20, 6, 8, 15, 10]
		},
		{
			name: "بازدیدهای صفحه",
			data: [35, 41, 62, 42, 13, 18, 29, 37, 36, 51, 32, 35]
		},
		{
			name: 'مجموع بازدید ها',
			data: [87, 57, 74, 99, 75, 38, 62, 47, 82, 56, 45, 47]
		}
	],
	title: {
		text: 'آمار صفحه',
		align: 'center'
	},
	markers: {
		size: 0,
		hover: {
			sizeOffset: 6
		}
	},
	xaxis: {
		categories: ['01 اسفند', '02 اسفند', '03 اسفند', '04 اسفند', '05 اسفند', '06 اسفند', '07 اسفند', '08 اسفند', '09 اسفند',
			'10 اسفند', '11 اسفند', '12 اسفند'
		],
	},
	tooltip: {
		y: [{
			title: {
				formatter: function (val) {
					return val + " (دقیقه)"
				}
			}
		}, {
			title: {
				formatter: function (val) {
					return val + " در هر جلسه"
				}
			}
		}, {
			title: {
				formatter: function (val) {
					return val;
				}
			}
		}]
	},
	grid: {
		row: {
			colors: ['#f5f9fe', '#ffffff'], // takes an array which will be repeated on columns
			opacity: 0.5
		},
	},
}

var chart = new ApexCharts(
	document.querySelector("#stepLineChart"),
	options
);

chart.render();