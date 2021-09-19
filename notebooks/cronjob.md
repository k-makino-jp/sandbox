# CronJob が 100 回連続で Job スケジューリングに失敗してもスケジューリングを継続する

## Reference

* [CronJob | Kubernetes: CronJobの制限](https://kubernetes.io/ja/docs/concepts/workloads/controllers/cron-jobs/#cron-job-limitations)
* [青山真也 (2020) Kubernetes 完全ガイド, 株式会社インプレス](https://book.impress.co.jp/books/1119101148)

## 本記事で解消するエラー

本記事の内容を実施することで、以下のエラーが解消されます。
```
Cannot determine if job needs to be started. Too many missed start time (> 100). Set or decrease .spec.startingDeadlineSeconds or check clock skew.
```

## エラーの詳細を理解する

本エラーは CronJob による Job のスケジューリングが 100 回連続で失敗した場合に発生します。
本エラーが発生すると、 CronJob を再作成しない限り、二度と Job は作成されません。以下、詳細について説明します。

CronJob は、指定した時刻になると Master Node が Job を作成します。
このため、 Master Node が一時的にダウンしていた場合など、 Job の開始時刻が遅れる場合があります。

例えば、運用者全員が帰宅した金曜日の 17:01 に CronJob コントローラーが応答不能になり、
火曜日の 00:00 に誰かが問題を発見した場合を考えます。
Job が 1 時間おきに作成される場合、金曜日の 17:01 から火曜日の 00:00 までに 80 個の Job の開始が保留されています。
CronJob の同時実行 (`spec.concurrencyPolicy`) と遅延開始 (`spec.startingDeadlineSeconds`) が許可されている場合に、
火曜日の 00:00 にコントローラーを再起動すると、保留されている Job は問題なく開始される必要があります。

ただし、CronJob コントローラーのサーバーまたは creationTimestamp を設定するための apiservers の時刻が正しくない場合は、開始時間の個数が非常に多くなり (数十年以上ずれている可能性があります) 、コントローラーのすべての CPU およびメモリを消費する可能性があります。

以上を考慮して、開始されるべきであったが開始されなかった Job のカウント数の上限値を 100 に設定しています。[(参考: kubernetes/utils.go at release-1.22)](https://github.com/kubernetes/kubernetes/blob/v1.22.2/pkg/controller/cronjob/utils.go#L127-L143)

## エラーに対処する

### Required: spec.startingDeadlineSeconds を設定する
`spec.startingDeadlineSeconds` は、 CronJob による Job のスケジューリングの開始時刻が遅れた場合に許容できる秒数を指定します。デフォルトは無制限であり、どれだけ開始時刻が遅れた場合でも Job をスケジューリングします。

`spec.startingDeadlineSeconds` を設定すると、最後に実行された時刻から現在までではなく、`spec.startingDeadlineSeconds` の値から現在までで、どれだけ Job を逃したのかをコントローラーがカウントします。例えば、`spec.startingDeadlineSeconds` が 200 の場合、過去 200 秒間に Job が失敗した回数を記録します。[(参考: CronJob | Kubernetes: CronJobの制限)](https://kubernetes.io/ja/docs/concepts/workloads/controllers/cron-jobs/#cron-job-limitations)

### Optional: spec.concurrencyPolicy を設定する
`spec.concurrencyPolicy` は、 Job の同時実行に関するポリシーを設定します。Job の実行が意図した間隔内で正常終了している時は、`spec.concurrencyPolicy` を指定せずとも、同時実行はされることなく、新しい Job が作成されます。

一方、古い Job がまだ実行している時は、ポリシーで新たな Job を実行するかどうかを制御したいケースがあります。このようなケースでは、`spec.concurrencyPolicy` に `Forbid` (前の Job が終了していない場合、次の Job は実行しない) を設定します。ただし、作成されなかった Job は、実行に失敗した Job としてコントローラーがカウントするため、 `spec.concurrencyPolicy` の設定は本記事のエラーの解消に寄与しないことに注意が必要です。[(参考: CronJob | Kubernetes: CronJobの制限)](https://kubernetes.io/ja/docs/concepts/workloads/controllers/cron-jobs/#cron-job-limitations)

`spec.concurrencyPolicy` に設定可能な値は [CronJobを使用して自動化タスクを実行する | Kubernetes](https://kubernetes.io/ja/docs/tasks/job/automated-tasks-with-cron-jobs/#concurrency-policy) から参照できます。

### マニフェストの例

以下は、CronJob で 30 分おきに Job が起動するようスケジューリングする例です。

```yaml
apiVersion: batch/v1
kind: CronJob
metadata:
  name: hello
spec:
  schedule: "*/30 * * * *"
  startingDeadlineSeconds: 1800
  concurrencyPolicy: Forbid
  jobTemplate:
    spec:
      template:
        spec:
          containers:
          - name: hello
            image: busybox
            command:
            - /bin/sh
            - -c
            - date; echo Hello from the Kubernetes cluster
          restartPolicy: OnFailure
```

