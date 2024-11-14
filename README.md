# プロジェクトの立ち上げ方  

事前にgit@github.com:Semi-koron/korocupBackend2024.gitをクローンしマイグレーションまで終わらせておくこと

1 `git clone git@github.com:simesaba80/kcl-cron.git`  
2 `cd kcl-cron`  
3 `curl -L https://github.com/golang-migrate/migrate/releases/download/v4.17.1/migrate.linux-amd64.tar.gz | tar xvz` 環境に合わせて要変更  
4 `docker compose up -d`  
5 `docker-compose exec app ./migrate -database "DBURL2" -path ./db/migrations up` DBURLは対応したものを

## タスクを追加する時は

まずtaskディレクトリで`task_template.go`に従ってタスクを定義する。この時ORM部分についてはcrudディレクトリ以下に切り分けること  
最後にmain.goで呼び出す部分を追加  
Airで動かしているので変更を保存すれば即座に反映されます。
