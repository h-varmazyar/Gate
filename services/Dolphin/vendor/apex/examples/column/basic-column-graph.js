var options = {
	chart: {
		height: 350,
		type: 'bar',
	},
	plotOptions: {
		bar: {
			horizontal: false,
			endingShape: 'rounded',
			columnWidth: '35%',
		},
	},
	dataLabels: {
		enabled: false
	},
	stroke: {
		show: true,
		width: 2,
		colors: ['transparent']
	},
	series: [{
		name: 'سود خالص',
		data: [44, 55, 57, 56, 61, 58, 63, 60, 66]
	}, {
		name: 'بازگشتی',
		data: [76, 85, 101, 98, 87, 105, 91, 114, 94]
	}, {
		name: 'جریان ازاد نقدینگی',
		data: [35, 41, 36, 26, 45, 48, 52, 53, 41]
	}],
	xaxis: {
		categories: ['تیر', 'مرداد', 'شهریور', 'مهر', 'آبان', 'آذر', 'دی', 'بهمن', 'اسفند'],
	},
	yaxis: {
		title: {
			text: 'تومان (هزاران)'
		}
	},
	fill: {
		opacity: 1
	},
	tooltip: {
		y: {
			formatter: function(val) {
				return "با " + val + " تومان"
			}
		}
	},
	grid: {
		row: {
			colors: ['#f5f9fe', '#ffffff'], // takes an array which will be repeated on columns
			opacity: 0.5
		},
	},
	colors: ['#1a8e5f', '#262b31', '#434950', '#63686f', '#868a90'],
}
var chart = new ApexCharts(
	document.querySelector("#basic-column-graph"),
	options
);
chart.render();
