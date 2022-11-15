// Daterange
$(function() {
	var start = moment().subtract(29, 'days');
	var end = moment();
	function cb(start, end) {
		$('#reportrange span').html(start.format('MMM D, YYYY') + ' - ' + end.format('MMM D, YYYY'));
	}
	$('#reportrange').daterangepicker({
		opens: 'right',
		startDate: start,
		endDate: end,
		ranges: {
			'امروز': [moment(), moment()],
			'دیروز': [moment().subtract(1, 'days'), moment().subtract(1, 'days')],
			'7 روز قبل': [moment().subtract(6, 'days'), moment()],
			'30 روز قبل': [moment().subtract(29, 'days'), moment()],
			'ماه اخیر': [moment().subtract(1, 'month').startOf('month'), moment().subtract(1, 'month').endOf('month')]
		}
	}, cb);
	cb(start, end);
});