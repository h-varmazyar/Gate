var options = {
	chart: {
		height: 220,
		type: 'rangeBar',
		toolbar: {
			show: false,
		},
	},
	plotOptions: {
		bar: {
			horizontal: true,
		}
	},
	series: [{
		name: 'جعفر',
		data: [{
			x: 'وایرهای فریم',
			y: [new Date('2019-03-02').getTime(), new Date('2019-03-03').getTime()]
		}, {
			x: 'طراحی',
			y: [new Date('2019-03-02').getTime(), new Date('2019-03-04').getTime()]
		 
		}, {
			x: 'کد',
			y: [new Date('2019-03-04').getTime(), new Date('2019-03-07').getTime()]
		}]
	}, {
		name: 'مجتبی',
		data: [{
			x: 'وایرهای فریم',
			y: [new Date('2019-03-01').getTime(), new Date('2019-03-02').getTime()] 
		}, {
			x: 'طراحی',
			y: [new Date('2019-03-03').getTime(), new Date('2019-03-07').getTime()] 
		}, {
			x: 'کد',
			y: [new Date('2019-03-06').getTime(), new Date('2019-03-09').getTime()]
		}]
	}],
	yaxis: {
		min: new Date('2019-03-01').getTime(),
		max: new Date('2019-03-14').getTime()
	},
	xaxis: {
		type: 'datetime'
	},
	colors: ['#1a8e5f', '#cc2626', '#434950', '#63686f', '#868a90'],
}

var chart = new ApexCharts(
	document.querySelector("#project-timeline"),
	options
);

chart.render();