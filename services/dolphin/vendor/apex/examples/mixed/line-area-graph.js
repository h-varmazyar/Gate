var options = {
	chart: {
		height: 350,
		type: 'line',
		zoom: {
			enabled: false
		},
	},
	stroke: {
		curve: 'straight'
	},
	series: [{
		name: 'تیم یک',
		type: 'area',
		data: [40, 55, 35, 45, 30, 35, 27, 32, 33, 41, 30]
	}, {
		name: 'تیم دو',
		type: 'line',
		data: [20, 40, 25, 35, 20, 35, 37, 52, 44, 61, 60]
	}],
	fill: {
		type:'solid',
		opacity: [0.35, 1],
	},
	labels: ['اسفند 01', 'اسفند 02','اسفند 03','اسفند 04','اسفند 05','اسفند 06','اسفند 07','اسفند 08','اسفند 09 ','اسفند 10','اسفند 11'],
	markers: {
		size: 0
	},
	yaxis: [
		{
			title: {
				text: 'سری 1',
			},
		},
		{
			opposite: true,
			title: {
				text: 'سری 2',
			},
		},
	],
	tooltip: {
		shared: true,
		intersect: false,
		y: {
			formatter: function (y) {
				if(typeof y !== "undefined") {
					return  y.toFixed(0) + " نکته ها";
				}
				return y;
				
			}
		}
	},
	colors: ['#1a8e5f', '#262b31', '#434950', '#63686f', '#868a90'],
}
var chart = new ApexCharts(
	document.querySelector("#line-area-graph"),
	options
);
chart.render();