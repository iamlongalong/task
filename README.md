<div align="center">
  <a href="https://taskfile.dev">
    <img src="docs/static/img/logo.svg" width="200px" height="200px" />
  </a>

  <h1>Task</h1>

  <p>
    Task is a task runner / build tool that aims to be simpler and easier to use than, for example, <a href="https://www.gnu.org/software/make/">GNU Make<a>.
  </p>

  <p>
    <a href="https://taskfile.dev/installation/">Installation</a> | <a href="https://taskfile.dev/usage/">Documentation</a> | <a href="https://twitter.com/taskfiledev">Twitter</a> | <a href="https://fosstodon.org/@task">Mastodon</a> | <a href="https://discord.gg/6TY36E39UK">Discord</a>
  </p>
</div>

## 说明

这个项目是从 [go-task/task](https://github.com/go-task/task) 中 fork 过来的。

这个项目我挺喜欢，用来做 scripts 管理很方便，目前在工作中被我用来做一些常用的环境准备。

相比于 ansible，这个项目最吸引我的点有两个： 
1. 依赖很少，一个 3MB 的二进制就能跑！
2. 基于 yaml 的配置，内容很简洁，没有 ansible 封装的奇奇怪怪的 module，所见即所得

但在使用中也遇到了一些问题，主要还是在 环境变量的传递问题！！！

作者在 Q&A 中也解释了不同 shell 间不能共享环境变量，这在同一层级的任务中其实很好理解。

我的场景是，在一个 task 中通过调用另一个 task，写法如下：
```yaml

tasks:
  parent_task:
    env:
      MY_ENV: "value from parent"
    cmds:
      - echo "this is parent, and MY_ENV is $MY_ENV"
      - task: child_task
  child_task:
    cmds:
      - 'echo "this is child, and I DO NOT HAS MY_ENV !!!  see: $MY_ENV"'
```

这非常令人疑惑 😵😵😵 为什么在子 cmd 中的 task 却不能用父 task 的环境变量呢？

所以，我 fork 了一版，增加了一个 feature: 在子 cmd 中，若使用了 task, 则会继承父 task 的环境变量。
按照原有逻辑，子 task 的环境变量继承顺序，因此修改后的环境变量继承逻辑为:

taskfile envs => dotenv envs => parent task envs => task envs。

但是需要注意，我的改动仅实现了在一个 task 的 cmd 中调用了另一个 task 的情况，**不包含**使用 dep 的方式。
