FROM opensearchproject/opensearch:2.18.0

# シングルノードモードで起動（クラスタ構成を無効化）
ENV discovery.type=single-node
# セキュリティプラグインを無効化（認証なしでアクセス可能）
ENV plugins.security.disabled=true

EXPOSE 9200