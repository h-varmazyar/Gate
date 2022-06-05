var options = {
	chart: {
		height: 300,
		type: 'bar',
		stacked: true,
		stackType: '100%'
	},
	plotOptions: {
		bar: {
			horizontal: true,
		},
	},
	stroke: {
		width: 1,
		colors: ['#fff']
	},
	series: [{
		name: 'آیفون',
		data: [44, 55, 41, 37, 22, 43, 21]
	},{
		name: 'آیپد',
		data: [53, 32, 33, 52, 13, 43, 32]
	},{
		name: 'آیمک',
		data: [12, 17, 11, 9, 15, 11, 20]
	},{
		name: 'مک بوک',
		data: [9, 7, 5, 8, 6, 9, 4]
	},{
		name: 'مک مینی',
		data: [25, 12, 19, 32, 25, 24, 10]
	}],
	title: {
		text: 'اپ استو'
	},
	xaxis: {
		categories: [2008, 2009, 2010, 2011, 2012, 2013, 2014],
	},
	tooltip: {
		y: {
			formatter: function(val) {
				return val + "هزار"
			}
		}
	},
	fill: {
		opacity: 1
	},
	legend: {
		position: 'bottom',
		horizontalAlign: 'center',
	},
	colors: ['#1a8e5f', '#262b31', '#434950', '#63686f', '#868a90'],
}

var chart = new ApexCharts(
	document.querySelector("#basic-bar-stack-graph-full-width"),
	options
);

chart.render();
