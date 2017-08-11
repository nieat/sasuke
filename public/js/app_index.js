'use strict';

// (function () {
	console.log('aaa');
	var query_list = {};
	query_list['table'] = 'users';
	query_list['columns'] = ['name', 'id'];
	query_list['relation'] = ['articles'];
	// query_list['options']['not_in'] = {'id':2}
	query_list['options'] = { 'equal': {}, 'not_in': {}, 'in': {} };
	query_list['options']['not_in'] = { 'id': 2 };
	console.log(query_list);
	var form = new Vue({
		'el': '#query_forms',
		data: {
			params: {
				table: '',
				relation: '',
				columns: {},
				options: {}
			},
			checkedColumns: []
		},
		methods: {
			excute: function excute() {
				// Ajax通信ライブラリ
				axios.get('/result', {
					params: form.params
				}).then(function (response) {
					console.log(response);
				}).catch(function (error) {
					console.log(error);
				});
			}
		}
	});
	var a = new Vue({
		'el': '#aaa',
		data: {
			num: 1
		},
		methods: {
			test: function test() {
				alert(1);
			}
		}
	});
	console.log(a);
// });