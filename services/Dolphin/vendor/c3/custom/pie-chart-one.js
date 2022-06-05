var chart12 = c3.generate({
	bindto: '#pieChart1',
	data: {
		// iris data from R
		columns: [
			['تیر', 219],
			['مرداد', 167],
			['شهریور', 115],
			['مهر', 87],
			['آبان', 60],
			['اذر', 30],
		],
		type : 'pie',
		colors: {
			Mon: '#1a8e5f',
			Tue: '#262b31',
			Wed: '#434950',
			Thu: '#63686f',
			Fri: '#868a90',
			Sat: '#999999',
		},
		onclick: function (d, i) { console.log("onclick", d, i); },
		onmouseover: function (d, i) { console.log("onmouseover", d, i); },
		onmouseout: function (d, i) { console.log("onmouseout", d, i); }
	},
});