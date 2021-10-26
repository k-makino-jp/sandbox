# Kubernetes Operator

## Reference
* [Kubernetes Operator とは redhat.com](https://www.redhat.com/ja/topics/containers/what-is-a-kubernetes-operator)

## Kubernetes Operator とは
Operator (カスタムコントローラー) はカスタムリソースを使用する Kubernetes へのソフトウェア拡張であり、
複雑なステートフルアプリケーションのインスタンス作成、スケーリング、属性変更等を行うコントローラーである。

## カスタムリソースとは
リソースは、Kubernetes API のエンドポイントで、特定の API オブジェクトのコレクションを保持する。
例えば、Pod の Kubernetes API はリソースであり、この API は、インスタンス化 (create)、属性変更 (apply)、状態取得 (get)、削除 (delete)などのコレクションを含有している。

カスタムリソースは、Kubernetes API の拡張で、
Kubernetes 本体のコードを変更せずに、
独自のリソースを容易に追加するための機能である。

カスタムリソースは、稼働しているクラスターに動的に登録され、現れたり、消えたりし、クラスター管理者はクラスター自体とは無関係にカスタムリソースを更新できる。
一度、カスタムリソースがインストールされると、ユーザーは kubectl を使い、Pod と同様に、オブジェクトを作成、アクセスすることが可能である。

## カスタムコントローラーとは
カスタムリソースは、単純に構造化データを格納、取り出す機能を提供する。カスタムリソースをカスタムコントローラーと組み合わせることで、インスタンス化、属性変更、状態取得等が可能となり、カスタムリソースは真の[宣言的 API](https://kubernetes.io/ja/docs/concepts/extend-kubernetes/api-extension/custom-resources/#%E5%AE%A3%E8%A8%80%E7%9A%84api) を提供する。

宣言的 API は、リソースのあるべき状態を宣言することを可能にし、Kubernetes オブジェクトの現在の状態を、あるべき状態に同期し続けるように動く。例えば、ユーザーが「コンテナは3つ起動されていること」という宣言を行えば、その宣言に従って Kubernetes が 3 つのコンテナを起動する。
宣言的 API の対義語は 命令的 API と呼ばれる。これはアプリケーションに対し、具体的な処理内容を命令することで順次実行させることを指す。

稼働しているクラスターのライフサイクルとは無関係に、カスタムコントローラーをデプロイ、更新することが可能である。
カスタムコントローラーはあらゆるリソースと連携できるが、
カスタムリソースと組み合わせると特に効果を発揮する。
カスタムリソースとカスタムコントローラーを組み合わせて使用することをオペレーターパターンと呼ぶ。

## カスタムリソース・カスタムコントローラーは何故必要なのか
Web アプリケーション、モバイルバックエンド、API サービスなどのステートレスアプリケーションは、Kubernetes の Deployment 等で管理・スケーリングできる。
これらのアプリケーションの運用方法について知識を追加する必要はない。

一方、データベースや監視システムなどのステートフルアプリケーションの場合は、
Kubernetes にはないドメイン固有の知識を追加する必要がある。
このようなアプリケーションをスケーリング、アップグレード、再構成するには、運用知識が必要である。

Kubernetes Operator はこの固有のドメイン知識を
Kubernetes 拡張として導入し、
アプリケーションのライフサイクルを管理して自動化する。

## カスタムリソースを追加する
Kubernetes は独自のリソース (カスタムリソース) を容易に追加して、Kubernetes を拡張できるように作られている。

カスタムリソース定義 (CustomrResourceDefinition; CRD) API リソースは、カスタムリソースを定義する。
CRDオブジェクトを定義することで、指定した名前、スキーマで新しいカスタムリソースが作成される。

カスタムリソースの追加は CRD を使用する以外に API アグリゲーションを使用することもできる。
カスタムリソースの追加方法の選択は[公式ドキュメント](https://kubernetes.io/ja/docs/concepts/extend-kubernetes/api-extension/custom-resources/#%E3%82%AB%E3%82%B9%E3%82%BF%E3%83%A0%E3%83%AA%E3%82%BD%E3%83%BC%E3%82%B9%E3%81%AE%E8%BF%BD%E5%8A%A0%E6%96%B9%E6%B3%95%E3%82%92%E9%81%B8%E6%8A%9E%E3%81%99%E3%82%8B)を参照する。

## Kubernetes Operator をデプロイする
オペレーターをデプロイする最も一般的な方法は、Custom Resource Definitionとそれに関連するコントローラーをクラスターに追加することである。
このコントローラーは通常のコンテナアプリケーションを動かすのと同じように、コントロールプレーン外で動作する。
例えば、コントローラーをDeploymentとしてクラスター内で動かすことができる。

## Kubernetes Operator を利用する

### etcd-operator

1. [coreos/etcd-operator](https://github.com/coreos/etcd-operator) にアクセスする。
2. Git Repository をクローンする。
```
git clone https://github.com/coreos/etcd-operator.git
```
3. カレントディレクトリを移動する。
```
cd etcd-operator
```
4. example/deployment-fix.yaml を作成する。

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: etcd-operator
spec:
  replicas: 1
  selector:
    matchLabels:
      name: etcd-operator
  template:
    metadata:
      labels:
        name: etcd-operator
    spec:
      containers:
      - name: etcd-operator
        image: quay.io/coreos/etcd-operator:v0.9.4
        command:
        - etcd-operator
        # Uncomment to act for resources in all namespaces. More information in doc/user/clusterwide.md
        #- -cluster-wide
        env:
        - name: MY_POD_NAMESPACE
          valueFrom:
            fieldRef:
              fieldPath: metadata.namespace
        - name: MY_POD_NAME
          valueFrom:
            fieldRef:
              fieldPath: metadata.name
```
5. etcd-operator をデプロイする。
```
kubectl create -f example/deployment-fix.yaml
```
6. etcd クラスタを作成する。
```
vi example/example-etcd-cluster.yaml
vi example/example-etcd-cluster-fix.yaml
```

```yaml
apiVersion: "etcd.database.coreos.com"
kind: "EtcdCluster"
metadata:
  name: "example-etcd-cluster"
  ## Adding this annotation make this cluster managed by clusterwide operators
  ## namespaced operators ignore it
  # annotations:
  #   etcd.database.coreos.com/scope: clusterwide
spec:
  size: 3
  version: "3.2.13"
```
```
kubectl create -f example/example-etcd-cluster.yaml
```

実行例

```
controlplane $ git clone https://github.com/coreos/etcd-operator.git
Cloning into 'etcd-operator'...
remote: Enumerating objects: 13742, done.
remote: Total 13742 (delta 0), reused 0 (delta 0), pack-reused 13742
Receiving objects: 100% (13742/13742), 3.24 MiB | 23.02 MiB/s, done.
Resolving deltas: 100% (8253/8253), done.

controlplane $ cd etcd-operator/

controlplane $ vi example/deployment-fix.yaml

controlplane $ kubectl create -f example/deployment-fix.yaml
deployment.apps/etcd-operator created

controlplane $ kubectl get crd
NAME                                    CREATED AT
etcdclusters.etcd.database.coreos.com   2021-10-26T23:40:06Z

controlplane $ cat example/example-etcd-cluster.yaml
apiVersion: "etcd.database.coreos.com/v1beta2"
kind: "EtcdCluster"
metadata:
  name: "example-etcd-cluster"
  ## Adding this annotation make this cluster managed by clusterwide operators
  ## namespaced operators ignore it
  # annotations:
  #   etcd.database.coreos.com/scope: clusterwide
spec:
  size: 3
  version: "3.2.13"

controlplane $ kubectl create -f example/example-etcd-cluster.yaml
etcdcluster.etcd.database.coreos.com/example-etcd-cluster created

controlplane $ kubectl get pods -l app=etcd
NAME                              READY   STATUS     RESTARTS   AGE
example-etcd-cluster-2xdxhzl9rd   1/1     Running    0          61s
example-etcd-cluster-86fswgnw26   1/1     Running    0          45s
example-etcd-cluster-8wfgjhzbrh   0/1     Init:0/1   0          5s

controlplane $ kubectl apply -f example/example-etcd-cluster.yaml 
Warning: kubectl apply should be used on resource created by either kubectl create --save-config or kubectl apply
etcdcluster.etcd.database.coreos.com/example-etcd-cluster configured
controlplane $ kubectl get pods -l app=etcd
NAME                              READY   STATUS     RESTARTS   AGE
example-etcd-cluster-2xdxhzl9rd   1/1     Running    0          2m11s
example-etcd-cluster-86fswgnw26   1/1     Running    0          115s
example-etcd-cluster-8wfgjhzbrh   1/1     Running    0          75s
example-etcd-cluster-mj76rc7qd4   0/1     Init:0/1   0          1s
```