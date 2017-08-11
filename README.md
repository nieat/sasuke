# FREE BI TOOL

go/revelで作るBIツールです

## 命名規則
変数はスネーク
メソッドはキャメルで


## Go to http://localhost:9000/ and you'll see:

## Code Layout

    conf/             Configuration directory
        app.conf      Main app configuration file
        routes        Routes definition file

    app/              App sources
        init.go       Interceptor registration
        controllers/  App controllers go here
        		   /api   APIとajaxはここ
        views/        Templates directory
        config/       viewとかバリデーションで使う設定ファイルはここにおいてcontrollerで呼び出す
        service/      generate.go みたいに機能ごとにファイル作っていきたい csv.goとかとか
        
    messages/         Message files

	resource/         
        	scss/     
        	js/      .vueファイルとか置く
      
    public/           Public static assets
        css/          CSS files
        js/           Javascript files
        images/       Image files	

    tests/            Test suites
    .env              設定とかがここに書かれるようにする
	npm 
	
設定は.envファイルに追記する
IP制限/adminユーザーしかできない設定 ->filterでやる

クエリービルダーは以下を使用する
https://github.com/doug-martin/goqu
