FROM opensearchproject/opensearch:2.18.0

# シングルノードモードで起動（クラスタ構成を無効化）
ENV discovery.type=single-node
# メモリ設定
ENV OPENSEARCH_JAVA_OPTS="-Xms1g -Xmx1g"

# セキュリティプラグインを無効化（認証なしでアクセス可能）
ENV plugins.security.disabled=true
# デモ設定のインストールを無効化
ENV DISABLE_INSTALL_DEMO_CONFIG=true
# 初期管理者パスワード（セキュリティ無効化のため実際には使用されない）
ENV OPENSEARCH_INITIAL_ADMIN_PASSWORD=admin

EXPOSE 9200