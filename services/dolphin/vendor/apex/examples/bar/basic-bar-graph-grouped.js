var options = {
	chart: {
		height: 300,
		type: 'bar',
	},
	plotOptions: {
		bar: {
			horizontal: true,
			dataLabels: {
				position: 'top',
			},
		}  
	},
	dataLabels: {
		enabled: true,
		offsetX: -6,
		style: {
			fontSize: '12px',
			colors: ['#fff']
		}
	},
	stroke: {
		show: true,
		width: 1,
		colors: ['#fff']
	},
	series: [{
		name: 'سری 1',
		data: [44, 55, 41, 64, 22]
	},{
		name: 'سری 2',
		data: [53, 32, 33, 44, 32]
	}],
	xaxis: {
		categories: [2010, 2012, 2014, 2016, 2018],
	},
	colors: ['#1a8e5f', '#262b31'],
}
var chart = new ApexCharts(
	document.querySelector("#basic-bar-graph-grouped"),
	options
);
chart.render();