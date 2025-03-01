# ドメイン

## Load Balanced Web Service
Copilot では、Load Balanced Web Service にカスタムドメインを使用する方法が 2 つあります。

1. アプリケーション作成時に `--domain` を使用すると、同じアカウントで Route 53 ドメインを関連付けることができます。
2. Environment 作成時に `--import-cert-arns` を使用して、検証済みの ACM 証明書を指定してください。

!!!attention
    現時点では、 `copilot app init` を実行した時のみに Route 53　ドメインが関連づけられます。
    Application のドメインをアップデートしたい場合 ([#3045](https://github.com/aws/copilot-cli/issues/3045)) 、新しいドメインを関連付けるために `--domain` を使用して新しい Application を作成する前に、`copilot app delete` を実行して古い Application を削除する必要があります。

## Load Balanced Web Service
[Application](../concepts/applications.en.md#追加のアプリケーション設定)で説明したように、 `copilot app init` を実行するときに Application のドメイン名を設定できます。 [Load Balanced Web Service](../concepts/services.ja.md#load-balanced-web-service) をデプロイすると以下のようなドメイン名を使ってアクセスできるようになります。



### Application に関連するルートドメインを使用する

[Application](../concepts/applications.en.md#追加のアプリケーション設定)で説明したように、 `copilot app init --domain` を実行するときに Application のドメイン名を設定できます。

**Service のデフォルトのドメイン名**

[Load Balanced Web Service](../concepts/services.ja.md#load-balanced-web-service) をデプロイすると、デフォルトでは以下のようなドメイン名を使ってアクセスできるようになります。
```
${SvcName}.${EnvName}.${AppName}.${DomainName}
```

上記は、Application Load Balancer を利用している場合です。

```
${SvcName}-nlb.${EnvName}.${AppName}.${DomainName}
```

これは、Network Load Balancer を利用している場合です。

具体的には、 `https://kudo.test.coolapp.example.aws` や `kudo-nlb.test.coolapp.example.aws:443` となります。

**ドメインエイリアスのカスタマイズ**

Copilot が Service に割り当てるデフォルトのドメイン名を使いたくない場合は、[Manifest](../manifest/overview.ja.md) の `alias` セクションを編集して、サービスのカスタム[エイリアス](https://docs.aws.amazon.com/ja_jp/Route53/latest/DeveloperGuide/resource-record-sets-choosing-alias-non-alias.html)を設定することができます。以下のスニペットは Service にエイリアスを設定する例です。

``` yaml
# copilot/{service name}/manifest.yml のなかで
http:
  path: '/'
  alias: example.aws
```

同様に、Service が Network Load Balancer を使用している場合、以下のように指定することができます。
```yaml
nlb:
  port: 443/tls
  alias: example-v1.aws
```

ただし、[サブドメインの権限を Route 53 に委任している](https://docs.aws.amazon.com/ja_jp/Route53/latest/DeveloperGuide/CreatingNewSubdomain.html#UpdateDNSParentDomain)ため、指定するエイリアスは、以下 3 つのホストゾーンのいずれかに含まれている必要があります。

- root: `${DomainName}`
- app: `${AppName}.${DomainName}`
- env: `${EnvName}.${AppName}.${DomainName}`

**裏で何がおきているか**

裏側では Copilot は

* Application を作成したアカウント内で `${AppName}.${DomainName}` というサブドメイン用のホストゾーンを作成し
* Environment があるアカウント内で `${EnvName}.${AppName}.${DomainName}` という新しい Environment 用のサブドメインのために別のホストゾーンを作成し
* Environment 用のサブドメインに使う ACM 証明書の作成と検証し
* ACM 証明書を
    - エイリアスが Application Load Balancer　(`http.alias`) として利用されている場合、HTTPS リスナーと関連づけて HTTP のトラフィックを HTTPS にリダイレクトし
    - エリアスが `nlb.alias` として利用されていて、 TLS ターミーネーションが有効な場合、 Network Load Balancer の TLS リスナーと関連づけて
* エイリアス用でオプションの A レコードを作成しています。

**どのように見えるか**

[![AWS Copilot CLI v1.8.0 Release Highlights](https://img.youtube.com/vi/Oyr-n59mVjI/0.jpg)](https://www.youtube.com/embed/Oyr-n59mVjI)

### 既存の有効な証明書に含まれるドメインを使用する

Route53 以外のドメインを所有している場合や、コンプライアンス上の理由から既存の ACM 証明書の `alias` を使用したい場合、Environment 構築時に `--import-cert-arns` フラグを指定することで、エイリアスを含む検証済みの ACM 証明書をインポートすることができます。例えば以下のようになります。

```
$ copilot env init --import-cert-arns arn:aws:acm:us-east-1:123456789012:certificate/12345678-1234-1234-1234-123456789012
```

Environment に Service をデプロイした後、その Environment に作成した ALB （Application Load Balancer） の DNS を、エイリアスのドメインがホストされている場所に A レコードとして追加してください。

ブログ記事にも[例](../../blogs/release-v118.ja.md#%E8%A8%BC%E6%98%8E%E6%9B%B8%E3%81%AE%E3%82%A4%E3%83%B3%E3%83%9D%E3%83%BC%E3%83%88)を掲載しています。

## Request-Driven Web Service
Request-Driven Web Service に[カスタムドメイン](https://docs.aws.amazon.com/ja_jp/apprunner/latest/dg/manage-custom-domains.html)を追加もできます。Load Balanced Web Service と同様に、Manifest の [`alias`](../manifest/rd-web-service.en.md#http-alias) フィールドを変更することで追加できます。
```yaml
# in copilot/{service name}/manifest.yml
http:
  path: '/'
  alias: web.example.aws
```

Load Balanced Web Service 同様、Request-Driven Web Service がドメインを使用するためには Application がドメイン (例：example.aws) と関連付けられている必要があります。

!!!info
    現時点では、`web.example.aws` のような 1 レベルのサブドメインのみをサポートしています。

    Environment レベルのドメイン (例：`web.${envName}.${appName}.example.aws`) や、Application レベルのドメイン (例：`web.${appName}.example.aws`)、
    ルートドメイン (例：`example.aws`) はまだサポートされていません。これは、サブドメインが Application 名と衝突してはいけないということでもあります。

Copilot は内部的には以下のような処理を行なっています。

* ドメインを app runner service に関連付けます。
* ルートドメインのホストゾーンにドメインレコードと検証レコードを作成します。
