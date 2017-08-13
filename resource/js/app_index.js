// ()=>{
	var add_form = new Vue({
		el: '#query_form',
		methods: {
			execute: function() {
				// 各input要素の追加
				this.append_form_input('table_name')
				this.append_form_input('column_names')
				this.append_form_input('relation_name')
				this.append_form_input('relation_column_names')
				document.ex_form.submit();
			},
			append_form_input: function(name_atr) {
				var add_form = document.getElementById('add_form')
				// SQLの条件をリクエストパラメータに追加
				var input = document.createElement('input')
				input.type = 'hidden'
				input.name = name_atr
				input.value = document.getElementById(name_atr).dataset.dataname
				add_form.appendChild(input)
			}
		}
	})

	// todo 追加オプション表示時に後ほど実装
	// var add_menu = new Vue({
	// 	el: '#add_menu',
	// 	data: {
	// 		show: false
	// 	}
	// })
// 
// }
