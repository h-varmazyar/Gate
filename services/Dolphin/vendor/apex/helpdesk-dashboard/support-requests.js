var options = {
	chart: {
		height: 150,
		type: 'area',
		toolbar: {
			show: false,
		},
	},
	dataLabels: {
		enabled: false
	},
	stroke: {
		curve: 'smooth',
		width: 3
	},
	series: [{
		name: 'درخواست پشتیبانی',
		data: [120, 100, 170, 480, 620, 500]
	}],
	grid: {
		row: {
			colors: ['#ffffff'], // takes an array which will be repeated on columns
			opacity: 0.5
		},
		padding: {
			left: 10,
			right: 10,
		},		
	},
	xaxis: {
		categories: ['شنبه', 'یکشنبه', 'سه شنبه', 'چهارشنبه', 'پنچ شنبه', 'جمعه'],
	},
	yaxis: {
		labels: {
			show: false,
		}
	},
	theme: {
		monochrome: {
			enabled: true,
			color: '#1a8e5f',
			shadeIntensity: 0.1
		},
	},
	markers: {
		size: 0,
		opacity: 0.2,
		colors: ["#1a8e5f"],
		strokeColor: "#fff",
		strokeWidth: 2,
		hover: {
			size: 7,
		}
	},
	tooltip: {
		x: {
			format: 'dd/MM/yy'
		},
	}
}

var chart = new ApexCharts(
	document.querySelector("#support-requests"),
	options
);

chart.render();
