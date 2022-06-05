var chart10 = c3.generate({
	bindto: '#donutChart',
	data: {
		columns: [
			['پسندیده ها', 60],
			['اشتراک ها', 30],
			['کلیک ها', 15],
		],
		type : 'donut',
		colors: {
			Likes: '#1a8e5f',
			Shares: '#262b31',
			Clicks: '#434950',
		},
		onclick: function (d, i) { console.log("onclick", d, i); },
		onmouseover: function (d, i) { console.log("onmouseover", d, i); },
		onmouseout: function (d, i) { console.log("onmouseout", d, i); }
	},
	donut: {
		title: "کلیک ها"
	},
});