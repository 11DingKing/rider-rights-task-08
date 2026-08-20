# Bug Reproduction

## 包的性质

当前 test_model_fix 保存的是被测模型修复后的结果源码，不是初始含 Bug 源码。要复现原始缺陷，必须检出下面固定的 parent SHA；不要在当前修复结果源码上期待重新出现修复前失败。生成系统使用的可信验证补丁和完整验证日志仅在本地留存，不提交到结果分支。

## 问题现象

临近期限的权益事项被升级后，新的办理期限应该从原承诺期限继续顺延，但现在总是从本轮扫描时间重新起算，已经拖延的事项反而多出一整段时间。先不要修改仓库代码，请沿调用链定位 Go 文件、符号和失效机制，用测试结果说明你确认的诊断。

## 含 Bug 版本

- 仓库：11DingKing/rider-rights-task-08
- 仓库地址：https://github.com/11DingKing/rider-rights-task-08.git
- parent SHA：f1db0f91153a32d66c5035079aeabded9c0bd6ef

## 复现步骤

```bash
git clone -- https://github.com/11DingKing/rider-rights-task-08.git bug-repro
cd bug-repro
git checkout --detach f1db0f91153a32d66c5035079aeabded9c0bd6ef
go test ./internal/domain -run "^TestEscalationDeadlineExtendsExistingDeadline$" -count=1
```

## 双架构完整错误信息

### linux/amd64

- 容器内复现预期退出码：1
- 容器内复现实际退出码：1

stdout：

```text
$ go test ./internal/domain -run "^TestEscalationDeadlineExtendsExistingDeadline$" -count=1
--- FAIL: TestEscalationDeadlineExtendsExistingDeadline (0.00s)
    task08_test.go:13: deadline moved from scan time: got 2026-08-19 14:30:00 +0000 UTC want 2026-08-19 12:30:00 +0000 UTC
FAIL
FAIL	riderguard/internal/domain	0.052s
FAIL

```

stderr：

```text
(empty)
```

### linux/arm64

- 容器内复现预期退出码：1
- 容器内复现实际退出码：1

stdout：

```text
$ go test ./internal/domain -run "^TestEscalationDeadlineExtendsExistingDeadline$" -count=1
--- FAIL: TestEscalationDeadlineExtendsExistingDeadline (0.00s)
    task08_test.go:13: deadline moved from scan time: got 2026-08-19 14:30:00 +0000 UTC want 2026-08-19 12:30:00 +0000 UTC
FAIL
FAIL	riderguard/internal/domain	0.001s
FAIL

```

stderr：

```text
(empty)
```

## 通过条件

先不要修改目标仓库代码。必须准确说明 internal/domain/escalation.go 的 NextEscalationDeadline 为什么使用扫描时刻作为基准，以及该结果如何经升级服务写回并产生额外办理时间，结论要包含具体文件、符号和完整因果机制；定向验证、相关包测试和全量回归命令的结果必须与结论一致；诊断结束时目标仓库代码、测试和配置零改动，不得删除、跳过或削弱测试。
