FROM docker.elastic.co/elasticsearch/elasticsearch:8.16.2

# シングルノードモードで起動（クラスタ構成を無効化）
ENV discovery.type=single-node

# セキュリティ機能を無効化（認証なしでアクセス可能）
ENV xpack.security.enabled=false

# メモリ設定
ENV ES_JAVA_OPTS="-Xms1g -Xmx1g"

EXPOSE 9200